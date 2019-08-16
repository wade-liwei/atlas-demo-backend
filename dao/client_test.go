package dao

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

type Student struct {
	Name string
	Age  int
}

func TestSetGet(t *testing.T) {

	s := Student{
		Name: "wade",
		Age:  100,
	}

	jsonByte, err := json.Marshal(s)

	if err != nil {
		t.Error(err)
	}

	err = client.Set("wade3", jsonByte, 300*time.Second).Err()
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	val, err := client.Get("wade3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
