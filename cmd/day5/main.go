package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type rule struct {
	a int
	b int
}

type update struct {
	pages  []int
	status bool
}

func (u update) checkRules(rules map[int][]rule) bool {
	for _, currentPage := range u.pages {
		aidx := slices.Index(u.pages, currentPage)
		if aidx == -1 {
			fmt.Println("a value not found, returning true")
			continue
		}
		//if the page is not in the rules, we're good for this one
		//if any rule is violated, this is false.
		for _, rule := range rules[currentPage] {
			bidx := slices.Index(u.pages, rule.b)
			if bidx == -1 {
				continue
			}
			// if the b in the rule is not in the pages, move on.
			//fmt.Printf("a: %v, b: %v, currentPage: %d, rule: %v, update: %v\n", aidx, bidx, currentPage, rule, u)
			if aidx > bidx {
				return false
			}
		}
	}

	return true
}

func (u update) middlePage() int {
	return u.pages[len(u.pages)/2]
}

func main() {
	fp, err := os.Open("cmd/day5/data5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	rules, data := loadFile(fp)

	ruleChart := loadRules(rules)
	updates := loadUpdates(data)

	for i, u := range updates {
		if u.checkRules(ruleChart) {
			updates[i].status = true
			fmt.Println(u)
		}
	}
	var count int
	var sum int
	for _, u := range updates {
		if u.status {
			//fmt.Println("passed", u)
			count++
			sum += u.middlePage()
		}
	}
	fmt.Println("number of valid updates", count)
	fmt.Println("middle page sum", sum)
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
