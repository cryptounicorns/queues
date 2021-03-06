package kafka

import (
	"github.com/corpix/loggers"

	"github.com/cryptounicorns/queues/consumer"
	"github.com/cryptounicorns/queues/producer"
)

const (
	Name = "kafka"
)

type Kafka struct {
	config Config
	log    loggers.Logger
}

func (q *Kafka) Producer() (producer.Producer, error) {
	return ProducerFromConfig(q.config, q.log)
}

func (q *Kafka) Consumer() (consumer.Consumer, error) {
	return ConsumerFromConfig(q.config, q.log)
}

func (q *Kafka) Close() error { return nil }

func FromConfig(c Config, l loggers.Logger) *Kafka {
	return &Kafka{
		config: c,
		log:    l,
	}
}
