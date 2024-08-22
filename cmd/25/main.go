package main

import (
	"fmt"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func main() {
	input := lib.GetInput()

	nums := lib.ParseNums(input)
	targetRow := nums[0]
	targetCol := nums[1]

	prev := 20151125
	for startRow := 2; ; startRow++ {
		row := startRow
		col := 1
		for row > 0 {
			next := (prev * 252533) % 33554393
			if row == targetRow && col == targetCol {
				fmt.Println(next)
				return
			}
			prev = next
			row--
			col++
		}
	}
}
