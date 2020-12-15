package shared

import (
	"strconv"
	"strings"
)

type ReadType int

const (
	NewLine ReadType = iota
	Comma
)

func ReadNumbers(fileName string, readType ReadType) (s []int, err error) {
	lines, err := ReadLines(fileName)
	if err != nil {
		return nil, err
	}

REDO:
	switch readType {
	case NewLine:
		for _, line := range lines {
			n, err := strconv.Atoi(line)
			if err != nil {
				return nil, err
			}

			s = append(s, n)
		}
	case Comma:
		lines = strings.Split(lines[0], ",")
		readType = NewLine
		goto REDO
	}

	return
}
