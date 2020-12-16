package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"aoc2020/cher"
	"aoc2020/shared"

	"github.com/kr/pretty"
)

var ruleRegex = regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)

func main() {
	part1()
	part2()
}

func part1() {
	rules, _, nearby, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	var allInvalid []int

	for _, ticket := range nearby {
		allInvalid = append(allInvalid, rules.GetInvalidItems(ticket)...)
	}

	var sum int

	for _, n := range allInvalid {
		sum += n
	}

	fmt.Println(sum)
}

func part2() {
	rules, _, nearby, err := parse()
	if err != nil {
		shared.HandleErr(err)
	}

	var validTickets [][]int

	for _, ticket := range nearby {
		if rules.IsValid(ticket) {
			validTickets = append(validTickets, ticket)
		}
	}

	ruleMapping := map[int]map[string]struct{}{}

	for i := 0; i < len(validTickets[0]); i++ {
		set := map[string]struct{}{}

		for _, rule := range rules {
			set[rule.Name] = struct{}{}
		}

		ruleMapping[i] = set
	}

	for i := 0; i < len(validTickets[0]); i++ {
		for _, ticket := range validTickets {
			num := ticket[i]

			for _, rule := range rules {
				if !rule.Validate(num) {
					delete(ruleMapping[i], rule.Name)

					if len(ruleMapping[i]) == 1 {
						lastOne := getLastItem(ruleMapping[i])

						for j := 0; j < len(validTickets[0]); j++ {
							if i == j {
								continue
							}

							delete(ruleMapping[j], lastOne)
						}
					}
				}
			}
		}
	}

	pretty.Println(ruleMapping)
}

func getLastItem(in map[string]struct{}) string {
	if len(in) != 1 {
		panic("wat")
	}

	for k := range in {
		return k
	}

	panic("wat 2")
}

func parse() (Rules, []int, [][]int, error) {
	groups, err := shared.ReadGroups("./input.txt")
	if err != nil {
		return nil, nil, nil, err
	}

	var rawRules, rawTicket, rawNearby = groups[0], groups[1], groups[2]

	rules, err := parseRules(rawRules)
	if err != nil {
		return nil, nil, nil, err
	}

	var ticket []int

	for _, i := range strings.Split(rawTicket[1], ",") {
		ticket = append(ticket, shared.MustInt(i))
	}

	var nearby [][]int

	for _, t := range rawNearby[1:] {
		var s []int

		for _, u := range strings.Split(t, ",") {
			s = append(s, shared.MustInt(u))
		}

		nearby = append(nearby, s)
	}

	return rules, ticket, nearby, nil
}

func parseRules(rawRules []string) (Rules, error) {
	rules := Rules{}

	for _, rr := range rawRules {
		matches := ruleRegex.FindStringSubmatch(rr)
		if len(matches) != 6 {
			return nil, cher.New("invalid_line", cher.M{"line": rr, "matches": matches})
		}

		range1 := []int{shared.MustInt(matches[2]), shared.MustInt(matches[3])}
		range2 := []int{shared.MustInt(matches[4]), shared.MustInt(matches[5])}

		sort.Ints(range1)
		sort.Ints(range2)

		r := Rule{
			Name:   matches[1],
			Ranges: [][]int{range1, range2},
		}

		rules = append(rules, r)
	}

	return rules, nil
}

type Rules []Rule

type Rule struct {
	Name   string
	Ranges [][]int
}

// IsValid returns false if any items don't match at least 1 rule
func (r Rules) IsValid(ticket []int) bool {
	for _, number := range ticket {
		if r.Validate(number) {
			return true
		}
	}

	return false
}

func (r Rules) GetInvalidItems(ticket []int) (ret []int) {
	for _, i := range ticket {
		if !r.Validate(i) {
			ret = append(ret, i)
		}
	}

	return
}

func (r Rules) Validate(in int) bool {
	for _, rule := range r {
		if rule.Validate(in) {
			return true
		}
	}

	return false
}

func (r Rule) Validate(i int) bool {
	if i >= r.Ranges[0][0] && i <= r.Ranges[0][1] {
		return true
	}

	if i >= r.Ranges[1][0] && i <= r.Ranges[1][1] {
		return true
	}

	return false
}
