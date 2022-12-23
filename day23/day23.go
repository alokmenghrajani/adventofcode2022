package day23

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type point struct {
	x, y int
}

var north = point{0, -1}
var northeast = point{1, -1}
var northwest = point{-1, -1}

var south = point{0, 1}
var southeast = point{1, 1}
var southwest = point{-1, 1}

var west = point{-1, 0}
var east = point{1, 0}

var moves = [][]point{
	{north, northeast, northwest},
	{south, southeast, southwest},
	{west, northwest, southwest},
	{east, northeast, southeast},
}

func (p point) add(other point) point {
	return point{p.x + other.x, p.y + other.y}
}

func Part1(input string) int {
	elves := []*point{}
	for j, line := range strings.Split(input, "\n") {
		for i, c := range line {
			if c == '#' {
				elves = append(elves, &point{x: i, y: j})
			}
		}
	}

	for round := 0; round < 10; round++ {
		moveAll(elves, round)
	}
	minX := utils.MaxInt
	maxX := utils.MinInt
	minY := utils.MaxInt
	maxY := utils.MinInt
	for _, elf := range elves {
		minX = utils.Min(minX, elf.x)
		maxX = utils.Max(maxX, elf.x)
		minY = utils.Min(minY, elf.y)
		maxY = utils.Max(maxY, elf.y)
	}

	return (maxX-minX+1)*(maxY-minY+1) - len(elves)
}

func Part2(input string) int {
	elves := []*point{}
	for j, line := range strings.Split(input, "\n") {
		for i, c := range line {
			if c == '#' {
				elves = append(elves, &point{x: i, y: j})
			}
		}
	}

	round := 0
	for {
		if moveAll(elves, round) == 0 {
			return round + 1
		}
		round++
	}
}

func moveAll(elves []*point, round int) int {
	proposalsCount := map[point]int{}
	proposals := make([]point, len(elves))
	didMove := 0
	for i, elf := range elves {
		p := move(elves, elf, round)
		proposalsCount[p]++
		proposals[i] = p
	}
	for i, elf := range elves {
		p := proposals[i]
		if proposalsCount[p] == 1 {
			if elf.x != p.x || elf.y != p.y {
				didMove++
			}
			*elf = p
		}
	}
	return didMove
}

func move(elves []*point, elf *point, round int) point {
	// count how many elves are neighboring
	neighbors := 0
	for _, m := range []point{northwest, north, northeast, west, east, southwest, south, southeast} {
		p := elf.add(m)
		if !available(elves, p) {
			neighbors++
		}
	}
	if neighbors == 0 {
		return *elf
	}

	// figure out where to move
outer:
	for i := round; i < round+4; i++ {
		m := moves[i%4]
		for j := 0; j < 3; j++ {
			p := elf.add(m[j])
			if !available(elves, p) {
				continue outer
			}
		}
		return elf.add(m[0])
	}
	return *elf
}

func available(elves []*point, p point) bool {
	for _, elf := range elves {
		if elf.x == p.x && elf.y == p.y {
			return false
		}
	}
	return true
}

func debug(elves []*point) {
	minX := utils.MaxInt
	maxX := utils.MinInt
	minY := utils.MaxInt
	maxY := utils.MinInt
	for _, elf := range elves {
		minX = utils.Min(minX, elf.x)
		maxX = utils.Max(maxX, elf.x)
		minY = utils.Min(minY, elf.y)
		maxY = utils.Max(maxY, elf.y)
	}
	for j := minY; j <= maxY; j++ {
	outer:
		for i := minX; i <= maxX; i++ {
			for _, elf := range elves {
				if elf.x == i && elf.y == j {
					fmt.Printf("#")
					continue outer
				}
			}
			fmt.Printf(".")
		}
		fmt.Println()
	}
	fmt.Println()
}
