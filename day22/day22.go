package day22

import (
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
	"github.com/alokmenghrajani/adventofcode2022/utils/grid2d"
)

type cell string

const (
	wall   cell = "#"
	open   cell = "."
	closed cell = " "
)

type dir int

const (
	right dir = 0
	down  dir = 1
	left  dir = 2
	up    dir = 3
)

type move struct {
	right bool
	left  bool
	move  int
}

type state struct {
	grid   *grid2d.Grid[cell]
	moves  []move
	posX   int
	posY   int
	dir    dir
	isTest bool
}

func Part1(input string) int {
	// parse
	pieces := strings.Split(input, "\n\n")
	s := &state{
		grid:  parseGrid(pieces[0]),
		moves: parseMoves(pieces[1]),
		posY:  0,
		dir:   right,
	}
	for i := 0; i < s.grid.SizeX(); i++ {
		if s.grid.Get(i, 0) == open {
			s.posX = i
			break
		}
	}

	// play the moves
	s.playMoves()

	return 1000*(s.posY+1) + 4*(s.posX+1) + int(s.dir)
}

func Part2(input string, isTest bool) int {
	// parse
	pieces := strings.Split(input, "\n\n")
	s := &state{
		grid:   parseGrid(pieces[0]),
		moves:  parseMoves(pieces[1]),
		posY:   0,
		dir:    right,
		isTest: isTest,
	}
	for i := 0; i < s.grid.SizeX(); i++ {
		if s.grid.Get(i, 0) == open {
			s.posX = i
			break
		}
	}

	// play the moves
	s.playMovesPart2()

	return 1000*(s.posY+1) + 4*(s.posX+1) + int(s.dir)
}

func parseGrid(input string) *grid2d.Grid[cell] {
	lines := strings.Split(input, "\n")
	maxX := 0
	for _, line := range lines {
		maxX = utils.Max(maxX, len(line))
	}

	g := grid2d.NewGrid(maxX, len(lines), closed)
	for j, line := range lines {
		for i, c := range line {
			g.Set(i, j, cell(c))
		}
	}

	return g
}

func parseMoves(input string) []move {
	r := []move{}
	re := regexp.MustCompile(`([0-9]+)|(L)|(R)`)
	for input != "" {
		pieces := re.FindStringSubmatch(input)
		if pieces[1] != "" {
			m := move{move: utils.MustAtoi(pieces[1])}
			r = append(r, m)
		} else if pieces[2] != "" {
			m := move{left: true}
			r = append(r, m)
		} else if pieces[3] != "" {
			m := move{right: true}
			r = append(r, m)
		} else {
			panic("bad input")
		}
		input = input[len(pieces[0]):]
	}
	return r
}

func (s *state) playMoves() {
	for _, m := range s.moves {
		s.playMove(m)
	}
}

func (s *state) playMove(m move) {
	if m.left {
		s.dir = (s.dir + 3) % 4
		return
	}
	if m.right {
		s.dir = (s.dir + 1) % 4
		return
	}
	dx := 0
	dy := 0
	switch s.dir {
	case right:
		dx = 1
	case down:
		dy = 1
	case left:
		dx = -1
	case up:
		dy = -1
	}
	for i := 0; i < m.move; i++ {
		newX := s.posX + dx
		newY := s.posY + dy
		c := s.grid.Get(newX, newY)
		switch c {
		case open:
			s.posX = newX
			s.posY = newY
		case wall:
			return
		case closed:
			// wrap around
			newX, newY, newD := s.wrap(newX, newY, dx, dy)
			c = s.grid.Get(newX, newY)
			switch c {
			case open:
				s.posX = newX
				s.posY = newY
				s.dir = newD
			case wall:
				return
			case closed:
				panic("unreachable")
			}
		}
	}
}

func (s *state) wrap(x, y, dx, dy int) (int, int, dir) {
	for {
		// go in reverse direction until we hit the other edge
		// this assumes the board is chasm-free.
		x = x - dx
		y = y - dy
		c := s.grid.Get(x, y)
		if c == closed {
			break
		}
	}
	return x + dx, y + dy, s.dir
}

func (s *state) playMovesPart2() {
	for _, m := range s.moves {
		s.playMovePart2(m)
	}
}

func (s *state) playMovePart2(m move) {
	if m.left {
		s.dir = (s.dir + 3) % 4
		return
	}
	if m.right {
		s.dir = (s.dir + 1) % 4
		return
	}
	for i := 0; i < m.move; i++ {
		if !s.playSingleMovePart2() {
			break
		}
	}
}

func (s *state) playSingleMovePart2() bool {
	faceId := s.faceId(s.posX, s.posY)
	if faceId == -1 {
		panic("invalid state")
	}
	dx := 0
	dy := 0
	switch s.dir {
	case right:
		dx = 1
	case down:
		dy = 1
	case left:
		dx = -1
	case up:
		dy = -1
	}
	newX := s.posX + dx
	newY := s.posY + dy
	newDir := s.dir
	newFaceId := s.faceId(newX, newY)
	if newFaceId != faceId || newFaceId == -1 {
		newX, newY, newDir = s.wrapPart2(faceId, s.posX, s.posY, s.dir)
	}
	c := s.grid.Get(newX, newY)
	switch c {
	case open:
		s.posX = newX
		s.posY = newY
		s.dir = newDir
		return true
	case wall:
		return false
	default:
		panic("unreachable")
	}
}

type point struct {
	x, y int
}

