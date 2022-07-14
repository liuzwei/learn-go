package algorithm

import (
	"reflect"
	"testing"

	"learn-go/util"
)

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		input []any
		out   []int
	}{
		{
			[]any{1, nil, 2, 3},
			[]int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		result := PostorderTraversal(util.Construct(tt.input))
		if !reflect.DeepEqual(result, tt.out) {
			t.Errorf("PostorderTraversal Error! got=%v, want=%v", result, tt.out)
		}
	}
}
