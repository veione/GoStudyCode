package main

import (
	"github.com/go-redis/redis"
	"fmt"
)


func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.1.1.95:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("feekey", "examples", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("feekey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("feekey", val)

	val2, err := client.Get("feekey2").Result()
	if err == redis.Nil {
		fmt.Println("feekey does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("feekey", val2)
	}
}

func main() {
	ExampleClient()
}
