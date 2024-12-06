package main

import (
	"bufio"
	"fmt"
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

func Test_searchForward(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(searchBlob))
	var data [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		data = append(data, line)
	}

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == "X" {
				if len(data[y]) > x+4 {
					if searchForward(data, "XMAS", x, y) {
						fmt.Println("found")
					}
					if searchForward(data, "SAMX", x, y) {
						fmt.Println("found")
					}
				}
			}
		}
	}

}
