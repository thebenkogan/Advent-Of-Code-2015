package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

var target = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

var CompoundRegex = regexp.MustCompile(`(\w+): (\d+)`)

func main() {
	input := lib.GetInput()
	split := strings.Split(input, "\n")

	for _, line := range split {
		nums := lib.ParseNums(line)
		sue := nums[0]
		good := true
		for _, compound := range CompoundRegex.FindAllStringSubmatch(line, -1) {
			name := compound[1]
			count, _ := strconv.Atoi(compound[2])
			if target[name] != count {
				good = false
				break
			}
		}
		if good {
			fmt.Println(sue)
		}
	}

	for _, line := range split {
		nums := lib.ParseNums(line)
		sue := nums[0]
		good := true
		for _, compound := range CompoundRegex.FindAllStringSubmatch(line, -1) {
			name := compound[1]
			count, _ := strconv.Atoi(compound[2])
			switch name {
			case "cats":
				fallthrough
			case "trees":
				if target[name] >= count {
					good = false
					break
				}
			case "pomeranians":
				fallthrough
			case "goldfish":
				if target[name] <= count {
					good = false
					break
				}
			default:
				if target[name] != count {
					good = false
					break
				}
			}
		}
		if good {
			fmt.Println(sue)
		}
	}
}
