package main

import (
	"coflnet-bot/internal/db"
	chatgen "coflnet-bot/internal/gen/chat"
	"coflnet-bot/internal/server"
	"coflnet-bot/internal/usecase"
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	othertrace "go.opentelemetry.io/otel/trace"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	const service = "coflnet-bot"
	const version = "0.0.1"

	// setup logging lib
	setupLogger()

	// setup open telemetry
	shutdown, err := setupOTelSDK(ctx, service, version)
	if err != nil {
		slog.Error("failed to initialize OTel SDK", "err", err)
		os.Exit(1)
	}
	defer func(ctx context.Context) {
		err = shutdown(ctx)
		if err != nil {
			slog.Error("failed to shutdown OTel SDK", "err", err)
		}
	}(ctx)

	// setup database
	err = db.StartDB(ctx)
	if err != nil {
		slog.Error("failed to start db", "err", err)
		panic("cannot start db")
	}

	// setup openapi clients
	chatClient, err := chatgen.NewClient(chatUrl())
	if err != nil {
		panic("cannot create chat client")
	}

	// setup some base dependencies
	discordMessageService := usecase.NewDiscordMessageService()
	redisMessageService := usecase.NewRedisMessageService()
	userService, err := usecase.NewUserService(paymentUrl(), proxyUrl(), mcConnectUrl())
	if err != nil {
		slog.Error("error setting up user service", "err", err)
		panic(errors.New("cannot setup user service"))
	}

	// setup background services
	chat, err := usecase.NewChat(discordMessageService, redisMessageService, chatClient, userService)
	if err != nil {
		slog.Error("error setting up chat", "err", err)
		panic(errors.New("cannot setup chat"))
	}

	go func() {
		err := chat.StartChatService(ctx)
		if err != nil {
			slog.Error("failed to start chat service", "err", err)
			panic(err)
		}
	}()

	// start mongo migration
	if migrationEnabled := os.Getenv("MIGRATION_ENABLED"); migrationEnabled == "true" {
		go func() {
			mongoMigrationService := usecase.NewMongoMigrationService(userService)
			err = mongoMigrationService.Migrate(ctx)
			if err != nil {
				slog.Error("failed to migrate mongo", "err", err)
			}
		}()
	}

	startApi()
}

func startApi() {
	apiServer := server.NewServer()

	err := apiServer.ListenAndServe()
	if err != nil {
		slog.Error("failed to start apiServer", "err", err)
		panic("cannot start apiServer")
	}
}

func proxyUrl() string {
	return mustLoadEnv("PROXY_URL")
}

func chatUrl() string {
	return mustLoadEnv("CHAT_URL")
}

func paymentUrl() string {
	return mustLoadEnv("PAYMENT_URL")
}

func mcConnectUrl() string {
	return mustLoadEnv("MC_CONNECT_URL")
}

func mustLoadEnv(key string) string {
	val, found := os.LookupEnv(key)
	if !found {
		panic(errors.New(fmt.Sprintf("missing env var %s", key)))
	}
	return val
}

func setupLogger() {
	var level slog.Level
	if true {
		level = slog.LevelDebug
	}

	var handler slog.Handler
	handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	})
	handler = TraceLogHandler{handler}
	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.Info(fmt.Sprintf("using log level %s", level.String()))
}

type TraceLogHandler struct {
	slog.Handler
}

func (h TraceLogHandler) Handle(ctx context.Context, r slog.Record) error {
	span := othertrace.SpanFromContext(ctx)
	if span == nil {
		return h.Handler.Handle(ctx, r)
	}
	if !span.SpanContext().IsValid() {
		return h.Handler.Handle(ctx, r)
	}
	traceId := span.SpanContext().TraceID().String()
	r.Add("trace", slog.StringValue(traceId))
	return h.Handler.Handle(ctx, r)
}

// OpenTelemetry initialization.
// setupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func setupOTelSDK(ctx context.Context, serviceName, serviceVersion string) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
			if err != nil {
				slog.Error("error during shutdown", "err", err)
			}
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		slog.Error("error during setup", "err", inErr)
		err = errors.Join(inErr, shutdown(ctx))
	}

	// Set up resource.
	res, err := newResource(serviceName, serviceVersion)
	if err != nil {
		handleErr(err)
		return
	}

	// Set up propagator.
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// Set up trace provider.
	tracerProvider, err := newTraceProvider(ctx, res)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	// Set up meter provider.
	meterProvider, err := newMeterProvider(res)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)
	go serveMetrics()

	return
}

func newResource(serviceName, serviceVersion string) (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		))
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTraceProvider(ctx context.Context, res *resource.Resource) (*trace.TracerProvider, error) {
	traceExporter, err := otlptracegrpc.New(ctx)

	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithResource(res),
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(traceExporter),
	)
	return traceProvider, nil
}

func newMeterProvider(res *resource.Resource) (*metric.MeterProvider, error) {
	prometheusExporter, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(prometheusExporter),
	)
	return meterProvider, nil
}

func serveMetrics() {
	slog.Info("serving metrics at localhost:2222/metrics")
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":2222", nil)
	if err != nil {
		slog.Error("error serving http", "err", err)
		return
	}
}
