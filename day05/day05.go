package day05

import (
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type Move struct {
	Amount int
	From   int
	To     int
}

func Part1(input string) string {
	stacks := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		stacks[i] = []byte{}
	}

	pieces := strings.Split(input, "\n\n")
	piece := strings.Split(pieces[0], "\n")
	for i := len(piece) - 2; i >= 0; i-- {
		line := piece[i]
		for j := 0; j*4+1 < len(line); j++ {
			if line[j*4+1] != ' ' {
				stacks[j] = append(stacks[j], line[j*4+1])
			}
		}
	}

	for _, line := range strings.Split(pieces[1], "\n") {
		m := Move{}
		utils.MustParseToStruct(regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`), line, &m)
		m.From--
		m.To--
		for i := 0; i < m.Amount; i++ {
			n := len(stacks[m.From])
			top := stacks[m.From][n-1]
			stacks[m.From] = stacks[m.From][0 : n-1]
			stacks[m.To] = append(stacks[m.To], top)
		}
	}

	var r strings.Builder
	for i := 0; i < len(stacks); i++ {
		n := len(stacks[i])
		if n > 0 {
			top := stacks[i][n-1]
			r.WriteByte(top)
		}
	}
	return r.String()
}

func Part2(input string) string {
	stacks := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		stacks[i] = []byte{}
	}

	pieces := strings.Split(input, "\n\n")
	piece := strings.Split(pieces[0], "\n")
	for i := len(piece) - 2; i >= 0; i-- {
		line := piece[i]
		for j := 0; j*4+1 < len(line); j++ {
			if line[j*4+1] != ' ' {
				stacks[j] = append(stacks[j], line[j*4+1])
			}
		}
	}

	for _, line := range strings.Split(pieces[1], "\n") {
		m := Move{}
		utils.MustParseToStruct(regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`), line, &m)
		m.From--
		m.To--

		n := len(stacks[m.From])
		top := stacks[m.From][n-m.Amount : n]
		stacks[m.From] = stacks[m.From][0 : n-m.Amount]
		stacks[m.To] = append(stacks[m.To], top...)
	}

	var r strings.Builder
	for i := 0; i < len(stacks); i++ {
		n := len(stacks[i])
		if n > 0 {
			top := stacks[i][n-1]
			r.WriteByte(top)
		}
	}
	return r.String()
}
