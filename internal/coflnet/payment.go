package coflnet

import (
	"context"
	"errors"
	"golang.org/x/exp/slog"
	"strconv"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/schemas/payments"
)

const (
	PaymentTraceName = "payment-api"
)

func NewPaymentApi() *PaymentApi {
	var err error
	r := &PaymentApi{}
	r.tracer = otel.Tracer(PaymentTraceName)

	paymentBaseUrl, err := utils.PaymentBaseUrl()
	if err != nil {
		slog.Error("error getting payment api url", err)
		return r
	}

	r.paymentClient, err = payments.NewClient(paymentBaseUrl)
	if err != nil {
		slog.Error("error creating payment api client", err)
	}

	return r
}

type PaymentApi struct {
	tracer        trace.Tracer
	paymentClient *payments.Client
}

func (p *PaymentApi) OwningTimesOfUser(ctx context.Context, userId int) ([]model.OwnedProducts, error) {
	if p.paymentClient == nil {
		return nil, errors.New("payment api client not initialized")
	}

	_, span := p.tracer.Start(ctx, "owning-times-of-user")
	defer span.End()
	span.SetAttributes(attribute.Int("user-id", userId))

	products := []string{"premium", "premium-plus", "pre-api"}

	result := make([]model.OwnedProducts, len(products))
	ch := make(chan model.OwnedProducts, len(products))
	wg := sync.WaitGroup{}

	for _, product := range products {
		wg.Add(1)
		go func(product string) {
			t, err := p.OwningTimeProductOfUser(ctx, userId, product)
			if err != nil {
				span.RecordError(err)
			}

			ch <- model.OwnedProducts{
				ProductSlug: product,
				ExpiresAt:   t,
			}
		}(product)
	}

	go func() {
		for r := range ch {
			result = append(result, r)
			wg.Done()
		}
	}()

	wg.Wait()
	close(ch)

	return result, nil
}

func (p *PaymentApi) OwningTimeProductOfUser(ctx context.Context, userId int, productSlug string) (time.Time, error) {
	if p.paymentClient == nil {
		return time.Time{}, errors.New("payment api client not initialized")
	}

	_, span := p.tracer.Start(ctx, "owning-time-product")
	defer span.End()

	span.SetAttributes(attribute.Int("user-id", userId))
	span.SetAttributes(attribute.String("product-slug", productSlug))

	if p.paymentClient == nil {
		return time.Time{}, errors.New("payment api client not initialized")
	}

	return p.paymentClient.UserUserIdOwnsProductSlugUntilGet(ctx, payments.UserUserIdOwnsProductSlugUntilGetParams{
		UserId:      strconv.Itoa(userId),
		ProductSlug: productSlug,
	})
}
