package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func GetConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "ec2-13-233-126-100.ap-south-1.compute.amazonaws.com:6379",
		Password: "foobared",
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
