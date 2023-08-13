package main

import (
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/api"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/Coflnet/coflnet-bot/internal/processor"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log/slog"
	"os"
	"os/signal"
)

type App struct {
	chatProcessor           *processor.ChatProcessor
	discordMessageProcessor *processor.DiscordMessageProcessor
	apiController           *api.ApiController
}

func newApp(chatProcessor *processor.ChatProcessor, discordMessageProcessor *processor.DiscordMessageProcessor, apiController *api.ApiController) *App {
	return &App{
		chatProcessor:           chatProcessor,
		discordMessageProcessor: discordMessageProcessor,
		apiController:           apiController,
	}
}

func main() {
	setupTracer()
	setupLogger()

	// metrics
	go metrics.Init()

	// setup database connection
	err := mongo.Init()
	if err != nil {
		slog.Error("error connecting to database", err)
		panic(err)
	}
	defer mongo.CloseConnection()

	app := wireApp()
	go app.startMessageProcessors()

	go func() {
		err = app.apiController.Start()
		if err != nil {
			panic(err)
		}
		slog.Info("api controller ended")
	}()

	slog.Info("waiting for interrupt")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	slog.Info("shutting down..")
}

func setupLogger() {
	var level slog.Level
	if utils.DebugEnabled() {
		level = slog.LevelDebug
	} else {
		level = slog.LevelInfo
	}
	opts := slog.HandlerOptions{
		Level: level,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	slog.SetDefault(logger)
	slog.Info(fmt.Sprintf("starting coflnet-bot with log level %s", level.String()))
}

func setupTracer() {
	tp, err := tracerProvider()
	if err != nil {
		slog.Error("failed to create tracer provider", err)
		panic(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)
}

func tracerProvider() (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithAgentEndpoint())
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("coflnet-bot"),
			semconv.ServiceVersionKey.String("v0.0.1"),
		)),
	)
	return tp, nil
}

func (a *App) startMessageProcessors() {

	processors := []processor.MessageProcessor{
		a.chatProcessor,
		a.discordMessageProcessor,
	}

	for _, p := range processors {
		err := p.StartProcessing()
		if err != nil {
			slog.Error("failed to start message processor", err)
		}
	}
}
