package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func setupRoutes() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"Received request: %s %s from %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)

		startTime := time.Now()

		person1 := Person{
			Name: "John",
			Age:  21,
		}
		person2 := Person{
			Name: "Jane",
			Age:  22,
		}
		payload := Payload{
			People: make([]Person, 0),
		}
		payload.People = append(payload.People, person1)
		payload.People = append(payload.People, person2)

		jsonResp, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)

		duration := time.Since(startTime)
		log.Printf("Request processed in %v", duration)
	})
}
