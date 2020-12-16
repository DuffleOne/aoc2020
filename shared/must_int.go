package shared

import (
	"strconv"

	"aoc2020/cher"
)

func MustInt(i string) int {
	n, err := strconv.Atoi(i)
	if err != nil {
		HandleErr(cher.New("not_int", cher.M{"string": i}, cher.Coerce(err)))
	}

	return n
}
