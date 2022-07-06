package algorithm

import "testing"

func TestFinalValueAfterOperations(t *testing.T) {
	tests := []struct {
		input  []string
		result int
	}{
		{input: []string{"--X", "X++", "X++"}, result: 1},
		{input: []string{"++X", "++X", "X++"}, result: 3},
		{input: []string{"X++", "++X", "--X", "X--"}, result: 0},
	}

	for _, tt := range tests {
		r := FinalValueAfterOperations(tt.input)
		if r != tt.result {
			t.Errorf("FinalValueAfterOperations error, input=%v, got=%v, want=%v", tt.input, r, tt.result)
		}
	}
}
