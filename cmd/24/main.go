package main

import (
	"fmt"
	"slices"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

// part 1 algorithm (assumes weights are unique):
// - get target weight, equal to total weight / 3
// - get all combos that sum to target weight
// - collect all combos with minimum length
// - sort combos first by length, then by entanglement
// - for each one:
// -     filter out group from all weights list
// -     get all combos from the filtered weight list
// -     for each second group, filter remaining weights to get third group, if length of third group > 0, the first is the answer

// for part 2, extend search one loop further

func comboSums(combos *[][]int, weights []int, target int, running int, combo []int, i int) {
	if running == target {
		*combos = append(*combos, slices.Clone(combo))
	} else if running > target {
		return
	}
	for ; i < len(weights); i++ {
		combo = append(combo, weights[i])
		comboSums(combos, weights, target, running+weights[i], combo, i+1)
		combo = combo[:len(combo)-1]
	}
}

func entanglement(combo []int) int {
	ent := 1
	for _, n := range combo {
		ent *= n
	}
	return ent
}

func filterCombo(c []int, weights []int) []int {
	set := make(map[int]struct{})
	for _, n := range c {
		set[n] = struct{}{}
	}
	filtered := make([]int, 0)
	for _, n := range weights {
		if _, ok := set[n]; !ok {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func getCombos(weights []int, target int) [][]int {
	var combos [][]int
	comboSums(&combos, weights, target, 0, make([]int, 0), 0)

	slices.SortFunc(combos, func(c1, c2 []int) int {
		if len(c1) != len(c2) {
			return len(c1) - len(c2)
		}
		return entanglement(c1) - entanglement(c2)
	})

	return combos
}

func part1(weights []int, total int) int {
	target := total / 3
	combos := getCombos(weights, target)

	for _, first := range combos {
		remaining := filterCombo(first, weights)

		secondGroups := make([][]int, 0)
		comboSums(&secondGroups, remaining, target, 0, make([]int, 0), 0)

		for _, second := range secondGroups {
			third := filterCombo(second, remaining)
			if len(third) > 0 {
				return entanglement(first)
			}
		}
	}

	panic("no solution found")
}

func part2(weights []int, total int) int {
	target := total / 4
	combos := getCombos(weights, target)

	for _, first := range combos {
		remaining := filterCombo(first, weights)

		secondGroups := make([][]int, 0)
		comboSums(&secondGroups, remaining, target, 0, make([]int, 0), 0)

		for _, second := range secondGroups {
			remaining2 := filterCombo(second, remaining)

			thirdGroups := make([][]int, 0)
			comboSums(&thirdGroups, remaining2, target, 0, make([]int, 0), 0)

			for _, third := range thirdGroups {
				fourth := filterCombo(third, remaining2)
				if len(fourth) > 0 {
					return entanglement(first)
				}
			}
		}
	}

	panic("no solution found")
}

func main() {
	input := lib.GetInput()

	total := 0
	weights := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		nums := lib.ParseNums(line)
		total += nums[0]
		weights = append(weights, nums[0])
	}

	fmt.Println(part1(weights, total))
	fmt.Println(part2(weights, total))
}
