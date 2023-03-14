package mongo

import (
	"context"
	"fmt"

	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel/attribute"
)

func UsersByDiscordId(ctx context.Context, id string) ([]*model.User, error) {
	ctx, span := tracer.Start(ctx, "user-by-discord-id")
    span.SetAttributes(attribute.String("search-discord-id", id))
	defer span.End()

    filter := bson.M{"discord_ids" : id}

    c, err := userCollection.Find(ctx, filter)
    if err != nil {
        span.RecordError(err)
        return nil, err
    }

    var users []*model.User
    for c.Next(ctx) {
        var u model.User
        err := c.Decode(&u)
        if err != nil {
            span.RecordError(err)
            return nil, err
        }
        users = append(users, &u)
    }
    span.SetAttributes(attribute.Int("found-users", len(users)))

    return users, nil
}

func UsersByDiscordTag(ctx context.Context, tag string) ([]*model.User, error) {
	ctx, span := tracer.Start(ctx, "user-by-discord-tag")
    span.SetAttributes(attribute.String("search-discord-tag", tag))
	defer span.End()

    filter := bson.M{"discord_names" : tag}

    c, err := userCollection.Find(ctx, filter)
    if err != nil {
        span.RecordError(err)
        return nil, err
    }

    var users []*model.User
    for c.Next(ctx) {
        var u model.User
        err := c.Decode(&u)
        if err != nil {
            span.RecordError(err)
            return nil, err
        }
        users = append(users, &u)
    }
    span.SetAttributes(attribute.Int("found-users", len(users)))

    return users, nil
}

// UserByDiscordAccount returns the user with the given discord account
// it searches by the discord id, if the id exists it returns the user
// if the id does not exist it searches by the discord tag
func UsersByDiscordAccount(ctx context.Context, user *discordgo.User) ([]*model.User, error) {
	ctx, span := tracer.Start(ctx, "user-by-discord-account")
	defer span.End()

    u, err := UsersByDiscordId(ctx, user.ID)
    if err == nil && u != nil {
        return u, nil
    } 
    if err != nil {
        span.RecordError(err)
    }

    tag := fmt.Sprintf("%s#%s", user.Username, user.Discriminator)
    u, err = UsersByDiscordTag(ctx, tag)

    if err != nil {
        span.RecordError(err)
    }

    return u, err
}

func UserById(ctx context.Context, id int) (*model.User, error) {
    ctx, span := tracer.Start(ctx, "user-by-id")
    defer span.End()
    span.SetAttributes(attribute.Int("search-user-id", id))

    filter := bson.M{"user_id" : id}

    var u model.User
    err := userCollection.FindOne(ctx, filter).Decode(&u)

    if err != nil {
        span.RecordError(err)
    }

    return &u, nil

}
