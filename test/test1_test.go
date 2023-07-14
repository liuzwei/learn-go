package test

import "testing"

func Test_myFunc(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "10 + 10",
			args:    args{10, 10},
			want:    20,
			wantErr: false,
		},
		{
			name:    "100 + 100",
			args:    args{100, 100},
			want:    200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := myFunc(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("myFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("myFunc() got = %v, want %v", got, tt.want)
			}
		})
	}
}
