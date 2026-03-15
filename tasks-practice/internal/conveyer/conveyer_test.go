package conveyer

import (
	"slices"
	"testing"
	"time"
)

func TestNew_IntTransform(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) int
		expected []int
	}{
		{
			name:     "double values",
			input:    []int{1, 2, 3, 4, 5},
			fn:       func(x int) int { return x * 2 },
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "square values",
			input:    []int{1, 2, 3, 4},
			fn:       func(x int) int { return x * x },
			expected: []int{1, 4, 9, 16},
		},
		{
			name:     "negate values",
			input:    []int{1, -2, 3, -4},
			fn:       func(x int) int { return -x },
			expected: []int{-1, 2, -3, 4},
		},
		{
			name:     "identity function",
			input:    []int{1, 2, 3},
			fn:       func(x int) int { return x },
			expected: []int{1, 2, 3},
		},
		{
			name:     "constant function",
			input:    []int{1, 2, 3},
			fn:       func(x int) int { return 42 },
			expected: []int{42, 42, 42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan int)
			go func() {
				for _, v := range tt.input {
					in <- v
				}
				close(in)
			}()

			out := New(in, tt.fn)

			var result []int
			for val := range out {
				result = append(result, val)
			}

			if !slices.Equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestNew_StringTransform(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func(string) string
		expected []string
	}{
		{
			name:     "uppercase",
			input:    []string{"hello", "world"},
			fn:       func(s string) string { return s + "!" },
			expected: []string{"hello!", "world!"},
		},
		{
			name:     "prepend prefix",
			input:    []string{"a", "abc", "abcdef"},
			fn:       func(s string) string { return "prefix:" + s },
			expected: []string{"prefix:a", "prefix:abc", "prefix:abcdef"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan string)
			go func() {
				for _, v := range tt.input {
					in <- v
				}
				close(in)
			}()

			out := New(in, tt.fn)

			var result []string
			for val := range out {
				result = append(result, val)
			}

			if !slices.Equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestNew_EmptyChannel(t *testing.T) {
	in := make(chan int)
	close(in)

	out := New(in, func(x int) int { return x * 2 })

	count := 0
	for range out {
		count++
	}

	if count != 0 {
		t.Errorf("expected 0 values from empty channel, got %d", count)
	}
}

func TestNew_ChannelClosed(t *testing.T) {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)

	out := New(in, func(x int) int { return x })

	// Consume all values
	for range out {
	}

	// Channel should be closed
	_, ok := <-out
	if ok {
		t.Errorf("expected output channel to be closed")
	}
}

func TestNew_NoDeadlock(t *testing.T) {
	in := make(chan int, 100)
	for i := 0; i < 100; i++ {
		in <- i
	}
	close(in)

	done := make(chan struct{})

	go func() {
		out := New(in, func(x int) int { return x * 2 })
		for range out {
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
		go func(id int) {
			in := make(chan int, countPerGoroutine)
			for j := 0; j < countPerGoroutine; j++ {
				in <- id*countPerGoroutine + j
			}
			close(in)

			out := New(in, func(x int) int { return x + 1 })
			for range out {
			}
			done <- struct{}{}
		}(i)
	}

	for i := 0; i < goroutines; i++ {
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			t.Fatal("concurrent call did not complete within timeout")
		}
	}
}

func TestNew_LargeDataset(t *testing.T) {
	const count = 10000

	in := make(chan int, count)
	expected := make([]int, 0, count)
	for i := 0; i < count; i++ {
		in <- i
		expected = append(expected, i*2)
	}
	close(in)

	out := New(in, func(x int) int { return x * 2 })

	var result []int
	for val := range out {
		result = append(result, val)
	}

	if len(result) != count {
		t.Errorf("expected %d values, got %d", count, len(result))
	}

	if !slices.Equal(result, expected) {
		t.Errorf("result mismatch")
	}
}

func TestNew_CustomStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	in := make(chan Person, 3)
	in <- Person{"Alice", 25}
	in <- Person{"Bob", 30}
	in <- Person{"Charlie", 35}
	close(in)

	out := New(in, func(p Person) Person {
		p.Age += 1
		return p
	})

	var result []Person
	for val := range out {
		result = append(result, val)
	}

	expected := []Person{
		{"Alice", 26},
		{"Bob", 31},
		{"Charlie", 36},
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d values, got %d", len(expected), len(result))
	}

	for i, p := range result {
		if p.Name != expected[i].Name || p.Age != expected[i].Age {
			t.Errorf("index %d: expected %v, got %v", i, expected[i], p)
		}
	}
}
