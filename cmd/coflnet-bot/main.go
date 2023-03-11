package main

import (
	"os"
	"os/signal"

	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/Coflnet/coflnet-bot/internal/usecase"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"golang.org/x/exp/slog"
)

func main() {
	// setup tracer
	setupTracer()

    // setup logger
    h := slog.HandlerOptions{Level: slog.LevelInfo}.NewTextHandler(os.Stdout)
    if utils.DebugEnabled() {
        h = slog.HandlerOptions{Level: slog.LevelDebug}.NewTextHandler(os.Stdout)
    }
    slog.SetDefault(slog.New(h))

	// metrics
	go metrics.Init()

	// setup database connection
	err := mongo.Init()
	if err != nil {
		slog.Error("error connecting to database", err)
        panic(err)
	}
	defer mongo.CloseConnection()

	// start the kafka message processors
	go startMessageProcessors()

    // register discord commands
    discordHandler, err := discord.NewDiscordHandler()
    if err != nil {
        panic(err)
    }
    err = discordHandler.RegisterCommands()
    if err != nil {
        panic(err)
    }


	// wait for interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	slog.Info("shutting down..")
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

func startMessageProcessors() {
	processors := []usecase.MessageProcessor{
		new(usecase.FlipProcessor),
        new(usecase.ChatProcessor),
        new(usecase.McVerifyProcessor),
        
        // new(usecase.DiscordMessageProcessor),
	}

	for _, p := range processors {
		go func(p usecase.MessageProcessor) {
			err := p.StartProcessing()
            slog.Error("error in message processor", err)
            panic(err)
		}(p)
	}
}
