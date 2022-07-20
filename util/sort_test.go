package util

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		ar     []int
		result []int
	}{
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		result := BubbleSort(tt.ar)
		fmt.Printf("%v", result)
	}
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		ar     []int
		result []int
	}{
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		result := InsertionSort(tt.ar)
		fmt.Printf("%v", result)
	}
}
