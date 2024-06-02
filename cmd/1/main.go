package main

import (
	"fmt"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func main() {
	input := lib.GetInput()

	total := 0
	basementPos := 0
	for i, c := range input {
		if c == '(' {
			total++
		} else if c == ')' {
			total--
		}
		if total == -1 && basementPos == 0 {
			basementPos = i + 1
		}
	}

	fmt.Println(total)
	fmt.Println(basementPos)
}
