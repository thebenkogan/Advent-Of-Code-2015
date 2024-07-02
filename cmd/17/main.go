package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func backtrack(sizes []int, index int, total int) int {
	if total == 150 {
		return 1
	} else if total > 150 {
		return 0
	}
	num := 0
	for i := index; i < len(sizes); i++ {
		num += backtrack(sizes, i+1, total+sizes[i])
	}
	return num
}

func backtrackWithTarget(sizes []int, index int, total int, numContainers int, targetContainers int) int {
	if total == 150 && numContainers == targetContainers {
		return 1
	} else if total >= 150 {
		return 0
	}
	num := 0
	for i := index; i < len(sizes); i++ {
		num += backtrackWithTarget(sizes, i+1, total+sizes[i], numContainers+1, targetContainers)
	}
	return num
}

func minContainers(sizes []int) int {
	// dp[i] = min number of containers needed to hold size i
	// dp[0] = 0
	// dp[i] = min(dp[i-j]) + 1 for all j in sizes
	dp := make([]int, 151)
	for i := 1; i <= 150; i++ {
		best := math.MaxInt
		for _, s := range sizes {
			if i-s >= 0 {
				best = min(best, dp[i-s]+1)
			}
		}
		dp[i] = best
	}
	return dp[150]
}

func main() {
	input := lib.GetInput()

	sizes := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		n, _ := strconv.Atoi(line)
		sizes = append(sizes, n)
	}

	fmt.Println(backtrack(sizes, 0, 0))

	targetContainers := minContainers(sizes)
	fmt.Println(backtrackWithTarget(sizes, 0, 0, 0, targetContainers))
}
