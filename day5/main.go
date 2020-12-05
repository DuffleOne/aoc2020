package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"

	"aoc2020/cher"
	"aoc2020/shared"
)

var re = regexp.MustCompile(`(?m)^[F|B]{7}[R|L]{3}$`)

var reversePowersCols = []int{6, 5, 4, 3, 2, 1, 0}
var reversePowersRows = []int{0, 0, 0, 0, 0, 0, 3, 2, 1, 0}

func main() {
	lines, err := parseFile()
	if err != nil {
		panic(err)
	}

	var highestSeatID int

	var seatAllocations []int

	for _, l := range lines {
		parts := strings.Split(l, "")

		col := struct {
			Min int
			Max int
		}{0, 127}

		row := struct {
			Min int
			Max int
		}{0, 7}

		for i, p := range parts {
			switch p {
			case "F":
				col.Max -= int(math.Pow(float64(2), float64(reversePowersCols[i])))
			case "B":
				col.Min += int(math.Pow(float64(2), float64(reversePowersCols[i])))
			case "R":
				row.Min += int(math.Pow(float64(2), float64(reversePowersRows[i])))
			case "L":
				row.Max -= int(math.Pow(float64(2), float64(reversePowersRows[i])))
			}
		}

		if col.Min != col.Max {
			panic(cher.New("col_mismatch", cher.M{"col": col, "row": row}))
		}

		if row.Min != row.Max {
			panic(cher.New("row_mismatch", cher.M{"col": col, "row": row}))
		}

		seatID := col.Min*8 + row.Min

		seatAllocations = append(seatAllocations, seatID)

		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	fmt.Print("Highest seat: ")
	fmt.Println(highestSeatID)

	sort.Ints(seatAllocations)

	var mySeatID int
	start := seatAllocations[0]

	for _, n := range seatAllocations {
		if start != n {
			mySeatID = start
			break
		}

		start++
	}

	fmt.Print("Your seat: ")
	fmt.Println(mySeatID)
}

func parseFile() ([]string, error) {
	var set []string

	lines, err := shared.ReadLines("./input.txt")
	if err != nil {
		return nil, err
	}

	for _, l := range lines {
		if !re.MatchString(l) {
			return nil, cher.New("invalid_line", cher.M{
				"line": l,
			})
		}

		set = append(set, l)
	}

	return set, nil
}
