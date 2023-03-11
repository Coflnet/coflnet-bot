// Code generated by ogen, DO NOT EDIT.

package mc_connect

import (
	"context"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/ogen-go/ogen/conv"
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

// ConnectMinecraftMcUuidGet invokes GET /Connect/minecraft/{mcUuid} operation.
//
// GET /Connect/minecraft/{mcUuid}
func (c *Client) ConnectMinecraftMcUuidGet(ctx context.Context, params ConnectMinecraftMcUuidGetParams) (*User, error) {
	res, err := c.sendConnectMinecraftMcUuidGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendConnectMinecraftMcUuidGet(ctx context.Context, params ConnectMinecraftMcUuidGetParams) (res *User, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ConnectMinecraftMcUuidGet",
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
	u.Path += "/Connect/minecraft/"
	{
		// Encode "mcUuid" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "mcUuid",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.McUuid))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeConnectMinecraftMcUuidGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ConnectUserUserIdGet invokes GET /Connect/user/{userId} operation.
//
// GET /Connect/user/{userId}
func (c *Client) ConnectUserUserIdGet(ctx context.Context, params ConnectUserUserIdGetParams) (*User, error) {
	res, err := c.sendConnectUserUserIdGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendConnectUserUserIdGet(ctx context.Context, params ConnectUserUserIdGetParams) (res *User, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ConnectUserUserIdGet",
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
	u.Path += "/Connect/user/"
	{
		// Encode "userId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "userId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.UserId))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeConnectUserUserIdGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ConnectUserUserIdPost invokes POST /Connect/user/{userId} operation.
//
// POST /Connect/user/{userId}
func (c *Client) ConnectUserUserIdPost(ctx context.Context, params ConnectUserUserIdPostParams) (*ConnectionRequest, error) {
	res, err := c.sendConnectUserUserIdPost(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendConnectUserUserIdPost(ctx context.Context, params ConnectUserUserIdPostParams) (res *ConnectionRequest, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ConnectUserUserIdPost",
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
	u.Path += "/Connect/user/"
	{
		// Encode "userId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "userId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.UserId))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "mcUuid" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "mcUuid",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.McUuid.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeConnectUserUserIdPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ConnectUserUserIdVerifyPost invokes POST /Connect/user/{userId}/verify operation.
//
// POST /Connect/user/{userId}/verify
func (c *Client) ConnectUserUserIdVerifyPost(ctx context.Context, params ConnectUserUserIdVerifyPostParams) error {
	res, err := c.sendConnectUserUserIdVerifyPost(ctx, params)
	_ = res
	return err
}

func (c *Client) sendConnectUserUserIdVerifyPost(ctx context.Context, params ConnectUserUserIdVerifyPostParams) (res *ConnectUserUserIdVerifyPostOK, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ConnectUserUserIdVerifyPost",
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
	u.Path += "/Connect/user/"
	{
		// Encode "userId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "userId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.UserId))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "/verify"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "mcUuid" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "mcUuid",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.McUuid.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeConnectUserUserIdVerifyPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ConnectUsersConnectedGet invokes GET /Connect/users/connected operation.
//
// Get all users stored which may or may not have a connected account.
//
// GET /Connect/users/connected
func (c *Client) ConnectUsersConnectedGet(ctx context.Context, params ConnectUsersConnectedGetParams) ([]User, error) {
	res, err := c.sendConnectUsersConnectedGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendConnectUsersConnectedGet(ctx context.Context, params ConnectUsersConnectedGetParams) (res []User, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ConnectUsersConnectedGet",
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
	u.Path += "/Connect/users/connected"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "amount" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "amount",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Amount.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "offset" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "offset",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Offset.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeConnectUsersConnectedGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ConnectUsersGet invokes GET /Connect/users operation.
//
// Get all users stored which may or may not have a connected account.
//
// GET /Connect/users
func (c *Client) ConnectUsersGet(ctx context.Context, params ConnectUsersGetParams) ([]User, error) {
	res, err := c.sendConnectUsersGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendConnectUsersGet(ctx context.Context, params ConnectUsersGetParams) (res []User, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ConnectUsersGet",
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
	u.Path += "/Connect/users"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "amount" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "amount",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Amount.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "offset" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "offset",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Offset.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeConnectUsersGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ConnectUsersIdsGet invokes GET /Connect/users/ids operation.
//
// Get all users which ids is in the list.
//
// GET /Connect/users/ids
func (c *Client) ConnectUsersIdsGet(ctx context.Context, params ConnectUsersIdsGetParams) ([]User, error) {
	res, err := c.sendConnectUsersIdsGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendConnectUsersIdsGet(ctx context.Context, params ConnectUsersIdsGetParams) (res []User, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ConnectUsersIdsGet",
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
	u.Path += "/Connect/users/ids"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "externalIds" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "externalIds",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range params.ExternalIds {
					if err := func() error {
						return e.EncodeValue(conv.StringToString(item))
					}(); err != nil {
						return errors.Wrapf(err, "[%d]", i)
					}
				}
				return nil
			})
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeConnectUsersIdsGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}