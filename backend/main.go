package main

import (
	"log"
	"net/http"
)

const PORT string = ":8080"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Payload struct {
	People []Person `json:"people"`
}

func main() {
	// register all http.HandleFunc to the server
	setupRoutes()

	log.Println("Starting backend on port", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}
