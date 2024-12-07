package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var searchBlob = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

//var searchBlob = `S..S..S
//.A.A.A.
//..MMM..
//SAMXMAS
//..MMM..
//.A.A.A.
//S..S..S`

func Test_searchPosition(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(searchBlob))
	var data [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		data = append(data, line)
	}
	var count int
	for y := 0; y < len(data); y++ {
		//fmt.Println("Checking line", y)
		for x := 0; x < len(data[y]); x++ {
			//fmt.Println("Checking position", y, x, data[y][x])
			if data[y][x] == "X" {
				count += searchPosition(data, "XMAS", x, y)
			}
		}
	}
	assert.Equal(t, 18, count)
}

func Test_searchMas(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(searchBlob))
	var data [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		data = append(data, line)
	}
	var count int
	for y := 0; y < len(data); y++ {
		//fmt.Println("Checking line", y)
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == "A" {
				//fmt.Println("found A at", x, y)
				if LRMas(data, x, y, 1) {
					count++
				}
				if LRMas(data, x, y, -1) {
					count++
				}
			}
		}
	}
	assert.Equal(t, 9, count)
}
