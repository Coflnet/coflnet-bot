// Code generated by ogen, DO NOT EDIT.

package chat

import (
	"context"
	"net/url"
	"time"
    "os"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// APIChatInternalClientPost invokes POST /api/Chat/internal/client operation.
//
// Create a nw Client.
//
// POST /api/Chat/internal/client
func (c *Client) APIChatInternalClientPost(ctx context.Context, request APIChatInternalClientPostReq) (APIChatInternalClientPostRes, error) {
	res, err := c.sendAPIChatInternalClientPost(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendAPIChatInternalClientPost(ctx context.Context, request APIChatInternalClientPostReq) (res APIChatInternalClientPostRes, err error) {
	var otelAttrs []attribute.KeyValue
	// Validate request before sending.
	switch request := request.(type) {
	case *APIChatInternalClientPostReqEmptyBody:
		// Validation is not needed for the empty body type.
	case *APIChatInternalClientPostApplicationJSON:
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	case *APIChatInternalClientPostTextJSON:
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	default:
		return res, errors.Errorf("unexpected request type: %T", request)
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "APIChatInternalClientPost",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/api/Chat/internal/client"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAPIChatInternalClientPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAPIChatInternalClientPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APIChatMuteDelete invokes DELETE /api/Chat/mute operation.
//
// Create a new mute for an user.
//
// DELETE /api/Chat/mute
func (c *Client) APIChatMuteDelete(ctx context.Context, request APIChatMuteDeleteReq) (APIChatMuteDeleteRes, error) {
	res, err := c.sendAPIChatMuteDelete(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendAPIChatMuteDelete(ctx context.Context, request APIChatMuteDeleteReq) (res APIChatMuteDeleteRes, err error) {
	var otelAttrs []attribute.KeyValue
	// Validate request before sending.
	switch request := request.(type) {
	case *APIChatMuteDeleteReqEmptyBody:
		// Validation is not needed for the empty body type.
	case *APIChatMuteDeleteApplicationJSON:
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	case *APIChatMuteDeleteTextJSON:
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	default:
		return res, errors.Errorf("unexpected request type: %T", request)
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "APIChatMuteDelete",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/api/Chat/mute"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAPIChatMuteDeleteRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAPIChatMuteDeleteResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APIChatMutePost invokes POST /api/Chat/mute operation.
//
// Create a new mute for an user.
//
// POST /api/Chat/mute
func (c *Client) APIChatMutePost(ctx context.Context, request APIChatMutePostReq) (APIChatMutePostRes, error) {
	res, err := c.sendAPIChatMutePost(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendAPIChatMutePost(ctx context.Context, request APIChatMutePostReq) (res APIChatMutePostRes, err error) {
	var otelAttrs []attribute.KeyValue
	// Validate request before sending.
	switch request := request.(type) {
	case *APIChatMutePostReqEmptyBody:
		// Validation is not needed for the empty body type.
	case *APIChatMutePostApplicationJSON:
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	case *APIChatMutePostTextJSON:
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	default:
		return res, errors.Errorf("unexpected request type: %T", request)
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "APIChatMutePost",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/api/Chat/mute"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

    r.Header.Set("authorization", os.Getenv("COFL_CHAT_API_KEY"))

	if err := encodeAPIChatMutePostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAPIChatMutePostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// APIChatSendPost invokes POST /api/Chat/send operation.
//
// Tracks a flip.
//
// POST /api/Chat/send
func (c *Client) APIChatSendPost(ctx context.Context, request APIChatSendPostReq) (APIChatSendPostRes, error) {
	res, err := c.sendAPIChatSendPost(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendAPIChatSendPost(ctx context.Context, request APIChatSendPostReq) (res APIChatSendPostRes, err error) {
	var otelAttrs []attribute.KeyValue
	// Validate request before sending.
	switch request := request.(type) {
	case *APIChatSendPostReqEmptyBody:
		// Validation is not needed for the empty body type.
	case *APIChatSendPostApplicationJSON:
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	case *APIChatSendPostTextJSON:
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	default:
		return res, errors.Errorf("unexpected request type: %T", request)
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "APIChatSendPost",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/api/Chat/send"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

    r.Header.Set("authorization", os.Getenv("COFL_CHAT_API_KEY"))

	if err := encodeAPIChatSendPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAPIChatSendPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
