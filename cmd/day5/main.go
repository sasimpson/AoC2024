package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	a int
	b int
}

type update struct {
	pages []int
}

func (u update) checkRules(rules map[int][]rule) {
	fmt.Println("Checking rules for", u)
	for _, page := range u.pages {
		fmt.Printf("%#v\n", rules[page])
	}
	fmt.Println()
	fmt.Println()
}

func main() {
	fp, err := os.Open("cmd/day5/data5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	rules, data := loadFile(fp)

	ruleChart := loadRules(rules)
	//fmt.Println(ruleChart[22])
	updates := loadUpdates(data)
	//fmt.Println(updates)

	for _, u := range updates {
		u.checkRules(ruleChart)
	}
}

func loadFile(file io.Reader) ([][]string, [][]string) {
	scanner := bufio.NewScanner(file)
	var rules [][]string
	var data [][]string
	for scanner.Scan() {
		line := scanner.Bytes()
		if ok, err := regexp.Match(`\|`, line); ok && err == nil {
			rules = append(rules, strings.Split(string(line), "|"))
		}
		if ok, err := regexp.Match(`,`, line); ok && err == nil {
			data = append(data, strings.Split(string(line), ","))
		}
	}
	return rules, data
}

func loadRules(rules [][]string) map[int][]rule {
	ruleChart := make(map[int][]rule)
	for _, r := range rules {
		a, _ := strconv.Atoi(r[0])
		b, _ := strconv.Atoi(r[1])

		x := rule{a: a, b: b}
		if _, ok := ruleChart[a]; !ok {
			ruleChart[a] = make([]rule, 0)
		}
		ruleChart[a] = append(ruleChart[a], x)
	}
	return ruleChart
}

func loadUpdates(data [][]string) []update {
	var updates []update
	for _, line := range data {
		var u update
		for _, x := range line {
			page, _ := strconv.Atoi(x)
			u.pages = append(u.pages, page)
		}
		updates = append(updates, u)
	}

	return updates
}
