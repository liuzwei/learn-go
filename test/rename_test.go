package test

import "testing"

func Test_renameFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "rename",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renameFile()
		})
	}
}
