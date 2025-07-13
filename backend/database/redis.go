package database

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func ConnectRedis() {
	redisURI := os.Getenv("REDIS_URI")
	if redisURI == "" {
		log.Println("REDIS_URI is not set, skipping Redis connection.")
		return
	}

	opt, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Fatalf("Failed to parse Redis URI: %v", err)
	}

	// The `rediss` scheme implies TLS. The go-redis library handles this automatically
	// if the URI scheme is "rediss". For explicit clarity or custom TLS configs,
	// you could add a TLS config to opt.TLSConfig.
	if strings.HasPrefix(redisURI, "rediss://") {
		opt.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	RDB = redis.NewClient(opt)

	ctx := context.Background()
	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis.")
}