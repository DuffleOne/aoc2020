package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"aoc2020/shared"
)

const goal = "shiny gold"

var rootReg = regexp.MustCompile(`^([a-z\s]+)\sbags contain\s`)
var subReg = regexp.MustCompile(`^(\d)\s([a-z\s]+)\sbags?.?$`)

func main() {
	part1()
	part2()
}

func part1() {
	bags, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	set := []string{goal}

beginning:
	for _, bag := range bags {
		for _, children := range bag.Children {
			if contains(set, children.Color) && !contains(set, bag.Color) {
				set = append(set, bag.Color)
				goto beginning
			}
		}
	}

	fmt.Println(len(set) - 1)
}

func part2() {
	bags, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	m := map[string]map[string]int{}

	for _, bag := range bags {
		children := map[string]int{}

		for _, child := range bag.Children {
			children[child.Color] = child.Count
		}

		m[bag.Color] = children
	}

	fmt.Println(countChildren(m, goal))
}

func countChildren(m map[string]map[string]int, str string) int {
	children := m[str]

	var total int

	for child, count := range children {
		total += count * (1 + countChildren(m, child))
	}

	return total
}

func parse() ([]*Bag, error) {
	lines, err := shared.ReadLines("./input.txt")
	if err != nil {
		return nil, err
	}

	var bags []*Bag

	for _, line := range lines {
		bags = append(bags, parseLine(line))
	}

	return bags, nil
}

func parseLine(line string) *Bag {
	matches := rootReg.FindAllStringSubmatch(line, -1)

	bagColor := matches[0][1]

	rest := strings.TrimPrefix(line, matches[0][0])
	parsedRest := parseRest(rest)

	return &Bag{
		Count:    1,
		Color:    bagColor,
		Children: parsedRest,
	}
}

func parseRest(line string) []*Bag {
	parts := strings.Split(line, ", ")

	set := []*Bag{}

	for _, part := range parts {
		matches := subReg.FindAllStringSubmatch(part, -1)

		if len(matches) == 0 {
			return nil
		}

		n, err := strconv.Atoi(matches[0][1])
		if err != nil {
			panic(err)
		}

		set = append(set, &Bag{Count: n, Color: matches[0][2]})
	}

	return set
}

type Bag struct {
	Count    int
	Color    string
	Children []*Bag
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
