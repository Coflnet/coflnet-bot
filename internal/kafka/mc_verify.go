package kafka

import (
	"context"
	kafkago "github.com/segmentio/kafka-go"
	"time"
)

func ConsumeMcVerify() (*kafkago.Message, error) {
	ctx := context.Background()
	msg, err := mcVerifyReader.FetchMessage(ctx)

	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func CommitMcVerify(msg *kafkago.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := mcVerifyReader.CommitMessages(ctx, *msg)

	if err != nil {
		return err
	}

	return nil
}
