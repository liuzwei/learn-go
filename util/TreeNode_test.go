package util

import (
	"fmt"
	"testing"
)

func TestConstruct(t *testing.T) {
	node := Construct([]any{1, 2, 3, nil, 4, nil, 5})
	fmt.Println(node)
}
