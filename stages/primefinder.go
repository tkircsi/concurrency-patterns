package stages

import (
	"context"
	"github.com/tkircsi/concurrency-patterns/snippets"
)

// PrimeFinder filters the prime numbers from input
func PrimeFinder(ctx context.Context, input <-chan interface{}) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-input:
				if ok == false {
					return
				}
				if val, ok := v.(int); ok {
					if snippets.IsPrime(val) {
						select {
						case <-ctx.Done():
						case outStream <- v:
						}

					}
				}
			}
		}
	}()
	return outStream
}
