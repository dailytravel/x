package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var (
	Redis *redis.Client
)

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis")
		log.Fatal(err)
	}

	return rdb
}

func SetEmail(ctx context.Context, email string) (string, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	key := uuid.New().String()

	err := client.Set(ctx, key, email, time.Minute*5).Err()
	if err != nil {
		return "", err
	}

	return key, nil
}

func GetEmail(ctx context.Context, code string) (string, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	email, err := client.Get(ctx, code).Result()
	if err == redis.Nil {
		return "", errors.New("verification code not found or expired")
	} else if err != nil {
		return "", err
	}

	return email, nil
}
