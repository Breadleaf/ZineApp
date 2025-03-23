package redis

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Redis struct {
	Connection net.Conn
}

func NewRedis(socket string) *Redis {
	conn, err := net.Dial("tcp", socket)
	if err != nil {
		panic(err)
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	fmt.Fprintf(
		conn,
		"*2\r\n$4\r\nAUTH\r\n$%d\r\n%s\r\n",
		len(redisPassword),
		redisPassword,
	)

	reader := bufio.NewReader(conn)
	line, _ := reader.ReadString('\n')
	if !strings.HasPrefix(line, "+OK") {
		panic("redis auth failed: " + line)
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