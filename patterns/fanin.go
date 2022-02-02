package patterns

import (
	"context"
	"sync"
)

func FanIn(
	ctx context.Context,
	channels ...<-chan interface{},
) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case multiplexedStream <- i:
			}
		}
	}
	wg.Add(len(channels))

	// Close multiplexedStream if all input channel is closed
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	// Start goroutines to read all input channels
	for _, c := range channels {
		go multiplex(c)
	}
	return multiplexedStream
}
