package generators

import (
	"context"
	"fmt"
	"time"
)

func NewFibonacciGenerator(ctx context.Context, limit int) <-chan int {
	fibStream := make(chan int)
	go func(limit int) {
		defer close(fibStream)
		a := 0
		b := 1
		fibStream <- a
		for b < limit {
			a, b = b, a+b
			time.Sleep(time.Millisecond * 200)
			select {
			case <-ctx.Done():
				fmt.Println("cancelled")
				return
			case fibStream <- a:
			}
		}
	}(limit)
	return fibStream
}
