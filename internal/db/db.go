package db

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
)

var (
	db     *gorm.DB
	tracer trace.Tracer
)

func StartDB(ctx context.Context) error {
	tracer = otel.Tracer("db")
	ctx, span := tracer.Start(ctx, "start-db")
	defer span.End()

	var err error
	db, err = gorm.Open(postgres.Open(LoadDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	err = migrations(ctx)
	if err != nil {
		slog.Error("Error running migrations", "err", err)
		span.RecordError(err)
		return err
	}

	slog.Info("Starting DB")
	return nil
}

func migrations(ctx context.Context) error {
	ctx, span := tracer.Start(ctx, "migrations")
	defer span.End()

	slog.Info("Starting Migrations")

	err := db.AutoMigrate(&Message{})
	if err != nil {
		return err
	}

	// err = db.AutoMigrate(&User{})
	// if err != nil {
	// 	return err
	// }

	err = db.AutoMigrate(&DiscordAccount{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&MinecraftAccount{})
	if err != nil {
		return err
	}

	return nil
}

func LoadDSN() string {
	dsn, found := os.LookupEnv("DATABASE_URL")

	if !found {
		panic("DATABASE_URL not found")
	}

	slog.Info("DATABASE_URL found", "dsn", len(dsn))
	return dsn
}
