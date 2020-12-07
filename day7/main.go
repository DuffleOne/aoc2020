package main

import (
	"fmt"
	"regexp"
	"strings"

	"aoc2020/shared"
)

const goal = "shiny gold"

var rootReg = regexp.MustCompile(`^([a-z\s]+)\sbags contain\s`)
var subReg = regexp.MustCompile(`^(\d)\s([a-z\s]+)\sbags?.?$`)

func main() {
	bags, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	set := []string{goal}

beginning:
	for _, bag := range bags {
		for _, children := range bag.Children {
			if contains(set, children) && !contains(set, bag.Color) {
				set = append(set, bag.Color)
				goto beginning
			}
		}
	}

	fmt.Println(len(set) - 1)
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
		Color:    bagColor,
		Children: parsedRest,
	}
}

func parseRest(line string) []string {
	parts := strings.Split(line, ", ")

	set := []string{}

	for _, part := range parts {
		matches := subReg.FindAllStringSubmatch(part, -1)

		if len(matches) == 0 {
			return []string{}
		}

		set = append(set, matches[0][2])
	}

	return set
}

type Bag struct {
	Color    string
	Children []string
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
