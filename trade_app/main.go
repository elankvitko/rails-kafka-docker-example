package main

import (
	"context"
	"fmt"
	"os"

	kafka "github.com/segmentio/kafka-go"
)

func getKafkaWriter(kafkaURL string) *kafka.Writer {
  return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    "rails_auth",
	})
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func main() {
	// get kafka reader using environment variables.
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")
	groupID := os.Getenv("groupID")

  fmt.Printf( "URL is - %s", kafkaURL )
	reader := getKafkaReader(kafkaURL, topic, groupID)
  writer := getKafkaWriter(kafkaURL)
	defer reader.Close()
  defer writer.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
		}
    // Do what needs to be done with the message you receive from rails and then write back:
    writer.WriteMessages(context.Background(),
    	kafka.Message{
    		Key:   []byte("Status"),
        Value: []byte("200"),
    	},
    )
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