func (s *state) wrapPart2(faceId, x, y int, dir dir) (int, int, dir) {
	// we get here when we need to change faces
	if s.isTest {
		faces := []point{{8, 0}, {0, 4}, {4, 4}, {8, 4}, {8, 8}, {12, 8}}
		boxSize := 4

		relX := x - faces[faceId].x
		relY := y - faces[faceId].y
		switch faceId {
		case 0:
			switch dir {
			case right:
				newFace := 5
				newDir := left
				newRelX := relY
				newRelY := relX
				if newRelX != boxSize-1 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case down:
				return x, y + 1, dir
			case left:
				newFace := 2
				newDir := down
				newRelX := relY
				newRelY := relX
				if newRelY != 0 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case up:
				newFace := 1
				newDir := down
				newRelX := relX
				newRelY := relY
				if newRelY != 0 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			}
		case 1:
			switch dir {
			case right:
				return x + 1, y, dir
			case left:
				newFace := 3
				newDir := left
				newRelX := boxSize - relX - 1
				newRelY := relY
				if newRelX != boxSize-1 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			}
		case 2:
			switch dir {
			case up:
				newFace := 0
				newDir := right
				newRelX := relY
				newRelY := relX
				if newRelX != 0 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			}
		case 3:
			switch dir {
			case right:
				newFace := 5
				newDir := down
				newRelX := boxSize - relY - 1
				newRelY := boxSize - relX - 1
				if newRelY != 0 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case down:
				return x, y + 1, dir
			}
		case 4:
			switch dir {
			case down:
				newFace := 1
				newDir := up
				newRelX := boxSize - relX - 1
				newRelY := relY
				if newRelY != boxSize-1 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			}
		case 5:
			switch dir {
			case left:
				return x - 1, y, dir
			}
		}
		panic("todo")
	} else {
		faces := []point{{50, 0}, {100, 0}, {50, 50}, {50, 100}, {0, 100}, {0, 150}}
		boxSize := 50

		relX := x - faces[faceId].x
		relY := y - faces[faceId].y
		switch faceId {
		case 0:
			switch dir {
			case right:
				return x + 1, y, dir
			case down:
				return x, y + 1, dir
			case left:
				newFace := 4
				newDir := right
				newRelX := relX
				newRelY := boxSize - relY - 1
				if newRelX != 0 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case up:
				newFace := 5
				newDir := right
				newRelX := relY
				newRelY := relX
				if newRelX != 0 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			}
		case 1:
			switch dir {
			case right:
				newFace := 3
				newDir := left
				newRelX := relX
				newRelY := boxSize - relY - 1
				if newRelX != boxSize-1 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case down:
				newFace := 2
				newDir := left
				newRelX := relY
				newRelY := relX
				if newRelX != boxSize-1 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case left:
				return x - 1, y, dir
			case up:
				newFace := 5
				newDir := up
				newRelX := relX
				newRelY := boxSize - relY - 1
				if newRelY != boxSize-1 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			}
		case 2:
			switch dir {
			case right:
				newFace := 1
				newDir := up
				newRelX := relY
				newRelY := relX
				if newRelY != boxSize-1 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case down:
				return x, y + 1, dir
			case left:
				newFace := 4
				newDir := down
				newRelX := relY
				newRelY := relX
				if newRelY != 0 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case up:
				return x, y - 1, dir
			}
		case 3:
			switch dir {
			case right:
				newFace := 1
				newDir := left
				newRelX := relX
				newRelY := boxSize - relY - 1
				if newRelX != boxSize-1 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case down:
				newFace := 5
				newDir := left
				newRelX := relY
				newRelY := relX
				if newRelX != boxSize-1 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case left:
				return x - 1, y, dir
			case up:
				return x, y - 1, dir
			}
		case 4:
			switch dir {
			case right:
				return x + 1, y, dir
			case down:
				return x, y + 1, dir
			case left:
				newFace := 0
				newDir := right
				newRelX := relX
				newRelY := boxSize - relY - 1
				if newRelX != 0 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case up:
				newFace := 2
				newDir := right
				newRelX := relY
				newRelY := relX
				if newRelX != 0 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			}
		case 5:
			switch dir {
			case right:
				newFace := 3
				newDir := up
				newRelX := relY
				newRelY := relX
				if newRelY != boxSize-1 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case down:
				newFace := 1
				newDir := down
				newRelX := relX
				newRelY := boxSize - relY - 1
				if newRelY != 0 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case left:
				newFace := 0
				newDir := down
				newRelX := relY
				newRelY := relX
				if newRelY != 0 {
					panic("unexpected")
				}
				newX := faces[newFace].x + newRelX
				newY := faces[newFace].y + newRelY
				return newX, newY, newDir
			case up:
				return x, y - 1, dir
			}
		}
		panic("unreachable")
	}
}

func (s *state) faceId(x, y int) int {
	var faces []point
	var boxSize int
	if s.isTest {
		faces = []point{{8, 0}, {0, 4}, {4, 4}, {8, 4}, {8, 8}, {12, 8}}
		boxSize = 4
	} else {
		faces = []point{{50, 0}, {100, 0}, {50, 50}, {50, 100}, {0, 100}, {0, 150}}
		boxSize = 50
	}
	for id, face := range faces {
		if (face.x <= x) && ((face.x + boxSize) > x) && (face.y <= y) && ((face.y + boxSize) > y) {
			return id
		}
	}
	return -1
}

// func (s *state) debug() {
// 	r := s.grid.StringWithFormatter(func(c cell, i, j int) string {
// 		if i == s.posX && j == s.posY {
// 			return "*"
// 		}
// 		return string(c)
// 	})
// 	fmt.Println(r)
// 	fmt.Println()
// }
