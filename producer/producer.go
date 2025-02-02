package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/segmentio/kafka-go"
)

func main() {
    // Kafka broker address as advertised for outside Docker
    brokerAddress := "localhost:29092"
    topic := "heshan-docker-xp-debug"

    // Create a new writer with the broker address and the topic
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  []string{brokerAddress},
        Topic:    topic,
        Balancer: &kafka.LeastBytes{}, // which partitioner? LeastBytes is a basic round-robin approach
    })
    defer writer.Close()

    log.Printf("Starting producer. Broker: %s, Topic: %s", brokerAddress, topic)

    // Send 10 messages, one per second
    for i := 0; i < 1000; i++ {
        // Construct the message
        messageValue := fmt.Sprintf("Hello Kafka %d", i)
        err := writer.WriteMessages(context.Background(),
            kafka.Message{
                Key:   []byte(fmt.Sprintf("Key-%d", i)),
                Value: []byte(messageValue),
            },
        )
        if err != nil {
            log.Fatal("Failed to write message:", err)
        }

        log.Printf("Sent message: %s", messageValue)
        time.Sleep(100 * time.Millisecond)
    }

    log.Println("Finished producing messages!")
}