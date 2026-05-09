package config

type kafkaConfig struct {
	Topic         string
	ConsumerGroup string
	Host          string
}

const (
	basicTopic = "local_topic"
	inboxTopic = "inbox_topic"
)

func NewKafkaConfig() *kafkaConfig {
	return &kafkaConfig{
		Topic:         inboxTopic,
		ConsumerGroup: "local_cg",
		Host:          "kafka:29092",
	}
}
