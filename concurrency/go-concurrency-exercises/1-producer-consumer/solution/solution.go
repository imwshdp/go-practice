package solution

import (
	"errors"
	"fmt"
	"time"
)

// === Mocks ===
var ErrEOF = errors.New("End of File")

type Tweet struct {
	Username string
	Text     string
}

func (t *Tweet) IsTalkingAboutGo() bool {
	return false
}

type Stream struct {
	pos    int
	tweets []Tweet
}

func (s *Stream) Next() (*Tweet, error) {
	return nil, nil
}

func GetMockStream() Stream {
	return Stream{}
}

// =============

func producer(stream Stream, tweetsChan chan<- *Tweet) {
	defer close(tweetsChan)

	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return
		}

		tweetsChan <- tweet
	}
}

func consumer(tweetsChan <-chan *Tweet) {
	for t := range tweetsChan {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tweetsChan := make(chan *Tweet)

	go producer(stream, tweetsChan)

	consumer(tweetsChan)

	fmt.Printf("Process took %s\n", time.Since(start))
}
