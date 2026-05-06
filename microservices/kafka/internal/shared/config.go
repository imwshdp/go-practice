package config

type kafkaConfig struct {
	Topic         string
	ConsumerGroup string
	Host          string
}

func NewKafkaConfig() *kafkaConfig {
	return &kafkaConfig{
		Topic:         "local_topic",
		ConsumerGroup: "local_cg",
		Host:          "kafka:29092",
	}
}
