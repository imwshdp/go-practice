package types

import (
	"encoding/json"
	"fmt"

	"kafka/internal/inbox/repository"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Message struct {
	Metadata *kafka.TopicPartition
	Event    *repository.Event
}

func NewMessage(metadata *kafka.TopicPartition, data []byte) *Message {
	e := &repository.Event{}
	err := json.Unmarshal(data, e)
	if err != nil {
		panic(fmt.Sprintf("err unmarshalling event = %v\n", err))
	}
	return &Message{
		Metadata: metadata,
		Event:    e,
	}
}
