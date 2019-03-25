package main

import (
	"testing"
)

func testisEven(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "success - even number - 0",
			args:    args{num: 0},
			want:    true,
			wantErr: false,
		},
		{
			name:    "success - odd number - 1",
			args:    args{num: 1},
			want:    false,
			wantErr: false,
		},
		{
			name:    "success - even number - 2",
			args:    args{num: 2},
			want:    true,
			wantErr: false,
		},
		{
			name:    "success - odd number - 3",
			args:    args{num: 3},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Logf("Running %v", tt.name)

		got, err := isEven(tt.args.num)
		if (err != nil) && !tt.wantErr {
			t.Errorf("isEven() returned err: %v, wantErr = %v", err, tt.wantErr)
			return
		}
		if got != tt.want {
			t.Errorf("isEven() = %v, wanted %v", got, tt.want)
		}
	}
}
