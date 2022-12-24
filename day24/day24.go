package day24

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils/grid2d"
)

type cell string

const (
	right cell = ">"
	down  cell = "v"
	left  cell = "<"
	up    cell = "^"
	empty cell = "."
)

type state struct {
	minute     int
	posX, posY int
}

func Part1(input string) int {
	// parse input
	lines := strings.Split(input, "\n")
	sizeX := len(lines[0]) - 2
	sizeY := len(lines) - 2
	grid := grid2d.NewGrid(sizeX, sizeY, empty)
	for j := 0; j < sizeY; j++ {
		line := lines[j+1]
		for i := 0; i < sizeX; i++ {
			c := line[i+1]
			grid.Set(i, j, cell(c))
		}
	}

	// use BFS to find solution
	return solve(grid, [][]int{{sizeX - 1, sizeY}})
}

func Part2(input string) int {
	// parse input
	lines := strings.Split(input, "\n")
	sizeX := len(lines[0]) - 2
	sizeY := len(lines) - 2
	grid := grid2d.NewGrid(sizeX, sizeY, empty)
	for j := 0; j < sizeY; j++ {
		line := lines[j+1]
		for i := 0; i < sizeX; i++ {
			c := line[i+1]
			grid.Set(i, j, cell(c))
		}
	}

	// use BFS to find solution
	return solve(grid, [][]int{{sizeX - 1, sizeY}, {0, -1}, {sizeX - 1, sizeY}})
}

func solve(grid *grid2d.Grid[cell], goals [][]int) int {
	queue := []state{{minute: 0, posX: 0, posY: -1}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.posX == goals[0][0] && current.posY == goals[0][1] {
			// we reached a goal
			goals = goals[1:]
			if len(goals) == 0 {
				// we are done
				return current.minute
			}
			// empty queue but preserve memory
			queue = queue[:0]
		}

		// check where we can move
		moves := [][]int{{1, 0}, {0, 1}, {0, 0}, {-1, 0}, {0, -1}}
		for _, move := range moves {
			x := current.posX + move[0]
			y := current.posY + move[1]
			if valid(current.minute+1, x, y, grid) {
				newState := state{minute: current.minute + 1, posX: x, posY: y}
				// only add state if it's not been seen
				add := true
				for i := len(queue) - 1; i >= 0; i-- {
					t := queue[i]
					if t.minute < newState.minute {
						break
					}
					if t.minute == newState.minute && t.posX == newState.posX && t.posY == newState.posY {
						add = false
						break
					}
				}
				if add {
					queue = append(queue, newState)
				}
			}
		}
	}
	panic("no solution")
}

// func debug(s state, grid *grid2d.Grid[cell]) {
// 	fmt.Printf("minute: %d\n", s.minute)
// 	r := grid.StringWithFormatter(func(c cell, i, j int) string {
// 		if i == s.posX && j == s.posY {
// 			return "*"
// 		}
// 		return string(c)
// 	})
// 	fmt.Println(r)
// 	fmt.Println()
// }

type disallow struct {
	dx, dy int
	cell   cell
}

func valid(minute, x, y int, grid *grid2d.Grid[cell]) bool {
	if x == grid.SizeX()-1 && y == grid.SizeY() {
		return true
	}
	if x == 0 && y == -1 {
		return true
	}
	if x < 0 || x >= grid.SizeX() {
		return false
	}
	if y < 0 || y >= grid.SizeY() {
		return false
	}
	disallows := []disallow{{-1, 0, ">"}, {1, 0, "<"}, {0, -1, "v"}, {0, 1, "^"}}
	for _, disallow := range disallows {
		cx := positiveMod(x+disallow.dx*minute, grid.SizeX())
		cy := positiveMod(y+disallow.dy*minute, grid.SizeY())
		if grid.Get(cx, cy) == disallow.cell {
			return false
		}
	}
	return true
}

func positiveMod(x, m int) int {
	r := x % m
	if r >= 0 {
		return r
	}
	return r + m
}
