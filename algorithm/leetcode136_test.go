package algorithm

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{100, 200, 300, 200, 25, 100, 300}
	result := SingleNumber(nums)
	fmt.Println(result)
}
