package main

import (
	"fmt"
	"time"

	"kafka/internal/consumer"
	"kafka/internal/producer"
	config "kafka/internal/shared"
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

func SetupServer() (*Server, error) {
	kafkaCfg := config.NewKafkaConfig()

	kafkaProducer, err := producer.NewKafkaProducer(
		kafkaCfg.Topic, kafkaCfg.Host,
	)
	if err != nil {
		return nil, err
	}

	msgChan := make(chan string, 64)

	kafkaConsumer, err := consumer.NewKafkaConsumer(
		kafkaCfg.Topic, kafkaCfg.Host, kafkaCfg.ConsumerGroup, msgChan,
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

func main() {
	srv, err := SetupServer()
	if err != nil {
		panic(err)
	}

	go srv.StartProducing()

	for msg := range srv.msgChan {
		go srv.handleMsg(msg)
	}
}
