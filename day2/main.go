package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"aoc2020/cher"
	"aoc2020/shared"
)

var re = regexp.MustCompile(`^(\d+)-(\d+)\s([a-z0-9]):\s([a-z]*)$`)

func main() {
	part2()
}

func part1() {
	input, err := readFile()
	if err != nil {
		shared.HandleErr(err)
	}

	var validCount = 0

	for _, i := range input {
		actualCount := getLetterCount(i.Input, i.Letter)
		if actualCount >= i.Min && actualCount <= i.Max {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func part2() {
	input, err := readFile()
	if err != nil {
		shared.HandleErr(err)
	}

	var validCount = 0

	for _, i := range input {
		X := rune(i.Input[i.Min-1]) == i.Letter
		Y := rune(i.Input[i.Max-1]) == i.Letter

		// bitwise XOR
		if (X || Y) && !(X && Y) {
			validCount++
		}
	}

	fmt.Println(validCount)
}

type PasswordRule struct {
	Input  string
	Min    int
	Max    int
	Letter rune
}

func (pr PasswordRule) Validate() error {
	return errors.New("eof")
}

func readFile() ([]PasswordRule, error) {
	lines, err := shared.ReadLines("./input.txt")
	if err != nil {
		panic(err)
	}

	var out []PasswordRule

	for _, l := range lines {
		m := re.FindStringSubmatch(l)

		if len(m) != 5 {
			return nil, cher.New("invalid_input", cher.M{
				"m":    m,
				"line": l,
			})
		}

		min, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}

		max, err := strconv.Atoi(m[2])
		if err != nil {
			return nil, err
		}

		out = append(out, PasswordRule{
			Min:    min,
			Max:    max,
			Letter: rune(m[3][0]),
			Input:  m[4],
		})
	}

	return out, nil
}

func getLetterCount(input string, char rune) int {
	var c int

	for _, r := range input {
		if r == char {
			c++
		}
	}

	return c
}
