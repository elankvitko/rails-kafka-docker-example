package main

import (
	"fmt"
	"time"
	"context"
	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "my-topic"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)

	conn.SetReadDeadline(time.Now().Add(10*time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
	    _, err := batch.Read(b)
	    if err != nil {
	        break
	    }
	    fmt.Println(string(b))
	}

	batch.Close()
	conn.Close()
}
