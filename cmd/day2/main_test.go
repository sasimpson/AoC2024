package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSafe(t *testing.T) {
	type args struct {
		report []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid increase",
			args: args{report: []int{1, 2, 3, 5, 6, 8}},
			want: true,
		}, {
			name: "valid decrease",
			args: args{report: []int{9, 7, 6, 5, 4, 2}},
			want: true,
		}, {
			name: "invalid decrease",
			args: args{report: []int{9, 5, 4, 2}},
			want: false,
		}, {
			name: "invalid increase",
			args: args{report: []int{1, 2, 3, 6, 8}},
			want: false,
		}, {
			name: "invalid no increase",
			args: args{report: []int{1, 2, 3, 3, 6, 8}},
			want: false,
		}, {
			name: "invalid no decrease",
			args: args{report: []int{9, 5, 5, 4, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSafe(tt.args.report)
			assert.Equal(t, tt.want, got)
		})
	}
}
