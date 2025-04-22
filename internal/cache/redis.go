package cache

import (
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	for i := 0; i < 5; i++ {
		_, err := RedisClient.Ping().Result()
		if err == nil {
			log.Println("✅ Redis connected")
			return
		}
		log.Printf("Redis not ready (%v), retrying...\n", err)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Redis error: %v", RedisClient.Ping().Err())
}
