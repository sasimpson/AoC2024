package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fp, err := os.Open("cmd/day5/data5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var rules [][]string
	var updates [][]string
	for scanner.Scan() {
		line := scanner.Bytes()
		if ok, err := regexp.Match(`\|`, line); ok && err == nil {
			rules = append(rules, strings.Split(string(line), "|"))
		}
		if ok, err := regexp.Match(`,`, line); ok && err == nil {
			updates = append(updates, strings.Split(string(line), ","))
		}
	}

	//for _, r := range rules {
	//	fmt.Println(r)
	//}
	ruleChart := loadRules(rules)
	fmt.Println(ruleChart[22])
	for _, update := range updates {
		fmt.Println(update)
	}
}

type rule struct {
	a int
	b int
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
