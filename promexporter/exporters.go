package promexporter

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/prometheus/client_golang/prometheus"
)

func ExportTotalKeyCountInDB() {
	db := newRedisClient()
	defer db.Close()
	keysCount := len(db.Keys("*").Val())
	totalKeysInDB.Reset()
	totalKeysInDB.WithLabelValues("count").Add(float64(keysCount))
}

func ExportStatusCodeForEachEndpoint(path string, status int) {
	responseStatus.WithLabelValues(path, strconv.Itoa(status)).Inc()
}

func ExportLatencyForEachEndpoint(path string) *prometheus.Timer {
	return prometheus.NewTimer(httpDuration.WithLabelValues(time.Now().String(), path))
}

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
