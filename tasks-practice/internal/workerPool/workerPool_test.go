package workerPool

import (
	"slices"
	"sort"
	"testing"
	"time"
)

func TestNew_BasicProcessing(t *testing.T) {
	tests := []struct {
		name       string
		jobs       []*Job
		workerCount int
		expected   []int
	}{
		{
			name: "single worker single job",
			jobs: []*Job{
				{Value: 5, Func: func(x int) int { return x * 2 }},
			},
			workerCount: 1,
			expected:    []int{10},
		},
		{
			name: "single worker multiple jobs",
			jobs: []*Job{
				{Value: 1, Func: func(x int) int { return x * 2 }},
				{Value: 2, Func: func(x int) int { return x * 2 }},
				{Value: 3, Func: func(x int) int { return x * 2 }},
			},
			workerCount: 1,
			expected:    []int{2, 4, 6},
		},
		{
			name: "multiple workers multiple jobs",
			jobs: []*Job{
				{Value: 1, Func: func(x int) int { return x * 2 }},
				{Value: 2, Func: func(x int) int { return x * 2 }},
				{Value: 3, Func: func(x int) int { return x * 2 }},
				{Value: 4, Func: func(x int) int { return x * 2 }},
			},
			workerCount: 2,
			expected:    []int{2, 4, 6, 8},
		},
		{
			name: "square function",
			jobs: []*Job{
				{Value: 2, Func: func(x int) int { return x * x }},
				{Value: 3, Func: func(x int) int { return x * x }},
				{Value: 4, Func: func(x int) int { return x * x }},
			},
			workerCount: 2,
			expected:    []int{4, 9, 16},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jobsCh := make(chan *Job, len(tt.jobs))
			for _, job := range tt.jobs {
				jobsCh <- job
			}
			close(jobsCh)

			result := New(jobsCh, tt.workerCount)

			var values []int
			for val := range result {
				values = append(values, val)
			}

			sort.Ints(values)
			sort.Ints(tt.expected)

			if !slices.Equal(values, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, values)
			}
		})
	}
}

func TestNew_EmptyJobs(t *testing.T) {
	jobsCh := make(chan *Job)
	close(jobsCh)

	result := New(jobsCh, 3)

	count := 0
	for range result {
		count++
	}

	if count != 0 {
		t.Errorf("expected 0 results from empty jobs channel, got %d", count)
	}
}

func TestNew_ChannelClosed(t *testing.T) {
	jobsCh := make(chan *Job, 2)
	jobsCh <- &Job{Value: 1, Func: func(x int) int { return x }}
	jobsCh <- &Job{Value: 2, Func: func(x int) int { return x }}
	close(jobsCh)

	result := New(jobsCh, 2)

	// Consume all values
	for range result {
	}

	// Channel should be closed
	_, ok := <-result
	if ok {
		t.Errorf("expected result channel to be closed")
	}
}

func TestNew_NoDeadlock(t *testing.T) {
	jobsCh := make(chan *Job, 100)
	for i := 0; i < 100; i++ {
		jobsCh <- &Job{Value: i, Func: func(x int) int { return x * 2 }}
	}
	close(jobsCh)

	done := make(chan struct{})

	go func() {
		result := New(jobsCh, 10)
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

func TestNew_ConcurrentWorkers(t *testing.T) {
	const jobCount = 100
	const workerCount = 10

	jobsCh := make(chan *Job, jobCount)
	for i := 0; i < jobCount; i++ {
		jobsCh <- &Job{Value: i, Func: func(x int) int { return x + 1 }}
	}
	close(jobsCh)

	result := New(jobsCh, workerCount)

	var values []int
	for val := range result {
		values = append(values, val)
	}

	if len(values) != jobCount {
		t.Errorf("expected %d results, got %d", jobCount, len(values))
	}

	// Verify all values are correct (order may vary due to concurrency)
	sort.Ints(values)
	expected := make([]int, jobCount)
	for i := 0; i < jobCount; i++ {
		expected[i] = i + 1
	}

	if !slices.Equal(values, expected) {
		t.Errorf("result mismatch")
	}
}

func TestNew_MultipleWorkerCounts(t *testing.T) {
	tests := []struct {
		name        string
		workerCount int
	}{
		{"single worker", 1},
		{"two workers", 2},
		{"five workers", 5},
		{"ten workers", 10},
		{"many workers", 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const jobCount = 100

			jobsCh := make(chan *Job, jobCount)
			for i := 0; i < jobCount; i++ {
				jobsCh <- &Job{Value: i, Func: func(x int) int { return x }}
			}
			close(jobsCh)

			result := New(jobsCh, tt.workerCount)

			count := 0
			for range result {
				count++
			}

			if count != jobCount {
				t.Errorf("expected %d results with %d workers, got %d", jobCount, tt.workerCount, count)
			}
		})
	}
}

func TestNew_SlowJobs(t *testing.T) {
	jobsCh := make(chan *Job, 10)
	for i := 0; i < 10; i++ {
		jobsCh <- &Job{
			Value: i,
			Func: func(x int) int {
				time.Sleep(10 * time.Millisecond)
				return x * 2
			},
		}
	}
	close(jobsCh)

	start := time.Now()
	result := New(jobsCh, 5)

	var values []int
	for val := range result {
		values = append(values, val)
	}
	elapsed := time.Since(start)

	if len(values) != 10 {
		t.Errorf("expected 10 results, got %d", len(values))
	}

	// With 5 workers and 10 jobs taking 10ms each, should complete in ~20ms
	if elapsed > 100*time.Millisecond {
		t.Errorf("took too long: %v, expected ~20ms with 5 workers", elapsed)
	}
}

func TestNew_ResultOrder(t *testing.T) {
	jobsCh := make(chan *Job, 5)
	for i := 1; i <= 5; i++ {
		jobsCh <- &Job{Value: i, Func: func(x int) int { return x }}
	}
	close(jobsCh)

	result := New(jobsCh, 1)

	var values []int
	for val := range result {
		values = append(values, val)
	}

	// With single worker, order should be preserved
	expected := []int{1, 2, 3, 4, 5}
	if !slices.Equal(values, expected) {
		t.Errorf("expected %v, got %v", expected, values)
	}
}

func TestNew_StressTest(t *testing.T) {
	const jobCount = 1000
	const workerCount = 20

	jobsCh := make(chan *Job, jobCount)
	for i := 0; i < jobCount; i++ {
		jobsCh <- &Job{Value: i, Func: func(x int) int { return i }}
	}
	close(jobsCh)

	result := New(jobsCh, workerCount)

	count := 0
	for range result {
		count++
	}

	if count != jobCount {
		t.Errorf("expected %d results, got %d", jobCount, count)
	}
}

func TestNew_ConcurrentPoolCreation(t *testing.T) {
	const poolCount = 10
	const jobsPerPool = 50

	done := make(chan struct{}, poolCount)

	for p := 0; p < poolCount; p++ {
		go func(id int) {
			jobsCh := make(chan *Job, jobsPerPool)
			for i := 0; i < jobsPerPool; i++ {
				jobsCh <- &Job{Value: id*jobsPerPool + i, Func: func(x int) int { return x }}
			}
			close(jobsCh)

			result := New(jobsCh, 5)
			for range result {
			}
			done <- struct{}{}
		}(p)
	}

	for i := 0; i < poolCount; i++ {
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			t.Fatal("concurrent pool creation did not complete within timeout")
		}
	}
}
