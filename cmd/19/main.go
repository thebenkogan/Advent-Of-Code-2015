package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func getNexts(replacements [][2]string, start string) map[string]struct{} {
	distinct := make(map[string]struct{})
	for _, r := range replacements {
		s, e := r[0], r[1]
		for i := 0; i < len(start)-len(s)+1; i++ {
			if start[i:i+len(s)] == s {
				replaced := start[:i] + e + start[i+len(s):]
				distinct[replaced] = struct{}{}
			}
		}
	}
	return distinct
}

func main() {
	input := lib.GetInput()

	sections := strings.Split(input, "\n\n")
	start := sections[1]
	rs := strings.Split(sections[0], "\n")
	replacements := make([][2]string, 0)
	for _, r := range rs {
		split := strings.Split(r, " => ")
		replacements = append(replacements, [2]string{split[0], split[1]})
	}

	nexts := getNexts(replacements, start)
	fmt.Println(len(nexts))

	symbols := 0
	rnOrAr := 0
	y := 0
	i := 0
	for i < len(start) {
		symbol := string(start[i])
		if i < len(start)-1 && start[i+1] >= 'a' && start[i+1] <= 'z' {
			symbol += string(start[i+1])
		}
		symbols += 1
		if symbol == "Rn" || symbol == "Ar" {
			rnOrAr++
		}
		if symbol == "Y" {
			y++
		}
		i += len(symbol)
	}

	// https://www.reddit.com/r/adventofcode/comments/3xflz8/comment/cy4etju/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
	fmt.Println(symbols - rnOrAr - (2 * y) - 1)
}
