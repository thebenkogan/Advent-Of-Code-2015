package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func bestHappiness(pairs map[string]map[string]int) int {
	names := make([]string, 0, len(pairs))
	for name := range pairs {
		names = append(names, name)
	}

	best := 0
	for _, perm := range lib.Permutations(names) {
		happiness := 0
		for i, name := range perm {
			right := perm[(i+1)%len(perm)]
			happiness += pairs[name][right]
			happiness += pairs[right][name]
		}
		best = max(best, happiness)
	}

	return best
}

func main() {
	input := lib.GetInput()

	pairs := make(map[string]map[string]int)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		name := split[0]
		other := split[len(split)-1]
		other = other[:len(other)-1]
		happiness := lib.ParseNums(line)[0]
		if strings.Contains(line, "lose") {
			happiness *= -1
		}
		if _, ok := pairs[name]; !ok {
			pairs[name] = make(map[string]int)
		}
		pairs[name][other] = happiness
	}

	fmt.Println(bestHappiness(pairs))

	for _, others := range pairs {
		others["self"] = 0
	}
	pairs["self"] = make(map[string]int)
	for n := range pairs {
		if n != "self" {
			pairs["self"][n] = 0
		}
	}

	fmt.Println(bestHappiness(pairs))
}
