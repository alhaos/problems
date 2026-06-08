package pipeline

import (
	"context"
	"strings"
)

func Pipeline(ctx context.Context, prefix string, words ...string) <-chan string {
	gen := generate(ctx, words...)
	up := toUpper(gen)
	return filterPrefix(up, prefix)
}

func generate(ctx context.Context, words ...string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, n := range words {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func toUpper(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for word := range in {
			out <- strings.ToUpper(word)
		}
	}()
	return out
}

func filterPrefix(in <-chan string, prefix string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for word := range in {
			if strings.HasPrefix(word, prefix) {
				out <- word
			}
		}
	}()
	return out
}
