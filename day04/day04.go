package day04

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

func Part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		ranges := strings.SplitN(line, ",", 2)
		range1 := strings.SplitN(ranges[0], "-", 2)
		x1 := utils.MustAtoi(range1[0])
		x2 := utils.MustAtoi(range1[1])

		range2 := strings.SplitN(ranges[1], "-", 2)
		y1 := utils.MustAtoi(range2[0])
		y2 := utils.MustAtoi(range2[1])

		if x1 <= y1 && x2 >= y2 {
			sum++
		} else if y1 <= x1 && y2 >= x2 {
			sum++
		}
	}
	return sum
}

func Part2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		ranges := strings.SplitN(line, ",", 2)
		range1 := strings.SplitN(ranges[0], "-", 2)
		x1 := utils.MustAtoi(range1[0])
		x2 := utils.MustAtoi(range1[1])

		range2 := strings.SplitN(ranges[1], "-", 2)
		y1 := utils.MustAtoi(range2[0])
		y2 := utils.MustAtoi(range2[1])

		if x1 <= y1 && y1 <= x2 {
			sum++
		} else if x1 <= y2 && y2 <= x2 {
			sum++
		} else if y1 <= x1 && x1 <= y2 {
			sum++
		}
	}
	return sum
}
