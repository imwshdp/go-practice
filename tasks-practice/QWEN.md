# Project Overview

**tasks-practice** is a Go project for practicing and implementing algorithmic solutions. The project contains utility functions organized in the `internal` package structure.

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

## Project Structure

```
tasks-practice/
├── go.mod              # Go module definition (Go 1.25.5)
├── main.go             # Entry point with example usage
├── QWEN.md             # This documentation file
└── internal/
    ├── intersection/
    │   ├── intersection.go      # Intersection implementations
    │   └── intersection_test.go # Comprehensive test suite
    └── mergeChan/
        ├── mergeChan.go         # Channel merge implementation
        └── mergeChan_test.go    # Test suite for merge function
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

# Run tests with verbose output
go test ./... -v

# Build the project
go build .
```

## Development Conventions

### Testing Practices

- Tests follow a table-driven test pattern
- Each function has tests for multiple data types (int, string, float64, rune)
- Test cases cover:
  - Basic functionality
  - Empty inputs
  - Edge cases (duplicates, negative numbers, unsorted input)
  - Concurrent safety (for mergeChan)

### Code Style

- Generic functions use `cmp.Ordered` constraint for type-safe ordered operations
- Generic functions use `comparable` constraint for channel operations
- Clear function and variable naming
- Minimal comments (code is self-documenting)
- Results preserve order where applicable

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

// Merge channels with explicit type
strCh1 := make(chan string, 2)
strCh2 := make(chan string, 2)
result := mergeChan.Merge[string](strCh1, strCh2)
```
