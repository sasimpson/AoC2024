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
		data = append(lines, splitLine)
	}

	//patterns for search:
	/*
		XMAS 	data[i][j:j+4]
		SAMX

		X S		data[i:i+4][j]
		M A
		A M
		S X

		X S		data[i:i+4][j:j+4]
		 M A
		  A M
		   S X

		   S X	data[i:i+4][j+4:j]
		  A M
		 M A
		X S
	*/

	fmt.Println(lines)
}

func searchForward(data *[][]string, search string, x, y int) bool {
	s := data[y][x : x+4]
	return s == search
}
