package stages

import (
	"context"
)

func IntToInterface(ctx context.Context, inStream <-chan int) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-inStream:
				if ok == false {
					return
				}
				var iv interface{}
				iv = v
				select {
				case <-ctx.Done():
				case outStream <- iv:
				}
			}
		}
	}()
	return outStream
}
