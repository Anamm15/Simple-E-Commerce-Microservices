package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type Producer interface {
	SendMessage(topic string, key string, message interface{}) error
	Close() error
}

type producer struct {
	syncProducer sarama.SyncProducer
}

func NewProducer(brokers []string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll // Waiting for confirmation from all brokers
	config.Producer.Retry.Max = 5

	p, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &producer{syncProducer: p}, nil
}

func (p *producer) SendMessage(topic string, key string, message interface{}) error {
	// Encode message to JSON
	val, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(val),
	}

	partition, offset, err := p.syncProducer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	log.Printf("Message sent to topic %s [partition: %d, offset: %d]", topic, partition, offset)
	return nil
}

func (p *producer) Close() error {
	return p.syncProducer.Close()
}
