// Code generated by ogen, DO NOT EDIT.

package chat

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
)

// handleAPIChatInternalClientPostRequest handles POST /api/Chat/internal/client operation.
//
// Create a nw Client.
//
// POST /api/Chat/internal/client
func (s *Server) handleAPIChatInternalClientPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	var otelAttrs []attribute.KeyValue

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIChatInternalClientPost",
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIChatInternalClientPost",
			ID:   "",
		}
	)
	request, close, err := s.decodeAPIChatInternalClientPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *APIChatInternalClientPostOKApplicationJSON
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "APIChatInternalClientPost",
			OperationID:   "",
			Body:          request,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = *APIChatInternalClientPostReqApplicationJSON
			Params   = struct{}
			Response = *APIChatInternalClientPostOKApplicationJSON
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIChatInternalClientPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIChatInternalClientPost(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIChatInternalClientPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleAPIChatMuteDeleteRequest handles DELETE /api/Chat/mute operation.
//
// Create a new mute for an user.
//
// DELETE /api/Chat/mute
func (s *Server) handleAPIChatMuteDeleteRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	var otelAttrs []attribute.KeyValue

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIChatMuteDelete",
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIChatMuteDelete",
			ID:   "",
		}
	)
	request, close, err := s.decodeAPIChatMuteDeleteRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *APIChatMuteDeleteOKApplicationJSON
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "APIChatMuteDelete",
			OperationID:   "",
			Body:          request,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = *APIChatMuteDeleteReqApplicationJSON
			Params   = struct{}
			Response = *APIChatMuteDeleteOKApplicationJSON
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIChatMuteDelete(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIChatMuteDelete(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIChatMuteDeleteResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleAPIChatMutePostRequest handles POST /api/Chat/mute operation.
//
// Create a new mute for an user.
//
// POST /api/Chat/mute
func (s *Server) handleAPIChatMutePostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	var otelAttrs []attribute.KeyValue

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIChatMutePost",
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIChatMutePost",
			ID:   "",
		}
	)
	request, close, err := s.decodeAPIChatMutePostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *APIChatMutePostOKApplicationJSON
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "APIChatMutePost",
			OperationID:   "",
			Body:          request,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = *APIChatMutePostReqApplicationJSON
			Params   = struct{}
			Response = *APIChatMutePostOKApplicationJSON
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIChatMutePost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIChatMutePost(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIChatMutePostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleAPIChatSendPostRequest handles POST /api/Chat/send operation.
//
// Tracks a flip.
//
// POST /api/Chat/send
func (s *Server) handleAPIChatSendPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	var otelAttrs []attribute.KeyValue

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIChatSendPost",
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "APIChatSendPost",
			ID:   "",
		}
	)
	request, close, err := s.decodeAPIChatSendPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *APIChatSendPostOKApplicationJSON
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "APIChatSendPost",
			OperationID:   "",
			Body:          request,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = *APIChatSendPostReqApplicationJSON
			Params   = struct{}
			Response = *APIChatSendPostOKApplicationJSON
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIChatSendPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIChatSendPost(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIChatSendPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}