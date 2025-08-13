package main

import (
	"fmt"
	"sync"
	"time"

	"producer-consumer/internal/config"
	"producer-consumer/internal/consumer"
	"producer-consumer/internal/db"
	"producer-consumer/internal/model"
	"producer-consumer/internal/producer"
)

func main() {
	cfg := config.Load()

	database := db.ConnectDatabase(cfg.DSN)
	defer database.Close()

	messageCh := make(chan model.Message, cfg.RPS*cfg.Producers*int(cfg.Duration.Seconds()))
	var produced int64
	var consumed int64
	start := time.Now()

	var prodWG, consWG sync.WaitGroup

	fmt.Printf("[%v] 🚀 Starting application...\n", time.Now().Format(time.RFC3339))

	// Start consumers
	fmt.Printf("[%v] 🏁 Starting consumers...\n", time.Now().Format(time.RFC3339))
	consumer.StartConsumers(database, cfg.Consumers, messageCh, &consWG, &consumed)

	// Start producers
	fmt.Printf("[%v] 🏁 Starting producers...\n", time.Now().Format(time.RFC3339))
	producer.StartProducers(cfg.Producers, cfg.RPS, cfg.Duration, messageCh, &produced, &prodWG)

	// Wait for all producers to finish
	prodWG.Wait()

	fmt.Printf("[%v] ✅ All producers finished, closing channel...\n", time.Now().Format(time.RFC3339))
	close(messageCh)

	// Wait for all consumers to finish
	consWG.Wait()

	totalTime := time.Since(start)
	fmt.Printf("\n✅ Produced: %d | Consumed: %d\n", produced, consumed)
	fmt.Printf("⏱ Time taken to consume all messages: %v\n", totalTime)
}
