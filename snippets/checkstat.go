package snippets

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Url      string
	Error    error
	Response *http.Response
	Duration time.Duration
}

func CheckHttpStatus(ctx context.Context, urls ...string) <-chan Result {
	results := make(chan Result)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			start := time.Now()
			resp, err := http.Get(u)
			result := Result{
				Url:      u,
				Error:    err,
				Response: resp,
				Duration: time.Since(start) / time.Millisecond,
			}

			select {
			case <-ctx.Done():
				return
			case results <- result:
			}
		}(url)
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(results)
	}(&wg)
	return results
}
