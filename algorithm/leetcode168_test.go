package algorithm

import (
	"testing"
)

func TestConvertToTitle(t *testing.T) {

	tests := []struct {
		colum int
		want  string
	}{
		{
			701, "ZY",
		},
		{
			2147483647, "FXSHRXW",
		},
	}

	for _, tt := range tests {
		result := ConvertToTitle(tt.colum)
		if result != tt.want {
			t.Errorf("ConvertToTitle got=%v, want=%v", result, tt.want)
		}
	}
}
