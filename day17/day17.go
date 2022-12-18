package day17

import (
	sparsegrid "github.com/alokmenghrajani/adventofcode2022/utils/sparseGrid"
)

type shape [][]bool

type state struct {
	top          int
	grid         *sparsegrid.SparseGrid[bool]
	currentShape int
	currentGas   int
	posX         int
	posY         int
	shapes       []shape
	gasSequence  string
	totalShapes  int
}

func Part1(input string) int {
	s := droprocks(input, 2022)
	return -s.top
}

func Part2(input string) int {
	// assumption: there's a cyclic pattern
	// I can't explain why I need the 100 ¯\_(ツ)_/¯
	nonCycle := 100
	s := droprocks(input, nonCycle)

	// Find cycle
	previouslySeen := map[int]int{}
	for {
		for i := 0; i < len(s.shapes); i++ {
			s.dropNext()
		}
		if _, ok := previouslySeen[s.currentGas]; ok {
			break
		}
		previouslySeen[s.currentGas] = s.totalShapes
	}
	nonCycle = previouslySeen[s.currentGas]
	cycle := s.totalShapes - nonCycle

	// we now know the nonCycle and cycle values, start over and track -s.top
	s = droprocks(input, nonCycle)
	n1 := -s.top

	for i := 0; i < cycle; i++ {
		s.dropNext()
	}
	n2 := -s.top

	leftOver := (1000000000000 - nonCycle) % cycle
	for i := 0; i < leftOver; i++ {
		s.dropNext()
	}
	n3 := -s.top

	return n1 + (1000000000000-nonCycle-leftOver)/cycle*(n2-n1) + (n3 - n2)
}

func droprocks(input string, rocks int) state {
	s := state{
		top:          0,
		grid:         sparsegrid.NewGrid(false),
		currentShape: 0,
		currentGas:   0,
		shapes:       []shape{},
		gasSequence:  input,
	}

	s.shapes = append(s.shapes, shape{{true, true, true, true}})
	s.shapes = append(s.shapes, shape{{false, true, false}, {true, true, true}, {false, true, false}})
	s.shapes = append(s.shapes, shape{{false, false, true}, {false, false, true}, {true, true, true}})
	s.shapes = append(s.shapes, shape{{true}, {true}, {true}, {true}})
	s.shapes = append(s.shapes, shape{{true, true}, {true, true}})

	for i := 0; i < 7; i++ {
		s.grid.Set(i, 0, true)
	}
	for i := 0; i < rocks; i++ {
		s.dropNext()
	}

	return s
}

func (s *state) dropNext() {
	shape := s.shapes[s.currentShape]
	s.posX = 2
	s.posY = s.top - 3 - len(shape)
	s.drop()
	s.currentShape = (s.currentShape + 1) % len(s.shapes)
	s.totalShapes++
}

func (s *state) drop() {
	shape := s.shapes[s.currentShape]
	for {
		gas := s.gasSequence[s.currentGas]
		s.currentGas = (s.currentGas + 1) % len(s.gasSequence)
		dx := 0
		switch gas {
		case '<':
			dx = -1
		case '>':
			dx = 1
		default:
			panic("bad input")
		}
		if s.fits(shape, s.posX+dx, s.posY) {
			s.posX = s.posX + dx
		}

		if s.fits(shape, s.posX, s.posY+1) {
			s.posY = s.posY + 1
		} else {
			s.place(shape, s.posX, s.posY)
			return
		}
	}
}

func (s *state) fits(shape shape, posX, posY int) bool {
	for j := 0; j < len(shape); j++ {
		for i := 0; i < len(shape[j]); i++ {
			if shape[j][i] {
				if posX+i >= 7 {
					return false
				}
				if posX+i < 0 {
					return false
				}
				if s.grid.Get(posX+i, posY+j) {
					return false
				}
			}
		}
	}
	return true
}

func (s *state) place(shape shape, posX, posY int) {
	for j := 0; j < len(shape); j++ {
		for i := 0; i < len(shape[j]); i++ {
			if shape[j][i] {
				s.grid.Set(posX+i, posY+j, true)
			}
		}
	}
	s.top, _ = s.grid.SizeY()
}
