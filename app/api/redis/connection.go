package redis

import (
	"fmt"
	"github.com/dikhimartin/redis"
)

func Connection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pong)

	return client
}
