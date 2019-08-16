package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

func init() {

	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("pong:  %v \n", pong)
}

func Set(key string, value []byte) error {
	return client.Set(key, value, 6*time.Minute).Err()
}

func Get(key string) ([]byte, error) {
	return client.Get(key).Bytes()
}
