package day03

import (
	"strings"
)

func Part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		firstHalf := line[0 : len(line)/2]
		secondHalf := line[len(line)/2:]
		m := map[rune]bool{}
		for _, c := range firstHalf {
			m[c] = true
		}
		common := map[rune]bool{}
		for _, c := range secondHalf {
			if _, ok := m[c]; ok {
				common[c] = true
			}
		}
		for c := range common {
			sum += value(byte(c))
		}
	}
	return sum
}

func Part2(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 3 {
		m := map[rune]int{}
		for j := 0; j < 3; j++ {
			m2 := map[rune]bool{}
			line := lines[i+j]
			for _, c := range line {
				m2[c] = true
			}
			for c := range m2 {
				m[c]++
			}
		}
		for k, v := range m {
			if v != 3 {
				continue
			}
			sum += value(byte(k))
		}
	}
	return sum
}

func value(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 27)
	}
	panic("unexpected")
}
