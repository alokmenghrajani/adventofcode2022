package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`)
	assert.Equal(t, 13, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`)
	assert.Equal(t, 36, r)
}

func TestUpdateTail(t *testing.T) {
	r := &rope{
		size: 2,
		posX: []int{3, 1},
		posY: []int{1, 1},
	}
	r.updateTail()
	assert.Equal(t, 2, r.posX[1])
	assert.Equal(t, 1, r.posY[1])

	r = &rope{
		size: 2,
		posX: []int{1, 1},
		posY: []int{3, 1},
	}
	r.updateTail()
	assert.Equal(t, 1, r.posX[1])
	assert.Equal(t, 2, r.posY[1])

	r = &rope{
		size: 2,
		posX: []int{2, 1},
		posY: []int{1, 3},
	}
	r.updateTail()
	assert.Equal(t, 2, r.posX[1])
	assert.Equal(t, 2, r.posY[1])

	r = &rope{
		size: 2,
		posX: []int{3, 1},
		posY: []int{2, 3},
	}
	r.updateTail()
	assert.Equal(t, 2, r.posX[1])
	assert.Equal(t, 2, r.posY[1])
}
