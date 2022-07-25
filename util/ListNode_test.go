package util

import (
	"fmt"
	"testing"
)

func TestListNodeInstance(t *testing.T) {
	nums := []int{1, 2, 4}
	instance := ListNodeInstance(nums)
	fmt.Printf("%v", instance)
}
