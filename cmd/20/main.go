package main

import (
	"fmt"
	"math"
	"strconv"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

func factorTotal(n int) int {
	total := 0
	seen := make(map[int]bool)
	for i := 1; i <= int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 && !seen[i] {
			total += i * 10
			seen[i] = true
			if !seen[n/i] {
				total += (n / i) * 10
				seen[n/i] = true
			}
		}
	}
	return total
}

// wtf?
// https://www.reddit.com/r/adventofcode/comments/3xjpp2/comment/cy59l7k/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
func factorTotalLimit(n int) int {
	total := 0
	seen := make(map[int]bool)
	for i := 1; i <= int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 && !seen[i] {
			seen[i] = true
			if n/i <= 50 {
				total += i * 11
			}
			other := n / i
			if !seen[other] {
				seen[other] = true
				if n/other <= 50 {
					total += (other) * 11
				}
			}
		}
	}
	return total
}

func main() {
	input := lib.GetInput()
	target, _ := strconv.Atoi(input)

	for i := 1; ; i++ {
		total := factorTotal(i)
		if total >= target {
			fmt.Println(i)
			break
		}
	}

	for i := 1; ; i++ {
		total := factorTotalLimit(i)
		if total >= target {
			fmt.Println(i)
			break
		}
	}
}
