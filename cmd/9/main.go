package main

import (
	"fmt"
	"math"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
	"github.com/zyedidia/generic/heap"
)

type edge struct {
	dst  string
	dist int
}

type node struct {
	loc  string
	cost int
	seen map[string]bool
}

func getCheapestFromStart(adj map[string][]edge, start string) int {
	heap := heap.From(func(a, b node) bool { return a.cost < b.cost })
	heap.Push(node{start, 0, map[string]bool{start: true}})
	for heap.Size() > 0 {
		curr, _ := heap.Pop()
		if len(curr.seen) == len(adj) {
			return curr.cost
		}
		for _, e := range adj[curr.loc] {
			if !curr.seen[e.dst] {
				newSeen := make(map[string]bool)
				for k, v := range curr.seen {
					newSeen[k] = v
				}
				newSeen[e.dst] = true
				heap.Push(node{e.dst, curr.cost + e.dist, newSeen})
			}
		}
	}

	panic("no path found for start: " + start)
}

func getLongestFromStart(adj map[string][]edge, start string) int {
	stack := []node{{start, 0, map[string]bool{start: true}}}
	maxDist := 0
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(curr.seen) == len(adj) {
			maxDist = max(maxDist, curr.cost)
		}
		for _, e := range adj[curr.loc] {
			if !curr.seen[e.dst] {
				newSeen := make(map[string]bool)
				for k, v := range curr.seen {
					newSeen[k] = v
				}
				newSeen[e.dst] = true
				stack = append(stack, node{e.dst, curr.cost + e.dist, newSeen})
			}
		}
	}

	return maxDist
}

func main() {
	input := lib.GetInput()

	adj := make(map[string][]edge)
	for _, line := range strings.Split(input, "\n") {
		nums := lib.ParseNums(line)
		split := strings.Split(line, " ")
		src := split[0]
		dst := split[2]
		adj[src] = append(adj[src], edge{dst, nums[0]})
		adj[dst] = append(adj[dst], edge{src, nums[0]})
	}

	best := math.MaxInt
	longest := 0
	for start := range adj {
		best = min(best, getCheapestFromStart(adj, start))
		longest = max(longest, getLongestFromStart(adj, start))
	}

	fmt.Println(best)
	fmt.Println(longest)
}
