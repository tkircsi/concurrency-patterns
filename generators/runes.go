package generators

import "context"

func NewRuneGenerator(ctx context.Context, text string) <-chan rune {
	runeStream := make(chan rune)
	go func(s string) {
		defer close(runeStream)
		for _, r := range s {
			runeStream <- r
		}
	}(text)
	return runeStream
}
