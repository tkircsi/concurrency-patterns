package stages

import "context"

// Map is a function that applies the given function on the input
func Map(ctx context.Context, input <-chan interface{}, fn func(v interface{}) interface{}) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-input:
				vv := fn(v)
				outStream <- vv
			}
		}
	}()
	return outStream
}
