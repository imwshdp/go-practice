package inbox

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/rand"
	"time"

	"kafka/internal/config"
	"kafka/internal/inbox/consumer"
	"kafka/internal/inbox/producer"
	"kafka/internal/inbox/repository"
	"kafka/internal/inbox/types"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	producer  *producer.KafkaProducer
	consumer  *consumer.KafkaConsumer
	msgChan   chan *types.Message
	eventRepo *repository.EventRepo
}

func newServer(
	producer *producer.KafkaProducer,
	consumer *consumer.KafkaConsumer,
	eventRepo *repository.EventRepo,
	msgChan chan *types.Message,
) *Server {
	return &Server{
		producer:  producer,
		consumer:  consumer,
		msgChan:   msgChan,
		eventRepo: eventRepo,
	}
}

func SetupServer(ctx context.Context, eventRepo *repository.EventRepo) (*Server, error) {
	kafkaCfg := config.NewKafkaConfig()

	kafkaProducer, err := producer.NewKafkaProducer(
		kafkaCfg.Topic, kafkaCfg.Host,
	)
	if err != nil {
		return nil, err
	}

	msgChan := make(chan *types.Message, 64)

	kafkaConsumer, err := consumer.NewKafkaConsumer(
		ctx, kafkaCfg.Topic, kafkaCfg.Host, kafkaCfg.ConsumerGroup, 0, msgChan,
	)
	if err != nil {
		return nil, err
	}

	srv := newServer(
		kafkaProducer,
		kafkaConsumer,
		eventRepo,
		msgChan,
	)
	return srv, nil
}

func (s *Server) StartProducing() {
	ticker := time.NewTicker(time.Duration(time.Second))
	defer ticker.Stop()

	for range ticker.C {
		msg := repository.NewEvent()

		b, err := json.Marshal(msg)
		if err != nil {
			panic("unable to marshal json")
		}

		s.producer.Produce(b)
	}
}

func (s *Server) handleMessage(ctx context.Context, msg *types.Message) {
	// simulate processing time
	randInt := rand.Intn(5)
	time.Sleep(time.Duration(randInt+1) * time.Second)

	id, isDuplicate, err := s.saveToDB(ctx, msg)
	if isDuplicate {
		slog.Info("Event already existed:", slog.Int64("offset", int64(msg.Metadata.Offset)), slog.String("event_id", msg.Event.EventID))
		return
	}

	if err != nil {
		slog.Error("Failed to store event:", slog.String("error", err.Error()))
		return
	}

	slog.Info("Event stored:", slog.String("id", id), slog.Int64("offset", int64(msg.Metadata.Offset)))
}

func (s *Server) saveToDB(ctx context.Context, msg *types.Message) (id string, isDuplicate bool, err error) {
	id, err = repository.TxClosure(ctx, s.eventRepo, func(ctx context.Context, tx *sqlx.Tx) (string, error) {
		// INFO: сan avoid with setting event_id as unique in DB
		event, err := s.eventRepo.Get(ctx, tx, msg.Event.EventID)
		if err != nil {
			return "", err
		}

		if event != nil {
			isDuplicate = true
			return "", nil
		}

		return s.eventRepo.Insert(ctx, tx, msg.Event)
	})
	if err != nil {
		return "", false, err
	}

	s.consumer.MarkAsComplete(msg.Metadata)
	return id, isDuplicate, nil
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := repository.NewDBConn()
	if err != nil {
		panic(err)
	}

	eventRepo := repository.NewEventRepo(db)

	srv, err := SetupServer(ctx, eventRepo)
	if err != nil {
		panic(err)
	}

	go srv.StartProducing()

	for msg := range srv.msgChan {
		go srv.handleMessage(ctx, msg)
	}
}
