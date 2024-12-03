package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var a1, a2 []int
	var sum int

	fp, err := os.Open("cmd/day1/data1a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		if len(line) == 2 {
			a, err := strconv.Atoi(line[0])
			if err != nil {
				log.Fatal(err)
			}
			b, err := strconv.Atoi(line[1])
			if err != nil {
				log.Fatal(err)
			}
			a1 = append(a1, a)
			a2 = append(a2, b)
		}
	}

	slices.Sort(a1)
	slices.Sort(a2)

	for i := 0; i < len(a1); i++ {
		sum = sum + absInt(a1[i], a2[i])
	}

	fmt.Println("the sum is", sum)
}

func absInt(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
