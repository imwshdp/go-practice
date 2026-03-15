package mergeChan

import (
	"slices"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestMerge_NoChannels(t *testing.T) {
	result := Merge()

	_, ok := <-result
	if ok {
		t.Errorf("expected closed channel, got value")
	}
}

func TestMerge_SingleChannel(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	result := Merge(ch)

	var values []int
	for val := range result {
		values = append(values, val)
	}

	expected := []int{1, 2, 3}

	sort.Ints(values)
	sort.Ints(expected)

	if len(values) != len(expected) {
		t.Errorf("expected %d values, got %d", len(expected), len(values))
	}

	if !slices.Equal(values, expected) {
		t.Errorf("expected %v, got %v", expected, values)
	}
}

func TestMerge_MultipleChannels(t *testing.T) {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2
	close(ch1)

	ch2 <- 3
	ch2 <- 4
	close(ch2)

	ch3 <- 5
	ch3 <- 6
	close(ch3)

	result := Merge(ch1, ch2, ch3)

	var values []int
	for val := range result {
		values = append(values, val)
	}

	expected := []int{1, 2, 3, 4, 5, 6}

	sort.Ints(values)
	sort.Ints(expected)

	if len(values) != len(expected) {
		t.Errorf("expected %d values, got %d", len(expected), len(values))
	}

	if !slices.Equal(values, expected) {
		t.Errorf("expected %v, got %v", expected, values)
	}
}

func TestMerge_EmptyChannels(t *testing.T) {
	ch1 := make(chan int)
	close(ch1)

	ch2 := make(chan int)
	close(ch2)

	result := Merge(ch1, ch2)

	count := 0
	for range result {
		count++
	}

	if count != 0 {
		t.Errorf("expected 0 values from empty channels, got %d", count)
	}
}

func TestMerge_DifferentLengths(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 3)
	ch3 := make(chan int, 2)

	ch1 <- 1
	close(ch1)

	ch2 <- 2
	ch2 <- 3
	ch2 <- 4
	close(ch2)

	ch3 <- 5
	ch3 <- 6
	close(ch3)

	result := Merge(ch1, ch2, ch3)

	var values []int
	for val := range result {
		values = append(values, val)
	}

	expected := []int{1, 2, 3, 4, 5, 6}

	sort.Ints(values)
	sort.Ints(expected)

	if len(values) != len(expected) {
		t.Errorf("expected %d values, got %d", len(expected), len(values))
	}

	if !slices.Equal(values, expected) {
		t.Errorf("expected %v, got %v", expected, values)
	}
}

func TestMerge_Duplicates(t *testing.T) {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2
	close(ch1)

	ch2 <- 2
	ch2 <- 3
	close(ch2)

	result := Merge(ch1, ch2)

	var values []int
	for val := range result {
		values = append(values, val)
	}

	expected := []int{1, 2, 2, 3}

	sort.Ints(values)
	sort.Ints(expected)

	if len(values) != len(expected) {
		t.Errorf("expected %d values (including duplicates), got %d", len(expected), len(values))
	}

	if !slices.Equal(values, expected) {
		t.Errorf("expected %v, got %v", expected, values)
	}
}

func TestMerge_NegativeNumbers(t *testing.T) {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	ch1 <- -5
	ch1 <- -2
	close(ch1)

	ch2 <- 0
	ch2 <- 3
	close(ch2)

	result := Merge(ch1, ch2)

	var values []int
	for val := range result {
		values = append(values, val)
	}

	expected := []int{-5, -2, 0, 3}

	sort.Ints(values)
	sort.Ints(expected)

	if len(values) != len(expected) {
		t.Errorf("expected %d values, got %d", len(expected), len(values))
	}

	if !slices.Equal(values, expected) {
		t.Errorf("expected %v, got %v", expected, values)
	}
}

func TestMerge_LargeValues(t *testing.T) {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	ch1 <- 1000000
	ch1 <- 999999
	close(ch1)

	ch2 <- -1000000
	ch2 <- -999999
	close(ch2)

	result := Merge(ch1, ch2)

	var values []int
	for val := range result {
		values = append(values, val)
	}

	expected := []int{-1000000, -999999, 999999, 1000000}

	sort.Ints(values)
	sort.Ints(expected)

	if len(values) != len(expected) {
		t.Errorf("expected %d values, got %d", len(expected), len(values))
	}

	if !slices.Equal(values, expected) {
		t.Errorf("expected %v, got %v", expected, values)
	}
}

func TestMerge_ConcurrentSafety(t *testing.T) {
	const numChannels = 10
	const valuesPerChannel = 100

	channels := make([]<-chan int, numChannels)

	for i := 0; i < numChannels; i++ {
		ch := make(chan int, valuesPerChannel)
		for j := 0; j < valuesPerChannel; j++ {
			ch <- i*valuesPerChannel + j
		}
		close(ch)
		channels[i] = ch
	}

	result := Merge(channels...)

	var values []int
	var mu sync.Mutex
	done := make(chan struct{})

	go func() {
		for val := range result {
			mu.Lock()
			values = append(values, val)
			mu.Unlock()
		}
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for merge to complete")
	}

	expectedCount := numChannels * valuesPerChannel
	if len(values) != expectedCount {
		t.Errorf("expected %d values, got %d", expectedCount, len(values))
	}

	expected := make([]int, 0, expectedCount)
	for i := 0; i < numChannels; i++ {
		for j := 0; j < valuesPerChannel; j++ {
			expected = append(expected, i*valuesPerChannel+j)
		}
	}

	sort.Ints(values)
	sort.Ints(expected)

	if !slices.Equal(values, expected) {
		t.Errorf("expected %v, got %v", expected, values)
	}
}

func TestMerge_ChannelClosing(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	close(ch1)

	ch2 <- 2
	close(ch2)

	result := Merge(ch1, ch2)

	var values []int
	for val := range result {
		values = append(values, val)
	}

	expected := []int{1, 2}

	sort.Ints(values)
	sort.Ints(expected)

	if len(values) != len(expected) {
		t.Errorf("expected %d values, got %d", len(expected), len(values))
	}

	if !slices.Equal(values, expected) {
		t.Errorf("expected %v, got %v", expected, values)
	}

	_, ok := <-result
	if ok {
		t.Errorf("expected result channel to be closed")
	}
}
