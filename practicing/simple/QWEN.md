# Project Overview

**tasks-practice** is a Go project for practicing and implementing algorithmic solutions and concurrent patterns. The project contains utility functions organized in the `internal` package structure.

## Main Components

### 1. Intersection Package (`internal/intersection/`)

Provides two generic functions for finding the intersection of two unsorted slices:

- **`SolutionWithSet[T cmp.Ordered]`** - Uses a hash set approach for O(n+m) time complexity
- **`SolutionWithSort[T cmp.Ordered]`** - Uses sorting and two-pointer technique

Both functions:
- Work with any ordered type (`int`, `string`, `float64`, `rune`, etc.)
- Handle duplicates correctly (each element appears only once in result)
- Handle empty slices and edge cases

### 2. MergeChan Package (`internal/mergeChan/`)

Provides a generic function for merging multiple channels:

- **`Merge[T comparable](chans ...<-chan T) <-chan T`** - Merges N input channels into a single output channel
- Uses goroutines and WaitGroup for concurrent reading
- Properly closes the result channel after all input channels are exhausted
- Works with any comparable type (`int`, `string`, `float64`, `bool`, etc.)

### 3. RandGen Package (`internal/randGen/`)

Provides a random number generator via channels:

- **`New(count int) chan int`** - Generates `count` random integers
- Uses seeded `math/rand` with nanosecond precision
- Returns values in range `[0, math.MaxInt)`
- Automatically closes channel after generating all values

### 4. Conveyer Package (`internal/conveyer/`)

Provides a generic channel-based transformation pipeline:

- **`New[T comparable](in <-chan T, fn func(T) T) <-chan T`** - Applies transformation function to each input value
- Processes values concurrently in a separate goroutine
- Preserves order of values
- Automatically closes output channel when input is exhausted

### 5. WorkerPool Package (`internal/workerPool/`)

Provides a worker pool pattern for concurrent job processing:

- **`New(jobsCh <-chan *Job, count int) <-chan int`** - Spawns `count` workers to process jobs
- `Job` struct contains `Value` and `Func` fields
- Workers share the same job channel for load balancing
- Properly closes result channel after all workers complete

## Project Structure

```
tasks-practice/
├── go.mod              # Go module definition (Go 1.25.5)
├── main.go             # Entry point with example usage
├── README.md           # Project documentation
├── QWEN.md             # This documentation file
└── internal/
    ├── conveyer/
    │   ├── conveyer.go         # Channel transformation pipeline
    │   └── conveyer_test.go    # Test suite
    ├── intersection/
    │   ├── intersection.go     # Slice intersection implementations
    │   └── intersection_test.go# Test suite
    ├── mergeChan/
    │   ├── mergeChan.go        # Channel merge implementation
    │   └── mergeChan_test.go   # Test suite
    ├── randGen/
    │   ├── randGen.go          # Random number generator
    │   └── randGen_test.go     # Test suite
    └── workerPool/
        ├── workerPool.go       # Worker pool implementation
        └── workerPool_test.go  # Test suite
```

## Building and Running

### Prerequisites

- Go 1.25.5 or later

### Commands

```bash
# Run the main program
go run main.go

# Run all tests
go test ./...

# Run tests for a specific package
go test ./internal/intersection/...
go test ./internal/mergeChan/...
go test ./internal/randGen/...
go test ./internal/conveyer/...
go test ./internal/workerPool/...

# Run tests with verbose output
go test ./... -v

# Build the project
go build .
```

## Development Conventions

### Testing Practices

- Tests follow a table-driven test pattern
- Each function has tests for multiple data types (int, string, float64, custom structs)
- Test cases cover:
  - Basic functionality
  - Empty inputs
  - Edge cases (duplicates, negative numbers, unsorted input)
  - Concurrent safety and deadlock detection
  - Stress tests with large datasets

### Code Style

- Generic functions use `cmp.Ordered` constraint for type-safe ordered operations
- Generic functions use `comparable` constraint for channel operations
- Clear function and variable naming
- Minimal comments (code is self-documenting)
- Results preserve order where applicable
- Channels are properly closed to prevent goroutine leaks

### Git Workflow

- Local commits are made before pushing
- Check `git status` and `git diff` before committing

## Example Usage

```go
// Intersection
result := intersection.SolutionWithSet([]int{1, 2, 3}, []int{3, 4, 5})
// Returns: [3]

// Merge channels (type inferred)
ch1 := make(chan int, 2)
ch2 := make(chan int, 2)
// ... populate and close channels
result := mergeChan.Merge(ch1, ch2)
// Read merged values from result channel

// Generate random numbers
randCh := randGen.New(10)
for val := range randCh {
    fmt.Println(val)
}

// Transform channel values
in := make(chan int)
go func() {
    for i := 1; i <= 5; i++ {
        in <- i
    }
    close(in)
}()
out := conveyer.New(in, func(x int) int { return x * 2 })

// Worker pool
jobsCh := make(chan *workerPool.Job, 10)
for i := 0; i < 10; i++ {
    jobsCh <- &workerPool.Job{
        Value: i,
        Func:  func(x int) int { return x * 2 },
    }
}
close(jobsCh)
results := workerPool.New(jobsCh, 5) // 5 workers
```
