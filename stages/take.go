package stages

import "context"

// Take is a generic stage. Not depends on the streamed value
func Take(ctx context.Context, input <-chan interface{}, n int) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case v := <-input:
				outStream <- v
			}
		}
	}()
	return outStream
}
