package limiter

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func PromiseAll(
	ctx context.Context,
	timeout time.Duration,
	urls []string,
) {
	reqCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	resCh := make(chan string, len(urls))

	wg := sync.WaitGroup{}
	wg.Add(len(urls))

	for _, u := range urls {
		go func() {
			defer wg.Done()

			req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, u, nil)
			if err != nil {
				return
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				return
			}

			resCh <- string(body)
		}()
	}

	go func() {
		defer close(resCh)
		wg.Wait()
	}()

	for res := range resCh {
		fmt.Println(res)
	}
}

func Demo() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
		"https://jsonplaceholder.typicode.com/posts/5",
	}

	PromiseAll(ctx, 1*time.Second, urls)
}
