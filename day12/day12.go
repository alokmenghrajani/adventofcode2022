package day12

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
	"github.com/alokmenghrajani/adventofcode2022/utils/grid2d"
)

type cell struct {
	height   int
	distance int
}

func Part1(input string) int {
	// parse input
	lines := strings.Split(input, "\n")
	grid := grid2d.NewGrid[*cell](len(lines[0]), len(lines), nil)
	startX := -1
	startY := -1

	for j, line := range lines {
		for i, c := range line {
			newCell := &cell{height: int(c - 'a'), distance: utils.MaxInt}
			if c == 'S' {
				newCell.height = 0
				startX = i
				startY = j
			} else if c == 'E' {
				newCell.height = 25
				newCell.distance = 0
			}
			grid.Set(i, j, newCell)
		}
	}

	fillGrid(grid)

	c := grid.Get(startX, startY)
	return c.distance
}

func Part2(input string) int {
	// parse input
	lines := strings.Split(input, "\n")
	grid := grid2d.NewGrid[*cell](len(lines[0]), len(lines), nil)

	for j, line := range lines {
		for i, c := range line {
			newCell := &cell{height: int(c - 'a'), distance: utils.MaxInt}
			if c == 'S' {
				newCell.height = 0
			} else if c == 'E' {
				newCell.height = 25
				newCell.distance = 0
			}
			grid.Set(i, j, newCell)
		}
	}

	fillGrid(grid)
	min := utils.MaxInt
	for j := 0; j < grid.SizeY(); j++ {
		for i := 0; i < grid.SizeX(); i++ {
			c1 := grid.Get(i, j)
			if c1.height == 0 {
				min = utils.Min(min, c1.distance)
			}
		}
	}
	return min
}

func fillGrid(grid *grid2d.Grid[*cell]) {
	done := false
	for !done {
		done = true
		for j := 0; j < grid.SizeY(); j++ {
			for i := 0; i < grid.SizeX(); i++ {
				c1 := grid.Get(i, j)
				if c1.distance == utils.MaxInt {
					continue
				}
				moves := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
				for _, move := range moves {
					x2 := i + move[0]
					y2 := j + move[1]
					c2 := grid.Get(x2, y2)
					if c2 == nil {
						continue
					}
					if c2.height < c1.height-1 {
						continue
					}
					if c2.distance > c1.distance+1 {
						c2.distance = c1.distance + 1
						done = false
					}
				}
			}
		}
	}
}

func printGrid(grid *grid2d.Grid[*cell]) {
	r := grid.StringWithFormatter(func(c *cell, i, j int) string {
		if c.distance == utils.MaxInt {
			return "?? "
		}
		return fmt.Sprintf("%02d ", c.distance)
	})
	fmt.Println(r)
}
