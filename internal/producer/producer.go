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
	limiter := rate.NewLimiter(rate.Limit(rps), rps) // burst = rps to allow initial burst
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	for i := 0; i < producerCount; i++ {
		wgDone.Add(1)

		go func(id int) {
			defer wgDone.Done()
			for {
				if err := limiter.Wait(ctx); err != nil {
					return // duration ended
				}

				select {
				case <-ctx.Done():
					return
				default:
					cur := atomic.AddInt64(produced, 1)
					messageCh <- model.Message{
						Message:   fmt.Sprintf("data-%d", cur),
						CreatedAt: time.Now(),
					}
				}
			}
		}(i + 1)
	}
}
