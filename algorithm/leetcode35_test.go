package algorithm

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 2, 3}
	index := SearchInsert(nums, 4)
	fmt.Println("4的位置在", index)
}
