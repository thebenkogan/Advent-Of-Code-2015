package main

import (
	"encoding/json"
	"fmt"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func getDocTotal(doc interface{}) int {
	total := 0
	switch doc := doc.(type) {
	case map[string]interface{}:
		for _, v := range doc {
			total += getDocTotal(v)
		}
	case []interface{}:
		for _, v := range doc {
			total += getDocTotal(v)
		}
	case float64:
		total += int(doc)
	}
	return total
}

func getDocTotalIgnoreRed(doc interface{}) int {
	total := 0
	switch doc := doc.(type) {
	case map[string]interface{}:
		hasRed := false
		for _, v := range doc {
			if v == "red" {
				hasRed = true
				break
			}
		}
		if hasRed {
			break
		}
		for _, v := range doc {
			total += getDocTotalIgnoreRed(v)
		}
	case []interface{}:
		for _, v := range doc {
			total += getDocTotalIgnoreRed(v)
		}
	case float64:
		total += int(doc)
	}
	return total
}

func main() {
	input := lib.GetInput()

	var doc interface{}
	_ = json.Unmarshal([]byte(input), &doc)

	fmt.Println(getDocTotal(doc))
	fmt.Println(getDocTotalIgnoreRed(doc))
}
