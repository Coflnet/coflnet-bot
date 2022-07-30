package kafka

import (
	"context"
	kafkago "github.com/segmentio/kafka-go"
	"time"
)

func ConsumeFlipSummary() (*kafkago.Message, error) {

	ctx := context.Background()
	msg, err := flipSummaryReader.FetchMessage(ctx)

	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func CommitFlipSummary(msg *kafkago.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := flipSummaryReader.CommitMessages(ctx, *msg)

	if err != nil {
		return err
	}

	return nil
}
