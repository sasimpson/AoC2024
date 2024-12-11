package main

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadRules(t *testing.T) {
	type args struct {
		rules [][]string
	}
	tests := []struct {
		name  string
		rules [][]string
		want  map[int][]rule
	}{
		{
			name: "example 1",
			rules: [][]string{
				{"47", "53"}, {"97", "13"}, {"97", "61"}, {"97", "47"}, {"75", "29"}, {"61", "13"}, {"75", "53"}, {"29", "13"}, {"97", "29"}, {"53", "29"},
			},
			want: map[int][]rule{
				47: {{47, 53}},
				97: {{97, 13}, {97, 61}, {97, 47}, {97, 29}},
				75: {{75, 29}, {75, 53}},
				61: {{61, 13}},
				29: {{29, 13}},
				53: {{53, 29}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadRules(tt.rules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadRules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadUpdates(t *testing.T) {
	tests := []struct {
		name string
		data [][]string
		want []update
	}{
		{
			name: "example 1",
			data: [][]string{{"75", "47", "61", "53", "29"}, {"97", "61", "53", "29", "13"}, {"75", "29", "13"}},
			want: []update{
				{pages: []int{75, 47, 61, 53, 29}},
				{pages: []int{97, 61, 53, 29, 13}},
				{pages: []int{75, 29, 13}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadUpdates(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadUpdates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wanta    [][]string
		wantb    [][]string
	}{
		{
			name:     "example1",
			filename: "testdata/data5_test.txt",
			wanta:    [][]string{{"47", "53"}, {"97", "13"}, {"97", "61"}, {"97", "47"}, {"75", "29"}, {"61", "13"}, {"75", "53"}, {"29", "13"}, {"97", "29"}, {"53", "29"}, {"61", "53"}, {"97", "53"}, {"61", "29"}, {"47", "13"}, {"75", "47"}, {"97", "75"}, {"47", "61"}, {"75", "61"}, {"47", "29"}, {"75", "13"}, {"53", "13"}},
			wantb:    [][]string{{"75", "47", "61", "53", "29"}, {"97", "61", "53", "29", "13"}, {"75", "29", "13"}, {"75", "97", "47", "61", "53"}, {"61", "13", "29"}, {"97", "13", "75", "29", "47"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.filename)
			if err != nil {
				log.Fatal(err)
			}
			gota, gotb := loadFile(file)
			if !reflect.DeepEqual(gota, tt.wanta) {
				t.Errorf("loadFile() got = %v, want %v", gota, tt.wanta)
			}
			if !reflect.DeepEqual(gotb, tt.wantb) {
				t.Errorf("loadFile() got1 = %v, want %v", gotb, tt.wantb)
			}
		})
	}
}

func Test_update_checkRules(t *testing.T) {
	tests := []struct {
		name   string
		update update
		rules  map[int][]rule
		want   bool
	}{
		{
			name:   "valid rule test #1",
			update: update{pages: []int{11, 54, 23, 6}},
			rules: map[int][]rule{
				11: []rule{{11, 54}, {11, 23}},
				54: []rule{{54, 6}},
				23: []rule{{23, 6}},
			},
			want: true,
		}, {
			name:   "rule violation, test #1",
			update: update{pages: []int{11, 54, 23, 6}},
			rules: map[int][]rule{
				11: []rule{{11, 54}, {11, 23}},
				54: []rule{{54, 6}},
				6:  []rule{{6, 23}},
			},
			want: false,
		}, {
			name:   "valid rule test #2",
			update: update{pages: []int{78, 35, 41, 17, 49, 11, 23, 53, 83}},
			rules: map[int][]rule{
				78: {{78, 53}, {78, 23}, {78, 83}, {78, 41}, {78, 17}, {78, 49}, {78, 11}, {78, 35}},
				35: {{35, 53}, {35, 17}, {35, 41}, {35, 83}, {35, 23}, {35, 49}, {35, 11}},
				41: {{41, 23}, {41, 17}, {41, 83}, {41, 11}, {41, 53}, {41, 49}},
				17: {{17, 23}, {17, 53}, {17, 11}, {17, 83}, {17, 49}},
				49: {{49, 53}, {49, 83}, {49, 23}, {49, 11}},
				11: {{11, 53}, {11, 23}, {11, 83}},
				23: {{23, 53}, {23, 83}},
				53: {{53, 83}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.update
			got := u.checkRules(tt.rules)
			assert.Equal(t, tt.want, got)
		})
	}
}
