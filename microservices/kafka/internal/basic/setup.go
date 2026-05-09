package basic

import (
	"context"
	"fmt"
	"time"

	"kafka/internal/basic/consumer"
	"kafka/internal/basic/producer"
	"kafka/internal/config"
)

type Server struct {
	producer *producer.KafkaProducer
	consumer *consumer.KafkaConsumer
	msgChan  chan string
}

func newServer(
	producer *producer.KafkaProducer,
	consumer *consumer.KafkaConsumer,
	msgChan chan string,
) *Server {
	return &Server{
		producer: producer,
		consumer: consumer,
		msgChan:  msgChan,
	}
}

func SetupServer(ctx context.Context) (*Server, error) {
	kafkaCfg := config.NewKafkaConfig()

	kafkaProducer, err := producer.NewKafkaProducer(
		kafkaCfg.Topic, kafkaCfg.Host,
	)
	if err != nil {
		return nil, err
	}

	msgChan := make(chan string, 64)

	kafkaConsumer, err := consumer.NewKafkaConsumer(
		ctx, kafkaCfg.Topic, kafkaCfg.Host, kafkaCfg.ConsumerGroup, msgChan,
	)
	if err != nil {
		return nil, err
	}

	srv := newServer(kafkaProducer, kafkaConsumer, msgChan)
	return srv, nil
}

func (s *Server) StartProducing() {
	ticker := time.NewTicker(time.Duration(time.Second))
	defer ticker.Stop()

	id := 0

	for t := range ticker.C {
		msg := fmt.Sprintf(
			"msgID = %d, ts = %s",
			id,
			t.Format("15:04:05"),
		)

		s.producer.Produce(msg)
		id++
	}
}

func (s *Server) handleMsg(msg string) {
	// db operation for example
	fmt.Printf("Received message: %s\n", msg)
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv, err := SetupServer(ctx)
	if err != nil {
		panic(err)
	}

	go srv.StartProducing()

	for msg := range srv.msgChan {
		go srv.handleMsg(msg)
	}
}
