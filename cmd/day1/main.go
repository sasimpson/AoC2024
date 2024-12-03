package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("day 1")
	fp, err := os.Open("cmd/day1/data1a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	a1, a2 := makeLists(fp)

	fmt.Println("part 1 the sum is", partOne(a1, a2))
	fmt.Println("part 2 the sum is", partTwo(a1, a2))
}

func partOne(a1, a2 []int) int {
	var sum int

	slices.Sort(a1)
	slices.Sort(a2)

	for i := 0; i < len(a1); i++ {
		sum = sum + absInt(a1[i], a2[i])
	}
	return sum
}

func partTwo(a1, a2 []int) int {
	var sum int
	for _, left := range a1 {
		var count int
		for _, right := range a2 {
			if left == right {
				count++ // the left number exists in the right, so count it
			}
		}
		sum = sum + (count * left)
	}
	return sum
}

func makeLists(file io.Reader) ([]int, []int) {
	var a1, a2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		if len(line) == 2 {
			err := appendInt(&a1, line[0])
			if err != nil {
				log.Fatal(err)
			}
			err = appendInt(&a2, line[1])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return a1, a2
}

func absInt(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func appendInt(i *[]int, s string) error {
	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*i = append(*i, v)
	return nil
}
