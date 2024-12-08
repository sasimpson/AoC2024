package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fp, err := os.Open("cmd/day4/data4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var data [][]string
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "")
		data = append(data, splitLine)
	}

	part1(data)
	part2(data)
}

func part1(data [][]string) {
	var count int
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[0]); x++ {
			if data[y][x] == "X" {
				count += searchPosition(data, "XMAS", x, y)
			}
		}
	}
	fmt.Println("day 4 part 1", count)
}

func part2(data [][]string) {
	var count int
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[0]); x++ {
			if data[y][x] == "A" {
				if LRMas(data, x, y, 1) {
					count++
				}
				if LRMas(data, x, y, -1) {
					count++
				}
				if TBMas(data, x, y, 1) {
					count++
				}
				if TBMas(data, x, y, -1) {
					count++
				}
			}
		}
	}
	fmt.Println("day 4 part 2", count)
}

// LRMas - Left to Right MAS
//
//	0 1 2
//
// 0 M . S
// 1 . A .
// 2 M . S
// A's (1 <= x pos < len(data[y]) and (1 <= y pos < len(data))
// M's in (x-1,y-1) and (x-1,y+1)
func LRMas(data [][]string, x, y int, dir int) bool {
	if (x >= 1 && x < len(data[y])-1) && (y >= 1 && y < len(data)-1) {
		var modx, mody []int
		switch {
		case dir > 0:
			modx = []int{-1, -1}
			mody = []int{-1, 1}
		case dir < 0:
			modx = []int{1, 1}
			mody = []int{1, -1}
		}
		if searchDiagonalDown(data, "MAS", x+modx[0], y+mody[0], dir) &&
			searchDiagonalUp(data, "MAS", x+modx[1], y+mody[1], dir) {
			return true
		}
	}
	return false
}

// TBMas - Top to Bottom MAS
//
//	0 1 2
//
// 0 M . M
// 1 . A .
// 2 S . S
// A's (1 <=x pos < len(data[y]) and (1 <= y pos < len(data))
// M's in DD(x-1, y-1, +dir) and DU(x+1, y-1 -dir)
func TBMas(data [][]string, x, y int, dir int) bool {
	if (x >= 1 && x < len(data[y])-1) && (y >= 1 && y < len(data)-1) {
		var modx, mody []int
		switch {
		case dir > 0: //top to bottom
			modx = []int{x - 1, x + 1}
			mody = []int{y - 1, y - 1}
		case dir < 0: //bottom to top
			modx = []int{x + 1, x - 1}
			mody = []int{y + 1, y + 1}
		}
		if searchDiagonalDown(data, "MAS", modx[0], mody[0], dir) &&
			searchDiagonalUp(data, "MAS", modx[1], mody[1], -dir) {
			return true
		}
	}
	return false
}

func searchPosition(data [][]string, search string, x, y int) int {
	var count int
	if searchHorizontal(data, search, x, y, 1) {
		//fmt.Println("found H+ at", y, x)
		count++
	}
	if searchHorizontal(data, search, x, y, -1) {
		//fmt.Println("found H- at", y, x)
		count++
	}
	if searchVertical(data, search, x, y, 1) {
		//fmt.Println("found V+ at", y, x)
		count++
	}
	if searchVertical(data, search, x, y, -1) {
		//fmt.Println("found V- at", y, x)
		count++
	}
	if searchDiagonalDown(data, search, x, y, 1) {
		//fmt.Println("found DD+ at", y, x)
		count++
	}
	if searchDiagonalDown(data, search, x, y, -1) {
		//fmt.Println("found DD- at", y, x)
		count++
	}
	if searchDiagonalUp(data, search, x, y, 1) {
		//fmt.Println("found UD+ at", y, x)
		count++
	}
	if searchDiagonalUp(data, search, x, y, -1) {
		//fmt.Println("found UD- at", y, x)
		count++
	}
	return count
}

func searchHorizontal(data [][]string, search string, x, y, dir int) bool {
	searchLen := len(search)
	s := make([]string, searchLen)
	//if going forward, check that we don't run over the end
	if dir >= 1 {
		if len(data[y]) > x+(searchLen-1) {
			for i := 0; i < searchLen; i++ {
				s[i] = data[y][x+i]
			}
		}
	} else {
		if x >= searchLen-1 {
			for i := 0; i < searchLen; i++ {
				s[i] = data[y][x-i]
			}
		}
	}
	toMatch := strings.Join(s, "")
	if toMatch == search {
		return true
	}
	return false
}

func searchVertical(data [][]string, search string, x, y, dir int) bool {
	searchLen := len(search)
	s := make([]string, searchLen)
	if dir >= 1 {
		if len(data) > y+(searchLen-1) {
			for i := 0; i < searchLen; i++ {
				s[i] = data[y+i][x]
			}
		}
	} else {
		if y >= (searchLen - 1) {
			for i := 0; i < searchLen; i++ {
				s[i] = data[y-i][x]
			}
		}
	}
	toMatch := strings.Join(s, "")
	if toMatch == search {
		return true
	}
	return false
}

func searchDiagonalDown(data [][]string, search string, x, y, dir int) bool {
	searchLen := len(search)
	lenX := len(data[y])
	lenY := len(data)
	s := make([]string, searchLen)
	if dir >= 1 {
		if lenX > x+(searchLen-1) && lenY > y+searchLen-1 {
			for i := 0; i < searchLen; i++ {
				s[i] = data[y+i][x+i]
			}
		}
	} else {
		if x >= searchLen-1 && y >= searchLen-1 {
			for i := 0; i < searchLen; i++ {
				s[i] = data[y-i][x-i]
			}
		}
	}
	toMatch := strings.Join(s, "")
	if toMatch == search {
		return true
	}
	return false
}

func searchDiagonalUp(data [][]string, search string, x, y, dir int) bool {
	searchLen := len(search)
	s := make([]string, searchLen)
	if dir >= 1 {
		if len(data[y]) > x+(searchLen-1) && y >= searchLen-1 {
			for i := 0; i < searchLen; i++ {
				s[i] = data[y-i][x+i]
			}
		}
	} else {
		if x >= searchLen-1 && len(data) > y+(searchLen-1) {
			for i := 0; i < searchLen; i++ {
				s[i] = data[y+i][x-i]
			}
		}
	}
	toMatch := strings.Join(s, "")
	if toMatch == search {
		return true
	}
	return false
}
