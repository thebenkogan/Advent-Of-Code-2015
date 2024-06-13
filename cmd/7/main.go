package main

import (
	"fmt"
	"maps"
	"strconv"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

type node struct {
	op     string
	wires  []string
	output string
}

func (n *node) evaluate(signals map[string]uint16) {
	switch n.op {
	case "AND":
		v, err := strconv.Atoi(n.wires[0])
		val := uint16(v)
		if err != nil {
			val = signals[n.wires[0]]
		}
		signals[n.output] = val & signals[n.wires[1]]
	case "OR":
		signals[n.output] = signals[n.wires[0]] | signals[n.wires[1]]
	case "LSHIFT":
		val, _ := strconv.Atoi(n.wires[1])
		signals[n.output] = signals[n.wires[0]] << val
	case "RSHIFT":
		val, _ := strconv.Atoi(n.wires[1])
		signals[n.output] = signals[n.wires[0]] >> val
	case "NOT":
		signals[n.output] = ^signals[n.wires[0]]
	case "ASSIGN":
		signals[n.output] = signals[n.wires[0]]
	}
}

func (n *node) isReady(signals map[string]uint16) bool {
	for _, wire := range n.wires {
		_, err := strconv.Atoi(wire)
		if err == nil {
			continue
		}
		if _, ok := signals[wire]; !ok {
			return false
		}
	}
	return true
}

func evaluateNodes(nodes []*node, signals map[string]uint16) {
	seen := make(map[*node]bool)
	for len(seen) < len(nodes) {
		for _, n := range nodes {
			if n.isReady(signals) && !seen[n] {
				n.evaluate(signals)
				seen[n] = true
			}
		}
	}
}

func main() {
	input := lib.GetInput()

	signals := make(map[string]uint16)
	nodes := make([]*node, 0)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " -> ")
		src, dst := split[0], split[1]

		n, err := strconv.Atoi(src)
		if err == nil {
			signals[dst] = uint16(n)
			continue
		}

		node := node{output: dst}
		switch {
		case strings.Contains(src, "AND"):
			node.op = "AND"
			node.wires = strings.Split(src, " AND ")
		case strings.Contains(src, "OR"):
			node.op = "OR"
			node.wires = strings.Split(src, " OR ")
		case strings.Contains(src, "LSHIFT"):
			node.op = "LSHIFT"
			node.wires = strings.Split(src, " LSHIFT ")
		case strings.Contains(src, "RSHIFT"):
			node.op = "RSHIFT"
			node.wires = strings.Split(src, " RSHIFT ")
		case strings.Contains(src, "NOT"):
			node.op = "NOT"
			node.wires = []string{strings.Split(src, "NOT ")[1]}
		default:
			node.op = "ASSIGN"
			node.wires = []string{src}
		}
		nodes = append(nodes, &node)
	}

	nextSignals := make(map[string]uint16)
	maps.Copy(nextSignals, signals)

	evaluateNodes(nodes, signals)
	fmt.Println(signals["a"])

	nextSignals["b"] = signals["a"]
	evaluateNodes(nodes, nextSignals)
	fmt.Println(nextSignals["a"])
}
