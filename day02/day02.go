package day02

import (
	"strings"
)

func Part1(input string) int {
	score := 0
	for _, line := range strings.Split(input, "\n") {
		opponent := line[0]
		you := line[2]
		outcome := 0
		if opponent == 'A' { // rock
			if you == 'Y' { // paper
				outcome = 6
			} else if you == 'X' { // rock
				outcome = 3
			}
		} else if opponent == 'B' { // paper
			if you == 'Z' { // scissor
				outcome = 6
			} else if you == 'Y' { // paper
				outcome = 3
			}
		} else if opponent == 'C' { // scissor
			if you == 'X' { // rock
				outcome = 6
			} else if you == 'Z' { // scissor
				outcome = 3
			}
		} else {
			panic("unexpected")
		}

		score += outcome
		if you == 'X' {
			score++
		} else if you == 'Y' {
			score += 2
		} else if you == 'Z' {
			score += 3
		} else {
			panic("unexpected")
		}
	}
	return score
}

func Part2(input string) int {
	score := 0
	for _, line := range strings.Split(input, "\n") {
		opponent := line[0]
		desiredOutcome := line[2]
		var outcome int
		switch desiredOutcome {
		case 'X':
			outcome = 0
		case 'Y':
			outcome = 3
		case 'Z':
			outcome = 6
		default:
			panic("unexpected")
		}

		switch opponent {
		case 'A': // rock
			switch desiredOutcome {
			case 'X':
				outcome += 3 // scissor
			case 'Y':
				outcome += 1 // rock
			case 'Z':
				outcome += 2 // paper
			}
		case 'B': // paper
			switch desiredOutcome {
			case 'X':
				outcome += 1 // rock
			case 'Y':
				outcome += 2 // paper
			case 'Z':
				outcome += 3 // scissor
			}
		case 'C': // scissor
			switch desiredOutcome {
			case 'X':
				outcome += 2 // paper
			case 'Y':
				outcome += 3 // scissor
			case 'Z':
				outcome += 1 // rock
			}
		default:
			panic("unexpected")
		}

		score += outcome
	}
	return score
}
