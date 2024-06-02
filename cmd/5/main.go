package main

import (
	"fmt"
	"slices"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func main() {
	input := lib.GetInput()
	split := strings.Split(input, "\n")

	total := 0
	for _, line := range split {
		if hasThreeVowels(line) && hasDoubleLetter(line) && hasNoBadStrings(line) {
			total++
		}
	}

	fmt.Println(total)

	total = 0
	for _, line := range split {
		if hasRepeatingPairs(line) && hasInBetween(line) {
			total++
		}
	}

	fmt.Println(total)
}

var VOWELS = []rune{'a', 'e', 'i', 'o', 'u'}

func hasThreeVowels(s string) bool {
	total := 0
	for _, r := range s {
		if slices.Contains(VOWELS, r) {
			total++
		}
	}
	return total >= 3
}

func hasDoubleLetter(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

var BAD_STRINGS = []string{"ab", "cd", "pq", "xy"}

func hasNoBadStrings(s string) bool {
	for _, bad := range BAD_STRINGS {
		if strings.Contains(s, bad) {
			return false
		}
	}
	return true
}

func hasRepeatingPairs(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if strings.Contains(s[i+2:], pair) {
			return true
		}
	}
	return false
}

func hasInBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
