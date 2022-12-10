package sparsegrid

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type SparseGrid[T comparable] struct {
	minX, maxX, minY, maxY int
	data                   map[string]T
	empty                  T
}

func NewGrid[T comparable](empty T) *SparseGrid[T] {
	return &SparseGrid[T]{
		minX:  utils.MaxInt,
		maxX:  utils.MinInt,
		minY:  utils.MaxInt,
		maxY:  utils.MinInt,
		data:  map[string]T{},
		empty: empty,
	}
}

func (g *SparseGrid[T]) SizeX() (int, int) {
	return g.minX, g.maxX
}

func (g *SparseGrid[T]) SizeY() (int, int) {
	return g.minY, g.maxY
}

func (g *SparseGrid[T]) Visited() int {
	return len(g.data)
}

func (g *SparseGrid[T]) Get(x, y int) T {
	k := key(x, y)
	v, ok := g.data[k]
	if !ok {
		return g.empty
	}
	return v
}

func (g *SparseGrid[T]) Set(x, y int, v T) {
	k := key(x, y)
	current, ok := g.data[k]
	if ok && v == current {
		return
	} else if !ok && v == g.empty {
		return
	} else if v == g.empty {
		delete(g.data, k)
	} else {
		g.data[k] = v
		g.minX = utils.Min(g.minX, x)
		g.maxX = utils.Max(g.maxX, x)
		g.minY = utils.Min(g.minY, y)
		g.maxY = utils.Max(g.maxY, y)
	}
}

func (g *SparseGrid[T]) StringWithFormatter(formatter func(T, int, int) string) string {
	var r strings.Builder
	for j := g.minY; j <= g.maxY; j++ {
		for i := g.minX; i <= g.maxX; i++ {
			_, err := r.WriteString(formatter(g.Get(i, j), i, j))
			utils.PanicOnErr(err)
		}
		_, err := r.WriteRune('\n')
		utils.PanicOnErr(err)
	}
	return r.String()
}

func key(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}
