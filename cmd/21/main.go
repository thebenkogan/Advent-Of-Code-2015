package main

import (
	"fmt"
	"math"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

type player struct {
	hp    int
	dmg   int
	armor int
}

type item struct {
	cost  int
	dmg   int
	armor int
}

var weapons = []item{
	{8, 4, 0},
	{10, 5, 0},
	{25, 6, 0},
	{40, 7, 0},
	{74, 8, 0},
}

var armors = []*item{
	nil,
	{13, 0, 1},
	{31, 0, 2},
	{53, 0, 3},
	{75, 0, 4},
	{102, 0, 5},
}

var rings = []*item{
	nil,
	{25, 1, 0},
	{50, 2, 0},
	{100, 3, 0},
	{20, 0, 1},
	{40, 0, 2},
	{80, 0, 3},
}

func play(pl, boss player) bool {
	plTurn := true
	for pl.hp > 0 && boss.hp > 0 {
		if plTurn {
			boss.hp -= max(pl.dmg-boss.armor, 1)
		} else {
			pl.hp -= max(boss.dmg-pl.armor, 1)
		}
		plTurn = !plTurn
	}
	return pl.hp > 0
}

func main() {
	input := lib.GetInput()
	nums := lib.ParseNums(input)
	boss := player{nums[0], nums[1], nums[2]}

	leastToWin := math.MaxInt
	mostToLose := 0
	for _, weapon := range weapons {
		for _, armor := range armors {
			// 0-1 ring, good enough for part 1
			for _, ring := range rings {
				dmg := weapon.dmg
				armorPoints := 0
				fullCost := weapon.cost
				if armor != nil {
					armorPoints = armor.armor
					fullCost += armor.cost
				}
				if ring != nil {
					dmg += ring.dmg
					armorPoints += ring.armor
					fullCost += ring.cost
				}
				if play(player{100, dmg, armorPoints}, boss) {
					leastToWin = min(leastToWin, fullCost)
				} else {
					mostToLose = max(mostToLose, fullCost)
				}
			}

			// 2 rings
			for i := 1; i < len(rings); i++ {
				for j := i + 1; j < len(rings); j++ {
					ring1 := rings[i]
					ring2 := rings[j]
					dmg := weapon.dmg + ring1.dmg + ring2.dmg
					armorPoints := ring1.armor + ring2.armor
					fullCost := weapon.cost + ring1.cost + ring2.cost
					if armor != nil {
						armorPoints += armor.armor
						fullCost += armor.cost
					}
					if play(player{100, dmg, armorPoints}, boss) {
						leastToWin = min(leastToWin, fullCost)
					} else {
						mostToLose = max(mostToLose, fullCost)
					}
				}
			}
		}
	}

	fmt.Println(leastToWin)
	fmt.Println(mostToLose)
}
