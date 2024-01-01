package usecase

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"os"
	"time"
)

type MongoMigrationService struct {
	tracer      trace.Tracer
	client      *mongo.Client
	userService *UserService
}

func NewMongoMigrationService(userService *UserService) *MongoMigrationService {
	return &MongoMigrationService{
		tracer:      otel.Tracer("mongo-migration-service"),
		userService: userService,
	}
}

func (s *MongoMigrationService) Migrate(ctx context.Context) error {
	ctx, span := s.tracer.Start(ctx, "migrate")
	defer span.End()

	err := s.connectToDatabase(ctx)
	if err != nil {
		slog.Error("Error connecting to database", "err", err)
		span.RecordError(err)
		return err
	}

	slog.Info("Starting Migrations")
	coll := s.client.Database("discord").Collection("users")

	opts := options.Find()
	opts.SetSort(bson.D{{"_id", -1}})

	cur, err := coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		slog.Error("Error finding messages", "err", err)
		span.RecordError(err)
		return err
	}

	for cur.Next(ctx) {
		var result bson.M
		err = cur.Decode(&result)
		if err != nil {
			slog.Warn("Error decoding message", "err", err)
			continue
		}

		err = s.migrateDocument(ctx, result)
		if err != nil {
			slog.Warn("Error migrating document", "err", err)
			continue
		}
	}

	slog.Info("Finished Migrations")
	return nil
}

func (s *MongoMigrationService) migrateDocument(ctx context.Context, doc bson.M) error {
	ctx, span := s.tracer.Start(context.Background(), "migrate-document")
	defer span.End()

	uuids := doc["minecraft_uuids"]

	if uuids == nil {
		return nil
	}

	for _, uuid := range uuids.(primitive.A) {
		slog.Debug("Migrating uuid", "uuid", uuid)

		user, err := s.userService.LoadUserByUUID(ctx, uuid.(string))
		if err != nil {
			slog.Error("Error loading user", "err", err)
			span.RecordError(err)
			return err
		}

		slog.Info(fmt.Sprint("Migrated user with external id %s", user.ExternalId))
	}
	time.Sleep(10 * time.Second)

	return nil
}

func (s *MongoMigrationService) connectToDatabase(ctx context.Context) error {
	ctx, span := s.tracer.Start(ctx, "connect-to-database")
	defer span.End()
	slog.Info("Connecting to mongo database")

	mclient, err := mongo.Connect(ctx, options.Client().ApplyURI(s.mongoURI()))
	if err != nil {
		return err
	}

	s.client = mclient
	return nil
}

func (s *MongoMigrationService) mongoURI() string {
	val, found := os.LookupEnv("MONGO_URI")
	if !found {
		panic("MONGO_URI not found")
	}
	return val
}
