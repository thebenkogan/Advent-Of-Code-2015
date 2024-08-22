package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

type instruction interface {
	execute(registers map[string]int) int
}

type hlf struct {
	r string
}

func (h hlf) execute(registers map[string]int) int {
	registers[h.r] /= 2
	return 1
}

type tpl struct {
	r string
}

func (h tpl) execute(registers map[string]int) int {
	registers[h.r] *= 3
	return 1
}

type inc struct {
	r string
}

func (h inc) execute(registers map[string]int) int {
	registers[h.r] += 1
	return 1
}

type jmp struct {
	o int
}

func (h jmp) execute(registers map[string]int) int {
	return h.o
}

type jie struct {
	o int
	r string
}

func (h jie) execute(registers map[string]int) int {
	if registers[h.r]%2 == 0 {
		return h.o
	}
	return 1
}

type jio struct {
	o int
	r string
}

func (h jio) execute(registers map[string]int) int {
	if registers[h.r] == 1 {
		return h.o
	}
	return 1
}

func parseInstruction(line string) instruction {
	split := strings.Split(line, " ")
	switch split[0] {
	case "hlf":
		return hlf{split[1]}
	case "tpl":
		return tpl{split[1]}
	case "inc":
		return inc{split[1]}
	case "jmp":
		nums := lib.ParseNums(split[1])
		return jmp{nums[0]}
	case "jie":
		nums := lib.ParseNums(split[2])
		return jie{nums[0], string(split[1][0])}
	case "jio":
		nums := lib.ParseNums(split[2])
		return jio{nums[0], string(split[1][0])}
	}
	panic("unknown instruction: " + split[0])
}

func run(ins []instruction, a int) int {
	registers := map[string]int{
		"a": a,
		"b": 0,
	}
	pc := 0

	for pc >= 0 && pc < len(ins) {
		instruction := ins[pc]
		pc += instruction.execute(registers)
	}

	return registers["b"]
}

func main() {
	input := lib.GetInput()

	instructions := make([]instruction, 0)
	for _, line := range strings.Split(input, "\n") {
		instructions = append(instructions, parseInstruction(line))
	}

	fmt.Println(run(instructions, 0))
	fmt.Println(run(instructions, 1))
}
