package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

func CacheMiddleware(rdb *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			key := r.URL.Query().Get("city")

			cachedValue, err := rdb.Get(ctx, key).Result()

			if err == nil {
				fmt.Println("cache-hit")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(cachedValue))

				return
			} else if err != redis.Nil {
				fmt.Printf("Error in retrieving the cache value %s", err)
				return
			}

			bodyBytes, err := io.ReadAll(r.Body)
			defer r.Body.Close()
			recorder := &responseRecorder{w, http.StatusOK, bodyBytes}
			next.ServeHTTP(recorder, r)

			if recorder.status == http.StatusOK {
				fmt.Println("cache-miss")

				respBody := recorder.body

				var response interface{}

				json.Unmarshal(respBody, &response)

				jsonValue, _ := json.Marshal(response)

				err := rdb.Set(ctx, key, jsonValue, 12*time.Hour).Err()

				if err != nil {
					fmt.Println("Error caching data:", err)
				}
			}
		})
	}

}

type responseRecorder struct {
	http.ResponseWriter
	status int
	body   []byte
}

func (rec *responseRecorder) WriteHeader(status int) {
	rec.status = status
	rec.ResponseWriter.WriteHeader(status)
}

func (rec *responseRecorder) Write(b []byte) (int, error) {
	rec.body = b
	return rec.ResponseWriter.Write(b)
}
