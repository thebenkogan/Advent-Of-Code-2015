package main

import (
	"fmt"
	"slices"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func main() {
	input := lib.GetInput()

	total := 0
	ribbon := 0
	for _, line := range strings.Split(input, "\n") {
		nums := lib.ParseNums(line)
		slices.Sort(nums)
		l, w, h := nums[0], nums[1], nums[2]
		total += 2*l*w + 2*w*h + 2*h*l + l*w
		ribbon += 2*l + 2*w + l*w*h
	}

	fmt.Println(total)
	fmt.Println(ribbon)
}
