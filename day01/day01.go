package day01

import (
	"sort"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type elf struct {
	calories int
}

func Part1(input string) int {
	elves := common(input)
	return elves[0].calories
}

func Part2(input string) int {
	elves := common(input)
	return elves[0].calories + elves[1].calories + elves[2].calories
}

func common(input string) []elf {
	lines := strings.Split(input, "\n")

	elves := []elf{}
	e := elf{}
	for _, line := range lines {
		if line == "" {
			elves = append(elves, e)
			e = elf{}
		} else {
			e.calories += utils.MustAtoi(line)
		}
	}
	elves = append(elves, e)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})

	return elves
}
