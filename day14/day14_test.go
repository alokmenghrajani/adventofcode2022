package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`)
	assert.Equal(t, 24, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`)
	assert.Equal(t, 93, r)
}
