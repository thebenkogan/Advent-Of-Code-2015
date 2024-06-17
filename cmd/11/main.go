package main

import (
	"fmt"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func increment(s string) string {
	buf := []byte(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'z' {
			buf[i] = 'a'
		} else {
			buf[i]++
			break
		}
	}
	return string(buf)
}

func hasStraight(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i]+2 == s[i+2] {
			return true
		}
	}
	return false
}

func notHasIOL(s string) bool {
	for _, r := range s {
		if r == 'i' || r == 'o' || r == 'l' {
			return false
		}
	}
	return true
}

func hasTwoPairs(s string) bool {
	pairs := 0
	seenPair := ""
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] && seenPair != string(s[i:i+2]) {
			seenPair = string(s[i : i+2])
			pairs++
			i++
		}
	}
	return pairs >= 2
}

func getNextPassword(s string) string {
	for {
		s = increment(s)
		if hasStraight(s) && notHasIOL(s) && hasTwoPairs(s) {
			return s
		}
	}
}

func main() {
	input := lib.GetInput()

	password1 := getNextPassword(input)
	password2 := getNextPassword(password1)
	fmt.Println(password1)
	fmt.Println(password2)
}
