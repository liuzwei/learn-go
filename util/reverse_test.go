package util

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	head := ListNodeInstance([]int{1, 2, 3, 4, 5, 6})
	result := Reverse(head)
	fmt.Printf("result is : %v", result.toString())
}

func TestReverseN(t *testing.T) {
	head := ListNodeInstance([]int{1, 2, 3, 4, 5, 6})
	result := ReverseN(head, 3)
	// 预期为 【3 2 1 4 5 6】
	fmt.Printf("result is : %v", result.toString())
}

func TestReverseBetween(t *testing.T) {
	head := ListNodeInstance([]int{1, 2, 3, 4, 5, 6})
	result := ReverseBetween(head, 2, 4)
	// 预期为 【1 4 3 2 5 6】
	fmt.Printf("result is : %v", result.toString())
}
