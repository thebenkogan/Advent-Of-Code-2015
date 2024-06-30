package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Advent-Of-Code-2015"
)

type status int

const (
	resting status = iota
	moving
)

type reindeer struct {
	name      string
	speed     int
	duration  int
	rest      int
	status    status
	remaining int
	position  int
	points    int
}

func main() {
	input := lib.GetInput()

	reindeers := make([]*reindeer, 0)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		nums := lib.ParseNums(line)
		name := split[0]
		speed := nums[0]
		duration := nums[1]
		rest := nums[2]
		reindeers = append(reindeers, &reindeer{name, speed, duration, rest, moving, duration, 0, 0})
	}

	for range 2503 {
		for _, r := range reindeers {
			if r.status == moving {
				r.remaining--
				r.position += r.speed
				if r.remaining == 0 {
					r.status = resting
					r.remaining = r.rest
				}
			} else {
				r.remaining--
				if r.remaining == 0 {
					r.status = moving
					r.remaining = r.duration
				}
			}
		}
		farthest := 0
		for _, r := range reindeers {
			farthest = max(farthest, r.position)
		}
		for _, r := range reindeers {
			if r.position == farthest {
				r.points++
			}
		}
	}

	farthest := 0
	for _, r := range reindeers {
		farthest = max(farthest, r.position)
	}

	fmt.Println(farthest)

	mostPoints := 0
	for _, r := range reindeers {
		mostPoints = max(mostPoints, r.points)
	}

	fmt.Println(mostPoints)
}
