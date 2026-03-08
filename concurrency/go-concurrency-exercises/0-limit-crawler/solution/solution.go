package solution

import (
	"fmt"
	"sync"
	"time"
)

// === Mocks ===
type mockResult struct {
	body string
	urls []string
}

type MockFetcher map[string]*mockResult

func (f MockFetcher) Fetch(url string) (string, []string, error) {
	return "", nil, nil
}

var fetcher = MockFetcher{}

// =============

var ticker = time.NewTicker(1 * time.Second)

func Crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	<-ticker.C
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	wg.Add(len(urls))
	for _, u := range urls {
		go Crawl(u, depth-1, wg)
	}
}

func main() {
	defer ticker.Stop()
	var wg sync.WaitGroup

	wg.Add(1)
	Crawl("http://golang.org/", 4, &wg)
	wg.Wait()
}
