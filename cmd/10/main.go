package main

import (
	"fmt"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func step(input string) string {
	out := strings.Builder{}
	i := 0
	for i < len(input) {
		total := 1
		for i+1 < len(input) && input[i] == input[i+1] {
			total++
			i++
		}
		out.WriteString(strconv.Itoa(total))
		out.WriteByte(input[i])
		i++
	}
	return out.String()
}

func main() {
	input := lib.GetInput()

	for range 40 {
		input = step(input)
	}

	fmt.Println(len(input))

	for range 10 {
		input = step(input)
	}

	fmt.Println(len(input))
}
