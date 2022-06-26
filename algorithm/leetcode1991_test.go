package algorithm

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	nums := []int{1}
	index := PivotIndex(nums)
	fmt.Println(index)
}
