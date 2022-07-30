package metrics

import (
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	messagesProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "discord_coflnet_bot_messages_processed",
		Help: "Count of processed messages",
	})

	errorsOccured = promauto.NewCounter(prometheus.CounterOpts{
		Name: "discord_coflnet_bot_errors_occured",
		Help: "Count of errors occured",
	})

	userLoaded = promauto.NewCounter(prometheus.CounterOpts{
		Name: "discord_coflnet_bot_user_loaded",
		Help: "Count of users loaded",
	})

	inGameMuteTriggered = promauto.NewCounter(prometheus.CounterOpts{
		Name: "discord_coflnet_bot_in_game_mute_triggered",
		Help: "Count of in game mute triggered",
	})

	inGameUnmuteTriggered = promauto.NewCounter(prometheus.CounterOpts{
		Name: "discord_coflnet_bot_in_game_unmute_triggered",
		Help: "Count of in game unmute triggered",
	})

	flipSummaryProcessingError = promauto.NewCounter(prometheus.CounterOpts{
		Name: "discord_coflnet_bot_flip_summary_processing_error",
		Help: "Errors while processing flip summaries",
	})

	flipSummarySend = promauto.NewCounter(prometheus.CounterOpts{
		Name: "discord_coflnet_bot_flip_summary_send",
		Help: "The amount of flip summaries send",
	})
)

func Init() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		log.Panic().Err(err).Msgf("error starting metrics server")
		return
	}
}

func MessageProcessed() {
	messagesProcessed.Inc()
}

func ErrorOccurred() {
	errorsOccured.Inc()
}

func UserLoaded() {
	userLoaded.Inc()
}

func InGameMuteTriggered() {
	inGameMuteTriggered.Inc()
}

func InGameUnmuteTriggered() {
	inGameUnmuteTriggered.Inc()
}

func FlipSummaryProcessingError() {
	flipSummaryProcessingError.Inc()
}

func FlipSummarySend() {
	flipSummarySend.Inc()
}
