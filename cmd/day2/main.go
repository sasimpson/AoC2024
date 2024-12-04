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
	fp, err := os.Open("data2a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	reports := parseData(fp)

	var safe int
	var rescan int
	for _, report := range reports {
		i := isSafe(report)
		if i == -1 {
			safe++
		} else {
			dampedReport := make([]int, len(report)-1)
			copy(dampedReport, report[:i])
			copy(dampedReport, report[i:])
			j := isSafe(dampedReport)
			if j == -1 {
				rescan++
			} else {
				fmt.Println(report, dampedReport)
			}
		}

	}

	fmt.Println("day 2 part 1: ", safe)
	fmt.Println("day 2 part 2: ", rescan+safe)
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

func isSafe(report []int) int {
	increasing := report[0] < avg(report)

	for i := 1; i < len(report); i++ {

		valid := validateRisk(report[i-1], report[i], increasing)

		if !valid {
			return i
		}
	}

	return -1
}
