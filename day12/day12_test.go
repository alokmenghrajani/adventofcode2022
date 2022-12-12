package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`)
	assert.Equal(t, 31, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`)
	assert.Equal(t, 29, r)
}
