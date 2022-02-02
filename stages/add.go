package stages

import "context"

func Add(ctx context.Context, inStream <-chan int, value int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for n := range inStream {
			select {
			case <-ctx.Done():
				return
			case outStream <- n + value:
			}
		}
	}()
	return outStream
}
