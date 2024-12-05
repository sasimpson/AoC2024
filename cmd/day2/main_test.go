package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func Test_isSafe(t *testing.T) {
	tests := []struct {
		report string
		want   bool
		fixed  bool
	}{
		{"7 6 4 2 1", true, true},                //0
		{"1 2 7 8 9", false, false},              //1
		{"9 7 6 2 1", false, false},              //2
		{"1 3 2 4 5", false, true},               //3
		{"8 6 4 4 1", false, true},               //4
		{"1 3 6 7 9", true, true},                //5
		{"2 1 2 3 4", false, true},               //6
		{"1 2 3 5 5", false, true},               //7
		{"10 9 8 7 7", false, true},              //8
		{"10 9 7 7 7", false, false},             //9
		{"1 6 7 8 9", false, true},               //10
		{"29 28 27 25 26 25 22 20", false, true}, //11
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			line := strings.Split(tt.report, " ")
			report := make([]int, len(line))
			for i, v := range line {
				report[i], _ = strconv.Atoi(v)
			}
			got := isSafe(report)
			assert.Equal(t, tt.want, got)
			if got == false {
				gotFixed := bruteForceDamper(report)
				assert.Equal(t, tt.fixed, gotFixed)
			}
		})
	}
}
