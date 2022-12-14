package day14

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
	sparsegrid "github.com/alokmenghrajani/adventofcode2022/utils/sparseGrid"
)

type cell string

const (
	rock cell = "#"
	air  cell = "."
	sand cell = "o"
)

func Part1(input string) int {
	grid := sparsegrid.NewGrid(air)
	for _, line := range strings.Split(input, "\n") {
		drawLines(grid, strings.Split(line, " -> "))
	}
	i := 0
	for !dropSand(grid, 500, 0) {
		i++
	}
	return i
}

func Part2(input string) int {
	grid := sparsegrid.NewGrid(air)
	for _, line := range strings.Split(input, "\n") {
		drawLines(grid, strings.Split(line, " -> "))
	}
	_, maxY := grid.SizeY()
	i := 0
	for !dropSandPart2(grid, 500, 0, maxY) {
		i++
	}
	return i + 1 // ugh
}

func drawLines(grid *sparsegrid.SparseGrid[cell], points []string) {
	xy := [][]int{}
	for _, point := range points {
		pair := strings.Split(point, ",")
		x := utils.MustAtoi(pair[0])
		y := utils.MustAtoi(pair[1])
		xy = append(xy, []int{x, y})
	}

	for i := 1; i < len(xy); i++ {
		drawLine(grid, xy[i-1], xy[i])
	}
}

func drawLine(grid *sparsegrid.SparseGrid[cell], from []int, to []int) {
	x := from[0]
	y := from[1]
	dx := utils.Sign(to[0] - x)
	dy := utils.Sign(to[1] - y)
	for {
		grid.Set(x, y, rock)
		if x == to[0] && y == to[1] {
			break
		}
		x += dx
		y += dy
	}
}

func dropSand(grid *sparsegrid.SparseGrid[cell], sandX, sandY int) bool {
	_, maxY := grid.SizeY()
	for {
		if sandY > maxY {
			return true
		}
		c := grid.Get(sandX, sandY+1)
		if c == air {
			sandY++
			continue
		}
		c = grid.Get(sandX-1, sandY+1)
		if c == air {
			sandX--
			sandY++
			continue
		}
		c = grid.Get(sandX+1, sandY+1)
		if c == air {
			sandX++
			sandY++
			continue
		}
		grid.Set(sandX, sandY, sand)
		return false
	}
}

func dropSandPart2(grid *sparsegrid.SparseGrid[cell], sandX, sandY, maxY int) bool {
	for {
		var c cell
		if sandY+1 == maxY+2 {
			c = rock
		} else {
			c = grid.Get(sandX, sandY+1)
		}
		if c == air {
			sandY++
			continue
		}
		if sandY+1 == maxY+2 {
			c = rock
		} else {
			c = grid.Get(sandX-1, sandY+1)
		}
		if c == air {
			sandX--
			sandY++
			continue
		}
		if sandY+1 == maxY+2 {
			c = rock
		} else {
			c = grid.Get(sandX+1, sandY+1)
		}
		if c == air {
			sandX++
			sandY++
			continue
		}
		grid.Set(sandX, sandY, sand)
		if sandX == 500 && sandY == 0 {
			return true
		}
		return false
	}
}
