package randGen

import (
	"math"
	"slices"
	"testing"
	"time"
)

func TestNew_Count(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{"zero count", 0},
		{"single value", 1},
		{"small count", 5},
		{"medium count", 100},
		{"large count", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := New(tt.count)

			var values []int
			for val := range result {
				values = append(values, val)
			}

			if len(values) != tt.count {
				t.Errorf("expected %d values, got %d", tt.count, len(values))
			}
		})
	}
}

func TestNew_ValuesInRange(t *testing.T) {
	count := 100
	result := New(count)

	for val := range result {
		if val < 0 {
			t.Errorf("expected non-negative value, got %d", val)
		}
		if val > math.MaxInt {
			t.Errorf("value exceeds MaxInt, got %d", val)
		}
	}
}

func TestNew_ChannelClosed(t *testing.T) {
	count := 10
	result := New(count)

	// Consume all values
	for range result {
	}

	// Channel should be closed
	_, ok := <-result
	if ok {
		t.Errorf("expected channel to be closed after all values consumed")
	}
}

func TestNew_DifferentRunsProduceDifferentResults(t *testing.T) {
	count := 10

	result1 := New(count)
	values1 := make([]int, 0, count)
	for val := range result1 {
		values1 = append(values1, val)
	}

	// Small delay to ensure different time seed
	time.Sleep(10 * time.Millisecond)

	result2 := New(count)
	values2 := make([]int, 0, count)
	for val := range result2 {
		values2 = append(values2, val)
	}

	// Values should likely differ between runs (probability extremely high)
	if slices.Equal(values1, values2) {
		t.Errorf("different runs produced identical results, unlikely but possible")
	}
}

func TestNew_NoDeadlock(t *testing.T) {
	done := make(chan struct{})

	go func() {
		result := New(100)
		for range result {
		}
		close(done)
	}()

	select {
	case <-done:
		// Success
	case <-time.After(5 * time.Second):
		t.Fatal("possible deadlock: New() did not complete within timeout")
	}
}

func TestNew_ConcurrentCalls(t *testing.T) {
	const goroutines = 10
	const countPerGoroutine = 50

	done := make(chan struct{}, goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			result := New(countPerGoroutine)
			for range result {
			}
			done <- struct{}{}
		}()
	}

	for i := 0; i < goroutines; i++ {
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			t.Fatal("concurrent call did not complete within timeout")
		}
	}
}
