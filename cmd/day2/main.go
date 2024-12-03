package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fp, err := os.Open("cmd/day2/data2a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	reports := parseData(fp)

	for _, report := range reports {
		fmt.Println(isSafe(report))
	}

}

func parseData(file io.Reader) [][]int {
	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dataLine := make([]int, len(line))
		for i, v := range line {
			dataLine[i], _ = strconv.Atoi(v)
		}
		data = append(data, dataLine)
	}
	return data
}

func isSafe(report []int) bool {
	if report[0] < report[1] { //increasing
		for i := range report {
			switch {
			case i >= len(report)-1: // at the end, done, safe.
				return true
			case report[i+1]-report[i] <= 0: // no change or decreasing is unsafe
				return false
			case (report[i+1]-report[i] <= 2) && (report[i+1]-report[i] > 0):
				continue //change 1 or 2 is safe, keep going
			default:
				return false
			}
		}
	}
	// decreasing
	for i := range report {
		switch {
		case i >= len(report)-1: // at the end, done, safe.
			return true
		case report[i]-report[i+1] <= 0: // no change or increasing is unsafe
			return false
		case (report[i]-report[i+1] <= 2) && (report[i]-report[i+1] > 0):
			continue //change of 1 or 2 is safe, keep going
		default:
			return false
		}
	}
	return false
}
