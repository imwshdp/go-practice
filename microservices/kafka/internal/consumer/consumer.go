package consumer

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
	topic    string
	groupID  string
	msgCh    chan<- string
	isReady  bool
}

func NewKafkaConsumer(topic, host, groupID string, msgCh chan<- string) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": host,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	consumer := &KafkaConsumer{
		consumer: c,
		topic:    topic,
		groupID:  groupID,
		msgCh:    msgCh,
	}

	if errInit := consumer.initializeKafkaTopic(host, topic); errInit != nil {
		return nil, errInit
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}

	go consumer.checkReadyToAccept()
	go consumer.startConsuming()

	return consumer, nil
}

func (c *KafkaConsumer) initializeKafkaTopic(host, topicName string) error {
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

	// TODO: add ctx using from main package
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	results, err := adminClient.CreateTopics(ctx, []kafka.TopicSpecification{topicSpec})
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

func (c *KafkaConsumer) checkReadyToAccept() {
	for {
		isReady, err := c.readyCheck()

		if err != nil {
			slog.Error("Consumer readycheck:", slog.String("error", err.Error()))
		} else if isReady {
			slog.Info("Consumer readycheck:", slog.Bool("status", isReady))
			c.isReady = true
			return
		} else {
			slog.Warn("Consumer readycheck:", slog.Bool("status", isReady))
		}

		// TODO: change to ticker
		time.Sleep(1 * time.Second)
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

func (c *KafkaConsumer) startConsuming() {
	defer c.consumer.Close()

	for {
		msg, err := c.consumer.ReadMessage(time.Second)
		if err != nil {
			var kafkaErr kafka.Error

			if errors.As(err, &kafkaErr) && kafkaErr.IsTimeout() {
				continue
			}

			slog.Error("Consumer read error:", slog.String("error", err.Error()))
			continue
		}

		c.msgCh <- string(msg.Value)
	}
}
