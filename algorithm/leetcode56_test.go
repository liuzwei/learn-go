package algorithm

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	intervals := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	reuslt := Merge(intervals)
	fmt.Println(reuslt)
}
