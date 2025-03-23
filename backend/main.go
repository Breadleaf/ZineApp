package main

import (
	"backend/internal/redis"
	"fmt"
	"log"
	"net/http"
)

const PORT string = ":8080"

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Payload struct {
	People []Person `json:"people"`
}

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

	// register all http.HandleFunc to the server
	setupRoutes()

	log.Println("Starting backend on port", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}
