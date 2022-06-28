package algorithm

import (
	"fmt"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	nums := [][]int{{1, 2}, {3, 4}}
	result := FindDiagonalOrder(nums)
	fmt.Println(result)
}
