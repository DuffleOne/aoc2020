package shared

import (
	"encoding/json"
	"fmt"

	"aoc2020/cher"
)

func HandleErr(err error) {
	if err == nil {
		return
	}

	if v, ok := err.(cher.E); ok {
		bytes, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			panic(err)
		}

		fmt.Println("Handled error:")
		fmt.Println(string(bytes))
		panic(err)
	}

	panic(err)
}
