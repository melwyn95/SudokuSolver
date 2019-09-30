package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func GetConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-17762.c57.us-east-1-4.ec2.cloud.redislabs.com:17762",
		Password: "Ok8I0W63Uv2CJKGg0VVUcssLaG7mz9Tc",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

func GenerateCacheKey(puzzle string, mode string) string {
	return mode + "_" + puzzle
}

func GetKey(connection *redis.Client, cacheKey string) (string, bool) {
	val, err := connection.Get(cacheKey).Result()
	return val, err != redis.Nil
}

func SetKey(connection *redis.Client, cacheKey string, value string) bool {
	err := connection.Set(cacheKey, value, 0).Err()
	return err != nil
}
