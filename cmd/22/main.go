package main

import (
	"fmt"
	"maps"
	"math"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
	"github.com/zyedidia/generic/heap"
)

type player struct {
	hp   int
	mana int
	dmg  int
}

type effect struct {
	cost     int
	dmg      int
	armor    int
	duration int
	mana     int
	heal     int
}

var effects = map[string]effect{
	"magic missile": {53, 4, 0, 0, 0, 0},
	"drain":         {73, 2, 0, 0, 0, 2},
	"shield":        {113, 0, 7, 6, 0, 0},
	"poison":        {173, 3, 0, 6, 0, 0},
	"recharge":      {229, 0, 0, 5, 101, 0},
}

type state struct {
	spent         int
	activeEffects map[string]int // effect name -> time left
	p             player
	boss          player
}

func (s *state) copy() state {
	return state{s.spent, maps.Clone(s.activeEffects), s.p, s.boss}
}

func (s *state) applyEffects() int {
	armor := 0
	for e, t := range s.activeEffects {
		effect := effects[e]
		s.boss.hp -= effect.dmg
		armor += effect.armor
		s.p.mana += effect.mana
		s.p.hp += effect.heal
		if t <= 1 {
			delete(s.activeEffects, e)
		} else {
			s.activeEffects[e] = t - 1
		}
	}
	return armor
}

func run(bossHp, bossDmg int, hardMode bool) int {
	boss := player{bossHp, 0, bossDmg}
	p := player{50, 500, 0}

	heap := heap.From(func(a, b state) bool { return a.spent < b.spent })
	heap.Push(state{0, map[string]int{}, p, boss})
	numFound := 0 // check the first 5 since the last effect price could make the difference
	minCost := math.MaxInt
	for heap.Size() > 0 {
		curr, _ := heap.Pop()

		if hardMode {
			curr.p.hp -= 1
			if curr.p.hp <= 0 {
				continue
			}
		}

		// cast current effects before player turn
		curr.applyEffects()

		// player turn, try each effect
		for name, e := range effects {
			if curr.p.mana < e.cost {
				continue // too expensive
			}
			if _, ok := curr.activeEffects[name]; ok {
				continue // already active
			}

			next := curr.copy()
			next.activeEffects[name] = e.duration
			next.spent += e.cost
			next.p.mana -= e.cost

			// boss turn, first apply effects
			armor := next.applyEffects()
			if next.boss.hp <= 0 {
				numFound++
				minCost = min(minCost, next.spent)
				if numFound == 5 {
					return minCost
				}
			}
			next.p.hp -= max(next.boss.dmg-armor, 1)
			if next.p.hp <= 0 {
				continue
			}

			heap.Push(next)
		}
	}
	panic("uh oh")
}

func main() {
	input := lib.GetInput()
	nums := lib.ParseNums(input)

	fmt.Println(run(nums[0], nums[1], false))
	fmt.Println(run(nums[0], nums[1], true))
}
