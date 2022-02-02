package patterns

import "context"

func FanOut(
	ctx context.Context,
	mul int,
	inputStream <-chan interface{},
	stageFn func(context.Context, <-chan interface{}) <-chan interface{},
) []<-chan interface{} {
	multiplex := make([]<-chan interface{}, mul)
	for i := 0; i < mul; i++ {
		multiplex[i] = stageFn(ctx, inputStream)
	}
	return multiplex
}
