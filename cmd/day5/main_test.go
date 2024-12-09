package main

import (
	"log"
	"os"
	"reflect"
	"testing"
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
				{[]int{75, 47, 61, 53, 29}},
				{[]int{97, 61, 53, 29, 13}},
				{[]int{75, 29, 13}},
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
