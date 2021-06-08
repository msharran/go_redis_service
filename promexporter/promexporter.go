package promexporter

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var totalKeysInDB = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "goredis_keys_in_redis_total",
		Help: "Number of keys in redis.",
	},
	[]string{"key"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "goredis_response_status",
		Help: "Status of HTTP response",
	},
	[]string{"path", "status"},
)

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "goredis_each_endpoint_latency_seconds_per_request",
	Help: "Latency of each endpoint",
}, []string{"time", "path"})

func init() {
	prometheus.Register(totalKeysInDB)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)
}

func PrometheusMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			timer := ExportLatencyForEachEndpoint(c.Request().URL.Path)
			h := next(c)
			timer.ObserveDuration()
			ExportStatusCodeForEachEndpoint(c.Request().URL.Path, c.Response().Status)
			ExportTotalKeyCountInDB()
			return h
		}
	}
}
