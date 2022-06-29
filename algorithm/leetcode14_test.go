package algorithm

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	result := LongestCommonPrefix(strs)
	fmt.Println(result)
}
