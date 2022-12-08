package grid2d

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type Grid[T any] struct {
	sizeX, sizeY int
	matrix       [][]T
	empty        T
}

func NewGrid[T any](sizeX, sizeY int, empty T) *Grid[T] {
	matrix := make([][]T, sizeY)
	rows := make([]T, sizeX*sizeY)
	for i := 0; i < sizeX*sizeY; i++ {
		rows[i] = empty
	}

	j := 0
	for i := 0; i < sizeY; i++ {
		matrix[i] = rows[j : j+sizeX : j+sizeX]
		j += sizeX
	}
	return &Grid[T]{
		sizeX:  sizeX,
		sizeY:  sizeY,
		matrix: matrix,
		empty:  empty,
	}
}

func (g *Grid[T]) SizeX() int {
	return g.sizeX
}

func (g *Grid[T]) SizeY() int {
	return g.sizeY
}

func (g *Grid[T]) Get(x, y int) T {
	if x < 0 || x >= g.sizeX {
		return g.empty
	}
	if y < 0 || y >= g.sizeY {
		return g.empty
	}
	return g.matrix[y][x]
}

func (g *Grid[T]) Set(x, y int, v T) {
	if x < 0 || x >= g.sizeX {
		panic("invalid x")
	}
	if y < 0 || y >= g.sizeY {
		panic("invalid y")
	}
	g.matrix[y][x] = v
}

func (g *Grid[T]) StringWithFormatter(formatter func(T, int, int) string) string {
	var r strings.Builder
	for j := 0; j < g.sizeY; j++ {
		for i := 0; i < g.sizeX; i++ {
			_, err := r.WriteString(formatter(g.matrix[j][i], i, j))
			utils.PanicOnErr(err)
		}
		_, err := r.WriteRune('\n')
		utils.PanicOnErr(err)
	}
	return r.String()
}
