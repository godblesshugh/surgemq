package main

import (
	"fmt"
	"gopkg.in/redis.v5"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Errorf("Ping redis error %q", err)
	}
	val1, err := client.HGetAll("FeelingConfig").Result()
	if err != nil {
		fmt.Errorf("redis/main: %v", err)
	}
	for k, v := range val1 {
		fmt.Printf("key: %q \n", k)
		fmt.Printf("value: %q", v)
	}
}
