package kafka

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

// Handler function signature
type MessageHandler func(ctx context.Context, key []byte, value []byte) error

type ConsumerGroupHandler struct {
	Handler MessageHandler
}

func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// call handler injected
		if err := h.Handler(session.Context(), msg.Key, msg.Value); err != nil {
			log.Printf("Error processing message: %v", err)
			// not to mark message as processed (opsional, for retry strategy)
			continue
		}

		// Mark message as processed (Commit offset)
		session.MarkMessage(msg, "")
	}
	return nil
}

type Consumer struct {
	brokers []string
}

func NewConsumer(brokers []string) *Consumer {
	return &Consumer{brokers: brokers}
}

func (c *Consumer) StartConsumerGroup(ctx context.Context, groupID string, topics []string, handler MessageHandler) error {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err := sarama.NewConsumerGroup(c.brokers, groupID, config)
	if err != nil {
		return err
	}
	defer group.Close()

	consumerHandler := &ConsumerGroupHandler{Handler: handler}

	log.Printf("Kafka consumer group '%s' started for topics: %v", groupID, topics)

	for {
		// This loop is important because if rebalancing occurs, Consumption will return, and we will have to connect again.
		err := group.Consume(ctx, topics, consumerHandler)
		if err != nil {
			log.Printf("Error from consumer: %v", err)
			return err
		}
		// Check if context has cancelled (app shutdown)
		if ctx.Err() != nil {
			return nil
		}
	}
}
