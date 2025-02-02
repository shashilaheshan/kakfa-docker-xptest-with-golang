package main

import (
    "context"
    "fmt"
    "log"

    "github.com/segmentio/kafka-go"
)

func main() {
    // The broker address for outside Docker is "localhost:29092"
    brokers := []string{"localhost:29092"}
    topic := "heshan-docker-xp-debug"
    groupID := "my-consumer-group"

    // Create a new reader with the brokers and topic
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:  brokers,
        Topic:    topic,
        GroupID:  groupID,          // consumer group name
        MinBytes: 10e3,             // 10KB
        MaxBytes: 10e6,             // 10MB
    })
    defer r.Close()

    log.Printf("Starting consumer. Broker: %v, Topic: %s, GroupID: %s", brokers, topic, groupID)

    // Continuously read messages
    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            log.Fatal("error reading message:", err)
        }
        fmt.Printf("Received message at offset %d: %s = %s\n",
            m.Offset, string(m.Key), string(m.Value))
    }
}