package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func step(grid [][]rune, cornersOn bool) [][]rune {
	out := make([][]rune, len(grid))
	for i, row := range grid {
		out[i] = make([]rune, len(row))
		copy(out[i], row)
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			neighborsOn := 0
			for _, dir := range lib.ALL_DIRS {
				nx, ny := x+dir.X, y+dir.Y
				if nx >= 0 && nx < len(grid[0]) && ny >= 0 && ny < len(grid) && grid[ny][nx] == '#' {
					neighborsOn++
				}
			}
			if grid[y][x] == '#' && !(neighborsOn == 2 || neighborsOn == 3) {
				out[y][x] = '.'
			}
			if grid[y][x] == '.' && neighborsOn == 3 {
				out[y][x] = '#'
			}
		}
	}
	if cornersOn {
		turnOnCorners(out)
	}
	return out
}

func turnOnCorners(grid [][]rune) {
	grid[0][0] = '#'
	grid[0][len(grid[0])-1] = '#'
	grid[len(grid)-1][0] = '#'
	grid[len(grid)-1][len(grid[0])-1] = '#'
}

func run100(input string, cornersOn bool) int {
	grid := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]rune, 0)
		for _, c := range line {
			row = append(row, c)
		}
		grid = append(grid, row)
	}
	if cornersOn {
		turnOnCorners(grid)
	}

	for range 100 {
		grid = step(grid, cornersOn)
	}

	on := 0
	for _, row := range grid {
		for _, c := range row {
			if c == '#' {
				on++
			}
		}
	}
	return on
}

func main() {
	input := lib.GetInput()
	fmt.Println(run100(input, false))
	fmt.Println(run100(input, true))
}
