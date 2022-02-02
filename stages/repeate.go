package stages

import "context"

// Repeate is a generic stage. Not depends on the streamed value
func Repeate(ctx context.Context, values ...interface{}) <-chan interface{} {
	outStream := make(chan interface{}, len(values))
	go func() {
		defer close(outStream)
		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				case outStream <- v:
				}
			}
		}
	}()
	return outStream
}

// RepeateFn continuously call fn and produce the result via channel
func RepeateFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for {
			v := fn()
			select {
			case <-ctx.Done():
				return
			case outStream <- v:
			}
		}
	}()
	return outStream
}
