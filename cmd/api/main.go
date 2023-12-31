package main

import (
	chatgen "coflnet-bot/gen/chat"
	"coflnet-bot/internal/db"
	"coflnet-bot/internal/server"
	"coflnet-bot/internal/usecase"
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"log/slog"
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
	chat := usecase.NewChat(discordMessageService, redisMessageService, chatClient, userService)
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

	opts := slog.HandlerOptions{
		Level: level,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	slog.SetDefault(logger)
	slog.Info(fmt.Sprintf("using log level %s", level.String()))
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
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
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
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(metricExporter)), // Default is 1m. Set to 3s for demonstrative purposes.
	)
	return meterProvider, nil
}