package main

import (
	"fmt"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func main() {
	input := lib.GetInput()

	pos := lib.Vector{X: 0, Y: 0}
	seen := make(map[string]bool, len(input))
	seen[pos.Hash()] = true
	for _, move := range input {
		switch move {
		case '^':
			pos.Y++
		case 'v':
			pos.Y--
		case '<':
			pos.X--
		case '>':
			pos.X++
		}
		seen[pos.Hash()] = true
	}

	fmt.Println(len(seen))

	santa := lib.Vector{X: 0, Y: 0}
	robo := lib.Vector{X: 0, Y: 0}
	santaTurn := true
	seen = make(map[string]bool, len(input))
	seen[robo.Hash()] = true
	for _, move := range input {
		update := &santa
		if !santaTurn {
			update = &robo
		}
		switch move {
		case '^':
			update.Y++
		case 'v':
			update.Y--
		case '<':
			update.X--
		case '>':
			update.X++
		}
		seen[update.Hash()] = true
		santaTurn = !santaTurn
	}

	fmt.Println(len(seen))
}
