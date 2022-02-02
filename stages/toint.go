package stages

import "context"

// ToInt is a casting stage. Casts generic value to an int value
func ToInt(ctx context.Context, input <-chan interface{}) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for v := range input {
			select {
			case <-ctx.Done():
				return
			case outStream <- v.(int):
			}
		}
	}()
	return outStream
}
