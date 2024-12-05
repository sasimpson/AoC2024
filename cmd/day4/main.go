package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	fp, err := os.Open("cmd/day4/data4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
}
