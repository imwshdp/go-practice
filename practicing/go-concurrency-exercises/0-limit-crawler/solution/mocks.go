package solution

type mockResult struct {
	body string
	urls []string
}

type MockFetcher map[string]*mockResult

func (f MockFetcher) Fetch(url string) (string, []string, error) {
	return "", nil, nil
}

var fetcher = MockFetcher{}
