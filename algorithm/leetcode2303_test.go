package algorithm

import "testing"

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		bracts [][]int
		income int
		result float64
	}{
		{
			[][]int{{3, 50}, {7, 10}, {12, 25}},
			10,
			2.65000,
		},
	}

	for _, tt := range tests {
		r := CalculateTax(tt.bracts, tt.income)
		if r != tt.result {
			t.Errorf("TestCalculateTax Error ! got=%v, want=%v", r, tt.result)
		}

	}
}
