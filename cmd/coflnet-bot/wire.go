//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/db"
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/processor"
	"github.com/Coflnet/coflnet-bot/internal/redis"
	"github.com/Coflnet/coflnet-bot/internal/usecase"
	"github.com/google/wire"
)

func wireApp() *App {
	panic(wire.Build(
		wire.NewSet(db.NewDB, db.NewUserRepo),
		wire.NewSet(discord.NewMuteCommand, discord.NewUnmuteCommand),
		wire.NewSet(coflnet.NewMcConnectApi, coflnet.NewPaymentApi, coflnet.NewChatApi, coflnet.NewApiClient),
		wire.NewSet(discord.NewDiscordHandler, redis.NewRedisHandler),
		wire.NewSet(usecase.NewUserHandler),
		wire.NewSet(processor.NewChatProcessor, processor.NewDiscordMessageProcessor),
		newApp,
	),
	)
}
