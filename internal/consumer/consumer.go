package consumer

import (
	"database/sql"
	"fmt"
	"sync"
	"sync/atomic"

	"producer-consumer/internal/model"
)

func StartConsumers(db *sql.DB, consumerCount int, messageCh <-chan model.Message, wgDone *sync.WaitGroup, consumed *int64) {
	for i := 0; i < consumerCount; i++ {
		wgDone.Add(1)
		go func(id int) {
			defer wgDone.Done()
			for msg := range messageCh {
				_, err := db.Exec(
					"INSERT INTO messages ( message, created_at) VALUES ( ?, ?)",
					msg.Message, msg.CreatedAt,
				)
				if err != nil {
					fmt.Printf("DB insert error (consumer %d): %v\n", id, err)
					continue
				}
				atomic.AddInt64(consumed, 1)
			}
		}(i + 1)
	}
}
