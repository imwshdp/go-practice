package intersection

import (
	"slices"
	"testing"
)

func TestSolutionWithSet_Int(t *testing.T) {
	tests := []struct {
		name   string
		first  []int
		second []int
		want   []int
	}{
		{
			name:   "basic intersection",
			first:  []int{1, 2, 3, 4},
			second: []int{3, 4, 5, 6},
			want:   []int{3, 4},
		},
		{
			name:   "no intersection",
			first:  []int{1, 2, 3},
			second: []int{4, 5, 6},
			want:   []int{},
		},
		{
			name:   "empty first slice",
			first:  []int{},
			second: []int{1, 2, 3},
			want:   []int{},
		},
		{
			name:   "empty second slice",
			first:  []int{1, 2, 3},
			second: []int{},
			want:   []int{},
		},
		{
			name:   "both slices empty",
			first:  []int{},
			second: []int{},
			want:   []int{},
		},
		{
			name:   "duplicates in first slice",
			first:  []int{1, 2, 2, 3},
			second: []int{2, 3, 4},
			want:   []int{2, 3},
		},
		{
			name:   "duplicates in second slice",
			first:  []int{1, 2, 3},
			second: []int{2, 2, 3, 4},
			want:   []int{2, 3},
		},
		{
			name:   "identical slices",
			first:  []int{1, 2, 3},
			second: []int{1, 2, 3},
			want:   []int{1, 2, 3},
		},
		{
			name:   "negative numbers",
			first:  []int{-5, -2, 0, 3},
			second: []int{-2, 0, 4, 7},
			want:   []int{-2, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolutionWithSet(tt.first, tt.second)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SolutionWithSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionWithSet_String(t *testing.T) {
	tests := []struct {
		name   string
		first  []string
		second []string
		want   []string
	}{
		{
			name:   "basic intersection",
			first:  []string{"apple", "banana", "cherry"},
			second: []string{"banana", "cherry", "date"},
			want:   []string{"banana", "cherry"},
		},
		{
			name:   "no intersection",
			first:  []string{"apple", "banana"},
			second: []string{"cherry", "date"},
			want:   []string{},
		},
		{
			name:   "empty slices",
			first:  []string{},
			second: []string{"apple"},
			want:   []string{},
		},
		{
			name:   "duplicates",
			first:  []string{"a", "b", "b", "c"},
			second: []string{"b", "c"},
			want:   []string{"b", "c"},
		},
		{
			name:   "case sensitive",
			first:  []string{"Apple", "apple"},
			second: []string{"apple", "APPLE"},
			want:   []string{"apple"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolutionWithSet(tt.first, tt.second)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SolutionWithSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionWithSet_Float64(t *testing.T) {
	tests := []struct {
		name   string
		first  []float64
		second []float64
		want   []float64
	}{
		{
			name:   "basic intersection",
			first:  []float64{1.5, 2.5, 3.5},
			second: []float64{2.5, 3.5, 4.5},
			want:   []float64{2.5, 3.5},
		},
		{
			name:   "no intersection",
			first:  []float64{1.1, 2.2},
			second: []float64{3.3, 4.4},
			want:   []float64{},
		},
		{
			name:   "with zero and negative",
			first:  []float64{-1.5, 0.0, 1.5},
			second: []float64{0.0, 1.5, 2.5},
			want:   []float64{0.0, 1.5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolutionWithSet(tt.first, tt.second)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SolutionWithSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionWithSet_Rune(t *testing.T) {
	tests := []struct {
		name   string
		first  []rune
		second []rune
		want   []rune
	}{
		{
			name:   "basic intersection",
			first:  []rune{'a', 'b', 'c', 'd'},
			second: []rune{'c', 'd', 'e', 'f'},
			want:   []rune{'c', 'd'},
		},
		{
			name:   "no intersection",
			first:  []rune{'a', 'b'},
			second: []rune{'c', 'd'},
			want:   []rune{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolutionWithSet(tt.first, tt.second)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SolutionWithSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionWithSort_Int(t *testing.T) {
	tests := []struct {
		name   string
		first  []int
		second []int
		want   []int
	}{
		{
			name:   "basic intersection",
			first:  []int{1, 2, 3, 4},
			second: []int{3, 4, 5, 6},
			want:   []int{3, 4},
		},
		{
			name:   "no intersection",
			first:  []int{1, 2, 3},
			second: []int{4, 5, 6},
			want:   []int{},
		},
		{
			name:   "empty first slice",
			first:  []int{},
			second: []int{1, 2, 3},
			want:   []int{},
		},
		{
			name:   "empty second slice",
			first:  []int{1, 2, 3},
			second: []int{},
			want:   []int{},
		},
		{
			name:   "both slices empty",
			first:  []int{},
			second: []int{},
			want:   []int{},
		},
		{
			name:   "unsorted input first",
			first:  []int{4, 2, 1, 3},
			second: []int{3, 4, 5, 6},
			want:   []int{3, 4},
		},
		{
			name:   "unsorted input second",
			first:  []int{1, 2, 3},
			second: []int{6, 4, 5, 3},
			want:   []int{3},
		},
		{
			name:   "unsorted both",
			first:  []int{4, 2, 1, 3},
			second: []int{6, 4, 5, 3},
			want:   []int{3, 4},
		},
		{
			name:   "duplicates in first slice",
			first:  []int{1, 2, 2, 3},
			second: []int{2, 3, 4},
			want:   []int{2, 3},
		},
		{
			name:   "duplicates in second slice",
			first:  []int{1, 2, 3},
			second: []int{2, 2, 3, 4},
			want:   []int{2, 3},
		},
		{
			name:   "identical slices",
			first:  []int{1, 2, 3},
			second: []int{1, 2, 3},
			want:   []int{1, 2, 3},
		},
		{
			name:   "negative numbers",
			first:  []int{-5, -2, 0, 3},
			second: []int{-2, 0, 4, 7},
			want:   []int{-2, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolutionWithSort(tt.first, tt.second)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SolutionWithSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionWithSort_String(t *testing.T) {
	tests := []struct {
		name   string
		first  []string
		second []string
		want   []string
	}{
		{
			name:   "basic intersection",
			first:  []string{"apple", "banana", "cherry"},
			second: []string{"banana", "cherry", "date"},
			want:   []string{"banana", "cherry"},
		},
		{
			name:   "no intersection",
			first:  []string{"apple", "banana"},
			second: []string{"cherry", "date"},
			want:   []string{},
		},
		{
			name:   "empty slices",
			first:  []string{},
			second: []string{"apple"},
			want:   []string{},
		},
		{
			name:   "unsorted input",
			first:  []string{"cherry", "apple", "banana"},
			second: []string{"date", "banana", "cherry"},
			want:   []string{"banana", "cherry"},
		},
		{
			name:   "duplicates",
			first:  []string{"a", "b", "b", "c"},
			second: []string{"b", "c"},
			want:   []string{"b", "c"},
		},
		{
			name:   "case sensitive",
			first:  []string{"Apple", "apple"},
			second: []string{"apple", "APPLE"},
			want:   []string{"apple"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolutionWithSort(tt.first, tt.second)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SolutionWithSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionWithSort_Float64(t *testing.T) {
	tests := []struct {
		name   string
		first  []float64
		second []float64
		want   []float64
	}{
		{
			name:   "basic intersection",
			first:  []float64{1.5, 2.5, 3.5},
			second: []float64{2.5, 3.5, 4.5},
			want:   []float64{2.5, 3.5},
		},
		{
			name:   "no intersection",
			first:  []float64{1.1, 2.2},
			second: []float64{3.3, 4.4},
			want:   []float64{},
		},
		{
			name:   "with zero and negative",
			first:  []float64{-1.5, 0.0, 1.5},
			second: []float64{0.0, 1.5, 2.5},
			want:   []float64{0.0, 1.5},
		},
		{
			name:   "unsorted input",
			first:  []float64{3.5, 1.5, 2.5},
			second: []float64{4.5, 2.5, 3.5},
			want:   []float64{2.5, 3.5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolutionWithSort(tt.first, tt.second)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SolutionWithSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionWithSort_Rune(t *testing.T) {
	tests := []struct {
		name   string
		first  []rune
		second []rune
		want   []rune
	}{
		{
			name:   "basic intersection",
			first:  []rune{'a', 'b', 'c', 'd'},
			second: []rune{'c', 'd', 'e', 'f'},
			want:   []rune{'c', 'd'},
		},
		{
			name:   "no intersection",
			first:  []rune{'a', 'b'},
			second: []rune{'c', 'd'},
			want:   []rune{},
		},
		{
			name:   "unsorted input",
			first:  []rune{'d', 'b', 'a', 'c'},
			second: []rune{'f', 'c', 'd', 'e'},
			want:   []rune{'c', 'd'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolutionWithSort(tt.first, tt.second)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SolutionWithSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
