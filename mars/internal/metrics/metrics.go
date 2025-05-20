package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	SegmentProcessMilliseconds = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "segment_process_avg_milliseconds_gauge",
		Help: "Среднее время обработки сегмента",
	})

	RequestTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Общее количество HTTP запросов",
	})

	RequestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Длительность HTTP запросов",
		Buckets: prometheus.DefBuckets,
	})

	ErrorsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "errors_total",
		Help: "Количество ошибок по типам",
	}, []string{"type"})

	QueueSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "message_queue_size",
		Help: "Текущий размер очереди сообщений",
	})
)

func init() {
	prometheus.MustRegister(SegmentProcessMilliseconds)
	prometheus.MustRegister(RequestTotal)
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(ErrorsTotal)
	prometheus.MustRegister(QueueSize)
}
