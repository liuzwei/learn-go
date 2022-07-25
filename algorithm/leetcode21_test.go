package algorithm

import (
	"testing"

	"learn-go/util"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := util.ListNodeInstance([]int{1, 3, 4})
	list2 := util.ListNodeInstance([]int{1, 2, 4})
	MergeTwoLists(list1, list2)

}
