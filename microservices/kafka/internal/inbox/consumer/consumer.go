package consumer

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"kafka/internal/inbox/types"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	consumer            *kafka.Consumer
	topic               string
	partition           int32
	groupID             string
	msgCh               chan<- *types.Message
	isReady             bool
	mu                  sync.Mutex
	msgStateMap         map[kafka.Offset]bool
	lastCommittedOffset kafka.Offset
	lastReceivedOffset  kafka.Offset
	commitDuration      time.Duration
}

func NewKafkaConsumer(ctx context.Context, topic, host, groupID string, partition int, msgCh chan<- *types.Message) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  host,
		"group.id":           groupID,
		"enable.auto.commit": false,
		"auto.offset.reset":  "earliest",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	offsets, err := c.Committed([]kafka.TopicPartition{{
		Topic: &topic, Partition: int32(partition),
	}}, int(5*time.Second))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch committed offsets: %w", err)
	}

	lastCommitted := offsets[0].Offset

	if lastCommitted < 0 {
		lastCommitted = kafka.OffsetBeginning
		slog.Info("No previous offset found, starting from beginning")
	} else {
		slog.Info("Resuming from committed offset", slog.Int64("offset", int64(lastCommitted)))
	}

	consumer := &KafkaConsumer{
		consumer:            c,
		topic:               topic,
		partition:           int32(partition),
		groupID:             groupID,
		msgCh:               msgCh,
		isReady:             false,
		msgStateMap:         make(map[kafka.Offset]bool),
		lastCommittedOffset: lastCommitted,
		lastReceivedOffset:  lastCommitted,
		mu:                  sync.Mutex{},
		commitDuration:      5 * time.Second,
	}

	if errInit := consumer.initializeKafkaTopic(ctx, host, topic); errInit != nil {
		return nil, fmt.Errorf("failed to initialize kafka topic: %w", errInit)
	}

	err = c.Assign([]kafka.TopicPartition{{
		Topic:     &consumer.topic,
		Partition: consumer.partition,
		Offset:    lastCommitted,
	}})
	if err != nil {
		return nil, err
	}

	go consumer.commitOffsetLoop(ctx)
	go consumer.startConsuming(ctx, 1*time.Second)

	return consumer, nil
}

func (c *KafkaConsumer) initializeKafkaTopic(ctx context.Context, host, topicName string) error {
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": host,
	})
	if err != nil {
		return err
	}
	defer adminClient.Close()

	slog.Info("Initializing topic:", slog.String("topic", topicName))
	topicSpec := kafka.TopicSpecification{
		Topic:             topicName,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	initCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	results, err := adminClient.CreateTopics(initCtx, []kafka.TopicSpecification{topicSpec})
	if err != nil {
		return err
	}

	for _, result := range results {
		if result.Error.Code() == kafka.ErrTopicAlreadyExists {
			slog.Info("Topic already exists:", slog.String("error", result.Error.String()))
			continue
		}
		if result.Error.Code() != kafka.ErrNoError {
			return fmt.Errorf("failed to create topic: %w", result.Error)
		}
		slog.Info("Topic created successfully:", slog.String("topic", result.Topic))
	}

	return c.waitForTopicReady(host, topicName)
}

func (c *KafkaConsumer) waitForTopicReady(host, topicName string) error {
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": host,
	})
	if err != nil {
		return err
	}
	defer adminClient.Close()

	for {
		time.Sleep(1 * time.Second)

		metadata, err := adminClient.GetMetadata(&topicName, false, 5000)
		if err != nil {
			slog.Error("Metadata fetch failed:", slog.String("error", err.Error()))
			continue
		}

		topicMeta, exists := metadata.Topics[topicName]
		if !exists {
			continue
		}

		if len(topicMeta.Partitions) > 0 {
			allPartitionsReady := true
			for _, partition := range topicMeta.Partitions {
				if partition.Error.Code() != kafka.ErrNoError {
					allPartitionsReady = false
					break
				}
				if partition.Leader == -1 {
					allPartitionsReady = false
					break
				}
			}

			slog.Info("Topic initialization:", slog.Bool("status", allPartitionsReady))

			if allPartitionsReady {
				return nil
			}
		}
	}
}

func (c *KafkaConsumer) readyCheck() (bool, error) {
	assignment, err := c.consumer.Assignment()
	if err != nil {
		return false, fmt.Errorf("failed to get assignment: %w", err)
	}

	if len(assignment) == 0 {
		return false, nil
	}

	return true, nil
}

func (c *KafkaConsumer) startConsuming(ctx context.Context, timeout time.Duration) {
	defer c.consumer.Close()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := c.consumer.ReadMessage(timeout)
			if err != nil {
				var kafkaErr kafka.Error

				if errors.As(err, &kafkaErr) && kafkaErr.IsTimeout() {
					continue
				}

				slog.Error("Read error:", slog.String("error", err.Error()))
				continue
			}

			c.registerNewMessage(msg.TopicPartition.Offset)

			payload := types.NewMessage(&msg.TopicPartition, msg.Value)
			c.msgCh <- payload
		}
	}
}

func (c *KafkaConsumer) registerNewMessage(offset kafka.Offset) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.msgStateMap[offset] = false

	if offset > c.lastReceivedOffset {
		c.lastReceivedOffset = offset
	}
}

func (c *KafkaConsumer) MarkAsComplete(tp *kafka.TopicPartition) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.msgStateMap[tp.Offset] = true
	slog.Debug("Message processed", slog.Int64("offset", int64(tp.Offset)))
}

func (c *KafkaConsumer) commitOffsetLoop(ctx context.Context) {
	ticker := time.NewTicker(c.commitDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			c.mu.Lock()

			toCommitOffset := c.lastReceivedOffset

			startOffset := c.lastCommittedOffset
			if startOffset < 0 {
				startOffset = 0
			}

			if c.lastCommittedOffset > c.lastReceivedOffset {
				panic("last commit is above last received")
			}

			for offset := startOffset; offset < c.lastReceivedOffset; offset++ {
				completed, exists := c.msgStateMap[offset]
				if !exists {
					continue
				}

				if !completed {
					break
				}

				delete(c.msgStateMap, offset)
				toCommitOffset = offset + 1
			}

			c.mu.Unlock()

			// INFO: can happen when we have uncompleted messages, but they are not in the map
			// (e.g. we received messages, but they are not marked as complete yet)
			if toCommitOffset <= c.lastCommittedOffset {
				continue
			}

			_, err := c.consumer.CommitOffsets([]kafka.TopicPartition{
				{
					Topic:     &c.topic,
					Partition: c.partition,
					Offset:    toCommitOffset,
				},
			})
			if err != nil {
				slog.Error("Failed to commit offset", slog.Int64("offset", int64(toCommitOffset)), slog.String("error", err.Error()))
				continue
			}

			c.mu.Lock()
			c.lastCommittedOffset = toCommitOffset

			fmt.Printf("state AFTER commit\n")
			for offset, v := range c.msgStateMap {
				fmt.Printf("%d: %t\n", offset, v)
			}

			c.mu.Unlock()

			slog.Info("Committed offset:", slog.Int64("offset", int64(toCommitOffset)))
		}
	}
}
