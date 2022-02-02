package patterns

import (
	"context"
	"time"
)

func Or(ctx context.Context, cancel context.CancelFunc, channels ...<-chan time.Time) <-chan time.Time {
	outStream := make(chan time.Time)
	for i, ch := range channels {
		go func(n int, ch <-chan time.Time) {
			defer cancel()
			for {
				select {
				case <-ctx.Done():
					//fmt.Printf("%d chan cancelled\n", n)
					return
				case t := <-ch:
					outStream <- t
					cancel()
				}
			}
		}(i, ch)
	}
	return outStream
}
