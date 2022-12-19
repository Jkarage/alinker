package utils

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jkarage/alinker/env"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

// Initializing the store service and return a store pointer
func InitializeStore() *StorageService {
	address, err := env.Env("REDIS_ADDRESS", "localhost:6379")
	if err != nil {
		panic(err)
	}
	password, err := env.Env("REDIS_PASSWORD", "")
	if err != nil {
		panic(err)
	}
	dbNumber, err := env.Env("REDIS_DB", "0")
	if err != nil {
		panic(err)
	}
	db, err := strconv.Atoi(dbNumber)
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

/* We want to be able to save the mapping between the originalUrl
and the generated shortUrl url
*/

func SaveUrlMapping(shortUrl string, originalUrl string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

/*
We should be able to retrieve the initial long URL once the short
is provided. This is when users will be calling the shortlink in the
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
