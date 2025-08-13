package producer

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"producer-consumer/internal/model"

	"golang.org/x/time/rate"
)

func StartProducers(producerCount int, rps int, duration time.Duration, messageCh chan<- model.Message, produced *int64, wgDone *sync.WaitGroup) {
	limiter := rate.NewLimiter(rate.Limit(rps), rps)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	
	go func() {
		time.Sleep(duration)
		cancel()
	}()

	for i := 0; i < producerCount; i++ {
		wgDone.Add(1)

		go func(id int) {
			defer wgDone.Done()
			
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				if err := limiter.Wait(ctx); err != nil {
					return // Context deadline exceeded
				}

				cur := atomic.AddInt64(produced, 1)
				msg := model.Message{
					Message:   fmt.Sprintf("data-%d", cur),
					CreatedAt: time.Now(),
				}
				
				select {
				case messageCh <- msg:
					// Message sent successfully
				case <-ctx.Done():
					return
				}
			}
		}(i + 1)
	}
}