package algorithm

import (
	"testing"

	"learn-go/util"
)

func TestMergeKLists(t *testing.T) {
	list1 := util.ListNodeInstance([]int{-495, -480, -474, -474, -470, -467})

	l := []*util.ListNode{list1}
	MergeKLists(l)
}
