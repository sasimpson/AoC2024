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

	var safe int
	var damper int
	for _, report := range reports {
		if isSafe(report) {
			safe++
			continue
		}

		if bruteForceDamper(report) {
			damper++
			continue
		}

	}
	fmt.Println("day 2 part 1: ", safe)
	fmt.Println("day 2 part 2: ", damper+safe)
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

func avg(vals []int) int {
	var sum int
	for _, v := range vals {
		sum = sum + v
	}
	return sum / len(vals)
}

func validateRisk(a, b int, inc bool) bool {
	if inc {
		return (b-a <= 3) && (b-a > 0)
	}
	return (a-b <= 3) && (a-b > 0)
}

func isSafe(report []int) bool {
	increasing := report[0] < avg(report)
	for i := 1; i < len(report); i++ {
		valid := validateRisk(report[i-1], report[i], increasing)
		if !valid {
			return false
		}
	}
	return true
}

func fix(pos int, report []int) []int {
	newReport := make([]int, len(report)-1)
	copy(newReport[:pos], report[:pos])
	copy(newReport[pos:], report[pos+1:])
	return newReport
}

func bruteForceDamper(report []int) bool {
	for i := 0; i < len(report); i++ {
		if isSafe(fix(i, report)) {
			return true
		}
	}
	return false
}
