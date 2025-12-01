package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "traceId", "ABC1234")
	_, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done triggered:", ctx.Err())
				return
			default:
				fmt.Println("Running.. TraceID = ", ctx.Value("traceId"))
				time.Sleep(500 * time.Millisecond)
			}

		}
	}(context.Background())
	time.Sleep(4 * time.Second)
}
