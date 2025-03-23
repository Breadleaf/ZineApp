package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

const PORT string = ":8080"

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Payload struct {
	People []Person `json:"people"`
}

type Redis struct {
	Connection net.Conn
}

func NewRedis(socket string) *Redis {
	conn, err := net.Dial("tcp", socket)
	if err != nil {
		panic(err)
	}
	return &Redis{Connection: conn}
}

func (r *Redis) Close() {
	if err := r.Connection.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	}
}

func (r *Redis) parseReturn() (string, error) {
	reader := bufio.NewReader(r.Connection)
	header, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	switch header[0] {
	case '+': // simple string
		return strings.TrimSuffix(header[1:], "\r\n"), nil

	case '-': // error
		return "", errors.New(
			strings.TrimSuffix(header[1:], "\r\n"),
		)

	case '$': // bulk string
		var length int
		if _, err := fmt.Sscanf(header, "$%d\r\n", &length); err != nil {
			return "", err
		}
		if length == -1 {
			return "", nil // nil value in redis
		}

		buf := make([]byte, length)
		if _, err := reader.Read(buf); err != nil {
			return "", err
		}
		// discard trailing "\r\n"
		if _, err := reader.Discard(2); err != nil {
			return "", nil
		}

		return string(buf), nil

	default:  // not yet implemented reply
		return "", fmt.Errorf("unsupported reply: %s", header)
	}


}

func (r *Redis) Set(key, value string) (string, error) {
	fmt.Fprintf(
		r.Connection,
		"*3\r\n$3\r\nSET\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n",
		len(key),
		key,
		len(value),
		value,
	)
	return r.parseReturn()
}

func (r *Redis) Get(key string) (string, error) {
	fmt.Fprintf(
		r.Connection,
		"*2\r\n$3\r\nGET\r\n$%d\r\n%s\r\n",
		len(key),
		key,
	)
	return r.parseReturn()
}

func main() {
	redis := NewRedis("redis:6379")
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
			Age: 21,
		}
		person2 := Person{
			Name: "Jane",
			Age: 22,
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

	log.Println("Starting backend on port", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}
