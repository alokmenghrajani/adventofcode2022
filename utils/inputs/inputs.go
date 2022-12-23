package inputs

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
	"github.com/alokmenghrajani/adventofcode2022/utils/grid2d"
	sparsegrid "github.com/alokmenghrajani/adventofcode2022/utils/sparseGrid"
)

func ToInts(input string, sep string) []int {
	var r []int
	for _, line := range strings.Split(input, sep) {
		if line != "" {
			r = append(r, utils.MustAtoi(line))
		}
	}
	return r
}

func ToGrid2D[T any](input, rowSep, colSep string, empty T, conv func(string) T) *grid2d.Grid[T] {
	lines := strings.Split(input, rowSep)

	grid := grid2d.NewGrid(len(lines[0]), len(lines), empty)
	for y, line := range lines {
		for x, v := range strings.Split(line, colSep) {
			grid.Set(x, y, conv(v))
		}
	}

	return grid
}

func ToSparseGrid[T comparable](input, rowSep, colSep string, empty T, conv func(string) T) *sparsegrid.SparseGrid[T] {
	lines := strings.Split(input, rowSep)

	grid := sparsegrid.NewGrid(empty)
	for y, line := range lines {
		for x, v := range strings.Split(line, colSep) {
			grid.Set(x, y, conv(v))
		}
	}

	return grid
}
