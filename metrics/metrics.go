package metrics

import (
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
)

func Init() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func MessageProcessed() {
	messagesProcessed.Inc()
}
