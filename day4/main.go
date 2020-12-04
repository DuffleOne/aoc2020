package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"aoc2020/cher"
	"aoc2020/ptr"
	"aoc2020/shared"
)

var (
	re       = regexp.MustCompile(`(?m)([a-z]+):([^\s]*)`)
	digit4   = regexp.MustCompile(`^([0-9]{4})$`)
	hgt      = regexp.MustCompile(`^(\d*)([a-z]{2})$`)
	hex      = regexp.MustCompile(`^#[a-f0-9]*$`)
	passport = regexp.MustCompile(`^\d{9}$`)
)

func main() {
	set, err := parseFile()
	if err != nil {
		panic(err)
	}

	var validCount int
	for _, s := range set {
		if s.IsValid() {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func parseFile() ([]*Doc, error) {
	lines, err := shared.ReadLines("./input.txt")
	if err != nil {
		return nil, err
	}

	var set []*Doc

	var index *Doc
	for _, l := range lines {
		if len(l) == 0 {
			set = append(set, index)
			index = nil
			continue
		}

		if index == nil {
			index = &Doc{}
		}

		matches := re.FindAllStringSubmatch(l, -1)

		for _, m := range matches {
			if len(m) != 3 {
				return nil, cher.New("cannot_parse", cher.M{
					"line": l,
				})
			}

			err = index.AddElement(m[1], m[2])
			if err != nil {
				return nil, err
			}
		}
	}

	set = append(set, index)

	return set, nil
}

type Doc struct {
	ECL *string
	PID *string
	EYR *string
	HCL *string
	BYR *string
	IYR *string
	CID *string
	HGT *string
}

func (d *Doc) IsValid() bool {
	r := reflect.ValueOf(d).Elem()

	for i := 0; i < r.NumField(); i++ {
		k := r.Type().Field(i).Name
		v := r.Field(i).Interface().(*string)

		if k == "CID" {
			continue
		}

		if v == nil {
			return false
		}
	}

	return true
}

func (d *Doc) AddElement(key, value string) error {
	key = strings.ToLower(key)

	switch key {
	case "ecl":
		d.addECL(value)
	case "pid":
		d.addPID(value)
	case "eyr":
		d.addEYR(value)
	case "hcl":
		d.addHCL(value)
	case "byr":
		d.addBYR(value)
	case "iyr":
		d.addIYR(value)
	case "cid":
		d.CID = ptr.String(value)
	case "hgt":
		d.addHGT(value)
	default:
		return cher.New("unknown_key", cher.M{"key": key, "value": value})
	}

	return nil
}

func (d *Doc) addBYR(value string) {
	if !digit4.MatchString(value) {
		return
	}

	n, err := strconv.Atoi(value)
	if err != nil {
		return
	}

	if n < 1920 || n > 2002 {
		return
	}

	d.BYR = ptr.String(value)
}

func (d *Doc) addIYR(value string) {
	if !digit4.MatchString(value) {
		return
	}

	n, err := strconv.Atoi(value)
	if err != nil {
		return
	}

	if n < 2010 || n > 2020 {
		return
	}

	d.IYR = ptr.String(value)
}

func (d *Doc) addEYR(value string) {
	if !digit4.MatchString(value) {
		return
	}

	n, err := strconv.Atoi(value)
	if err != nil {
		return
	}

	if n < 2020 || n > 2030 {
		return
	}

	d.EYR = ptr.String(value)
}

func (d *Doc) addHGT(value string) {
	m := hgt.FindStringSubmatch(value)
	if len(m) != 3 {
		return
	}

	v, err := strconv.Atoi(m[1])
	if err != nil {
		return
	}

	switch m[2] {
	case "cm":
		if v < 150 || v > 193 {
			return
		}
	case "in":
		if v < 59 || v > 76 {
			return
		}
	}

	d.HGT = ptr.String(value)
}

func (d *Doc) addHCL(value string) {
	if !hex.MatchString(value) {
		return
	}

	d.HCL = ptr.String(value)
}

func (d *Doc) addECL(value string) {
	allowed := []string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}

	if !contains(allowed, value) {
		return
	}

	d.ECL = ptr.String(value)
}

func (d *Doc) addPID(value string) {
	if !passport.MatchString(value) {
		return
	}

	d.PID = ptr.String(value)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
