package main

import (
	"authentication/internal/redis"
	"fmt"
)

func main() {
	redis := redis.NewRedis("redis:6379")
	defer redis.Close()

	if val, err := redis.Set("name", "name@gmail.com"); err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("SET reply:", val)
	}

	if val, err := redis.Get("name"); err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("GET reply:", val)
	}
}
