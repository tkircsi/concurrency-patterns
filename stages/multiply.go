package stages

import "context"

func Multiply(ctx context.Context, inStream <-chan int, multiplier int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for n := range inStream {
			select {
			case <-ctx.Done():
				return
			case outStream <- n * multiplier:
			}
		}
	}()
	return outStream
}
