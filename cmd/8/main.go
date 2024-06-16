package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func main() {
	input := lib.GetInput()
	split := strings.Split(input, "\n")

	numStr := 0
	numMem := 0
	for _, line := range split {
		numStr += len(line)
		line := line[1 : len(line)-1]
		i := 0
		for i < len(line) {
			if line[i] == '\\' {
				if line[i+1] == 'x' {
					i += 4
				} else {
					i += 2
				}
				numMem++
			} else {
				i++
				numMem++
			}
		}
	}

	fmt.Println(numStr - numMem)

	numExpanded := 0
	for _, line := range split {
		numExpanded += 2
		for _, c := range line {
			if c == '\\' || c == '"' {
				numExpanded += 2
			} else {
				numExpanded++
			}
		}
	}

	fmt.Println(numExpanded - numStr)
}
