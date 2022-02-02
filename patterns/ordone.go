package patterns

import (
	"context"
)

// OrDone doesn't send any data if the source channel is closed
func OrDone(ctx context.Context, c <-chan interface{}) <-chan interface{} {
	val := make(chan interface{})
	go func() {
		defer close(val)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case <-ctx.Done():
				case val <- v:
				}
			}
		}
	}()
	return val
}
