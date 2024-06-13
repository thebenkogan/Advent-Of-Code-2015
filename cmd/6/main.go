package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func main() {
	input := lib.GetInput()

	on := make(map[[2]int]bool)
	level := make(map[[2]int]int)
	for _, line := range strings.Split(input, "\n") {
		nums := lib.ParseNums(line)
		x1, y1, x2, y2 := nums[0], nums[1], nums[2], nums[3]
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				pair := [2]int{x, y}
				switch {
				case strings.HasPrefix(line, "turn on"):
					on[pair] = true
					level[pair]++
				case strings.HasPrefix(line, "turn off"):
					delete(on, pair)
					level[pair] = max(level[pair]-1, 0)
				case strings.HasPrefix(line, "toggle"):
					if on[pair] {
						delete(on, pair)
					} else {
						on[pair] = true
					}
					level[pair] += 2
				}
			}
		}
	}

	fmt.Println(len(on))

	total := 0
	for _, l := range level {
		total += l
	}
	fmt.Println(total)
}
