package algorithm

import (
	"fmt"
	"testing"
)

func TestValidString(t *testing.T) {
	result := ValidString("{[()]}")
	fmt.Println(result)
}
