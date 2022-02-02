package generators

import (
	"context"
)

func NewIntGenerator(ctx context.Context, start int) <-chan int {
	outStream := make(chan int)
	go func(i int) {
		defer close(outStream)
		for {
			select {
			case <-ctx.Done():
				return
			case outStream <- i:
				i++
			}
		}
	}(start)
	return outStream
}

func NewIntSliceGenerator(ctx context.Context, nums []int) <-chan int {
	intStream := make(chan int, len(nums))
	go func() {
		defer close(intStream)
		for _, i := range nums {
			select {
			case <-ctx.Done():
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}
