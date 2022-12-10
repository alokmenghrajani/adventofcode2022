package day08

import (
	"github.com/alokmenghrajani/adventofcode2022/utils"
	"github.com/alokmenghrajani/adventofcode2022/utils/grid2d"
	"github.com/alokmenghrajani/adventofcode2022/utils/inputs"
)

func Part1(input string) int {
	grid := inputs.ToGrid2D(input, "\n", "", -1, utils.MustAtoi)
	visible := grid2d.NewGrid(grid.SizeX(), grid.SizeY(), false)

	for j := 0; j < grid.SizeY(); j++ {
		for i := 0; i < grid.SizeX(); i++ {
			directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
			for _, dir := range directions {
				v := true
				x := i
				y := j
				for {
					x += dir[0]
					y += dir[1]
					if x < 0 || x >= grid.SizeX() || y < 0 || y >= grid.SizeY() {
						break
					}
					if grid.Get(i, j) <= grid.Get(x, y) {
						v = false
						break
					}
				}
				if v {
					visible.Set(i, j, true)
					break
				}
			}
		}
	}

	sum := 0
	for j := 0; j < visible.SizeY(); j++ {
		for i := 0; i < visible.SizeX(); i++ {
			if visible.Get(i, j) {
				sum++
			}
		}
	}
	return sum
}

func Part2(input string) int {
	grid := inputs.ToGrid2D(input, "\n", "", -1, utils.MustAtoi)
	score := grid2d.NewGrid(grid.SizeX(), grid.SizeY(), 0)

	for j := 0; j < grid.SizeY(); j++ {
		for i := 0; i < grid.SizeX(); i++ {
			directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
			product := 1
			for _, dir := range directions {
				x := i
				y := j
				distance := 0
				for {
					x += dir[0]
					y += dir[1]
					if x < 0 || x >= grid.SizeX() || y < 0 || y >= grid.SizeY() {
						break
					}
					distance++
					if grid.Get(i, j) <= grid.Get(x, y) {
						break
					}
				}
				product *= distance
			}
			score.Set(i, j, product)
		}
	}

	best := 0
	for j := 0; j < score.SizeY(); j++ {
		for i := 0; i < score.SizeX(); i++ {
			best = utils.Max(best, score.Get(i, j))
		}
	}
	return best
}
