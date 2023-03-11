package coflnet

import (
	"context"
	"errors"
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

type PaymentApi struct {
    tracer trace.Tracer
    paymentClient *payments.Client
}

func NewPaymentApi() (*PaymentApi, error) {
    var err error
    r := &PaymentApi{}

    r.paymentClient, err = payments.NewClient(utils.PaymentBaseUrl())
    r.tracer = otel.Tracer(PaymentTraceName)

    return r, err
}


// TODO make this fancier
// load all products and map them to colors/roles
func (p *PaymentApi) OwningTimesOfUser(ctx context.Context, userId int) ([]model.OwnedProducts, error) {
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
                ExpiresAt: t,
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
    _, span := p.tracer.Start(ctx, "owning-time-product") 
    defer span.End()

    span.SetAttributes(attribute.Int("user-id", userId))
    span.SetAttributes(attribute.String("product-slug", productSlug))

    if p.paymentClient == nil {
        return time.Time{}, errors.New("payment api client not initialized")
    }

    return p.paymentClient.UserUserIdOwnsProductSlugUntilGet(ctx, payments.UserUserIdOwnsProductSlugUntilGetParams{
        UserId: strconv.Itoa(userId),
        ProductSlug: productSlug,
    })
}
