package metrics

import (
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	DiscordMessagesProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "coflnet_bot_discord_messages_processed",
		Help: "Count of processed discord messages",
	})

    MessagesForwardedToRedisChatChannel = promauto.NewCounter(prometheus.CounterOpts{
        Name: "coflnet_bot_redis_chat_channel_messages_forwarded",
        Help: "Count of discord messages forwarded to the redis chat channel",
    })

    MessagesInsertedIntoDatabase  = promauto.NewCounter(prometheus.CounterOpts{
        Name: "coflnet_bot_messages_inserted_into_database",
        Help: "Count of discord messages inserted into the database",
    })

    MessagesForwardedToCoflnetDiscordChatChannel  = promauto.NewCounter(prometheus.CounterOpts{
        Name: "coflnet_bot_messages_forwarded_to_coflnet_discord_chat_channel",
        Help: "Count of discord messages forwarded to the coflnet discord chat channel",
    })
)

func Init() {
    slog.Info("starting metrics server")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)


	if err != nil {
		slog.Error("error starting metrics server", err)
        panic(err)
	}
}
