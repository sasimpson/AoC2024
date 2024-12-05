package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	mulRegExp = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
)

func main() {
	fp, err := os.Open("cmd/day3/data3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	var sum int64
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		data := scanner.Bytes()

		matches := mulRegExp.FindAllSubmatch(data, -1)
		for _, match := range matches {
			a, _ := strconv.ParseInt(string(match[1]), 10, 64)
			b, _ := strconv.ParseInt(string(match[2]), 10, 64)
			prod := a * b
			sum = sum + prod
		}
	}
	fmt.Println("day 3 sum: ", sum)

}
