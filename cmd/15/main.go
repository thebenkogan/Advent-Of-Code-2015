package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	input := lib.GetInput()

	ingredients := make([]ingredient, 0)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		name := split[0]
		nums := lib.ParseNums(line)
		capacity, durability, flavor, texture, calories := nums[0], nums[1], nums[2], nums[3], nums[4]
		ingredients = append(ingredients, ingredient{name, capacity, durability, flavor, texture, calories})
	}

	best := 0
	best500 := 0
	for a := range 100 {
		for b := range 100 {
			for c := range 100 {
				d := 100 - a - b - c
				if d < 0 {
					continue
				}
				capacity := a*ingredients[0].capacity + b*ingredients[1].capacity + c*ingredients[2].capacity + d*ingredients[3].capacity
				durability := a*ingredients[0].durability + b*ingredients[1].durability + c*ingredients[2].durability + d*ingredients[3].durability
				flavor := a*ingredients[0].flavor + b*ingredients[1].flavor + c*ingredients[2].flavor + d*ingredients[3].flavor
				texture := a*ingredients[0].texture + b*ingredients[1].texture + c*ingredients[2].texture + d*ingredients[3].texture
				calories := a*ingredients[0].calories + b*ingredients[1].calories + c*ingredients[2].calories + d*ingredients[3].calories
				total := max(capacity, 0) * max(durability, 0) * max(flavor, 0) * max(texture, 0)
				best = max(best, total)
				if calories == 500 {
					best500 = max(best500, total)
				}
			}
		}
	}

	fmt.Println(best)
	fmt.Println(best500)
}
