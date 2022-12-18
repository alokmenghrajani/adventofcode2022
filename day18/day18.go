package day18

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type point struct {
	x, y, z int
}

func Part1(input string) int {
	cells := map[point]bool{}
	for _, line := range strings.Split(input, "\n") {
		cells[parse(line)] = true
	}

	r := 0
	for cell := range cells {
		for _, neighbor := range neighbors(cell) {
			if _, ok := cells[neighbor]; !ok {
				r++
			}
		}
	}
	return r
}

func Part2(input string) int {
	// parse input
	cells := map[point]bool{}
	for _, line := range strings.Split(input, "\n") {
		cells[parse(line)] = true
	}

	// compute bounds
	boundsMin := point{utils.MaxInt, utils.MaxInt, utils.MaxInt}
	boundsMax := point{utils.MinInt, utils.MinInt, utils.MinInt}
	for cell := range cells {
		boundsMin.x = utils.Min(boundsMin.x, cell.x)
		boundsMin.y = utils.Min(boundsMin.y, cell.y)
		boundsMin.z = utils.Min(boundsMin.z, cell.z)

		boundsMax.x = utils.Max(boundsMax.x, cell.x)
		boundsMax.y = utils.Max(boundsMax.y, cell.y)
		boundsMax.z = utils.Max(boundsMax.z, cell.z)
	}

	// calculate number of neighbors which aren't inside
	insideOutside := map[point]state{}
	r := 0
	for cell := range cells {
		for _, neighbor := range neighbors(cell) {
			if _, ok := cells[neighbor]; !ok {
				s := isInside(cells, insideOutside, boundsMin, boundsMax, neighbor)
				switch s {
				case inside:
					// do nothing
				case outside:
					r++
				default:
					panic("meh")
				}
			}
		}
	}
	return r
}

func parse(line string) point {
	pieces := strings.Split(line, ",")
	x := utils.MustAtoi(pieces[0])
	y := utils.MustAtoi(pieces[1])
	z := utils.MustAtoi(pieces[2])
	return point{x, y, z}
}

func neighbors(cell point) []point {
	moves := []point{
		{1, 0, 0}, {-1, 0, 0},
		{0, 1, 0}, {0, -1, 0},
		{0, 0, 1}, {0, 0, -1}}
	r := []point{}
	for _, move := range moves {
		x := cell.x + move.x
		y := cell.y + move.y
		z := cell.z + move.z
		r = append(r, point{x, y, z})
	}
	return r
}

type state int

const (
	inside state = iota
	outside
	processing
)

func isInside(cells map[point]bool, insideOutside map[point]state, boundsMin, boundsMax, cell point) state {
	if cell.x < boundsMin.x || cell.y < boundsMin.y || cell.z < boundsMin.z {
		return outside
	}
	if cell.x > boundsMax.x || cell.y > boundsMax.y || cell.z > boundsMax.z {
		return outside
	}
	if v, ok := insideOutside[cell]; ok {
		return v
	}
	if cells[cell] {
		return processing
	}
	insideOutside[cell] = processing
	moves := []point{
		{1, 0, 0}, {-1, 0, 0},
		{0, 1, 0}, {0, -1, 0},
		{0, 0, 1}, {0, 0, -1}}
	for _, move := range moves {
		c := point{cell.x + move.x, cell.y + move.y, cell.z + move.z}
		v := isInside(cells, insideOutside, boundsMin, boundsMax, c)
		switch v {
		case inside:
			insideOutside[cell] = inside
			return inside
		case outside:
			insideOutside[cell] = outside
			return outside
		}
	}
	insideOutside[cell] = inside
	return inside
}
