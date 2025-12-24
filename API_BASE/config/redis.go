package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func ConnectRedis(cfg *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Khong the ket noi Redis", err)
	}
	log.Println("Ket noi Redis thanh cong")
	return rdb
}
