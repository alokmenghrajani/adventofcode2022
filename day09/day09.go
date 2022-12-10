package day09

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
	sparsegrid "github.com/alokmenghrajani/adventofcode2022/utils/sparseGrid"
)

type rope struct {
	size int
	posX []int
	posY []int
}

func Part1(input string) int {
	rope := newRope(2)
	grid := sparsegrid.NewGrid(false)
	grid.Set(rope.posX[rope.size-1], rope.posY[rope.size-1], true)
	for _, line := range strings.Split(input, "\n") {
		dx := 0
		dy := 0
		switch line[0] {
		case 'R':
			dx = 1
		case 'U':
			dy = -1
		case 'D':
			dy = 1
		case 'L':
			dx = -1
		}
		n := utils.MustAtoi(line[2:])
		for i := 0; i < n; i++ {
			// move head
			rope.posX[0] += dx
			rope.posY[0] += dy

			// update tail
			rope.updateTail()
			grid.Set(rope.posX[rope.size-1], rope.posY[rope.size-1], true)
		}
	}
	return grid.Visited()
}

func Part2(input string) int {
	rope := newRope(10)
	grid := sparsegrid.NewGrid(false)
	grid.Set(rope.posX[rope.size-1], rope.posY[rope.size-1], true)
	for _, line := range strings.Split(input, "\n") {
		dx := 0
		dy := 0
		switch line[0] {
		case 'R':
			dx = 1
		case 'U':
			dy = -1
		case 'D':
			dy = 1
		case 'L':
			dx = -1
		}
		n := utils.MustAtoi(line[2:])
		for i := 0; i < n; i++ {
			// move head
			rope.posX[0] += dx
			rope.posY[0] += dy

			// update tail
			rope.updateTail()
			grid.Set(rope.posX[rope.size-1], rope.posY[rope.size-1], true)
		}
	}
	return grid.Visited()
}

func newRope(size int) *rope {
	return &rope{
		size: size,
		posX: make([]int, size),
		posY: make([]int, size),
	}
}

func (r *rope) updateTail() {
outer:
	for i := 1; i < r.size; i++ {
		diffX := utils.Abs(r.posX[i-1] - r.posX[i])
		diffY := utils.Abs(r.posY[i-1] - r.posY[i])
		if diffX <= 1 && diffY <= 1 {
			// no need to update node if it's touching
			continue
		}
		moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, move := range moves {
			t := utils.Max(
				utils.Abs(r.posX[i-1]-r.posX[i]+move[0]),
				utils.Abs(r.posY[i-1]-r.posY[i]+move[1]))
			if t == 1 {
				r.posX[i] = r.posX[i-1] + move[0]
				r.posY[i] = r.posY[i-1] + move[1]
				continue outer
			}
		}
		moves = [][]int{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}}
		for _, move := range moves {
			t := utils.Max(
				utils.Abs(r.posX[i-1]-r.posX[i]+move[0]),
				utils.Abs(r.posY[i-1]-r.posY[i]+move[1]))
			if t == 1 {
				r.posX[i] = r.posX[i-1] + move[0]
				r.posY[i] = r.posY[i-1] + move[1]
				continue outer
			}
		}
		panic("unreachable")
	}
}
