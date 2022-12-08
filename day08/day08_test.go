package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(`30373
25512
65332
33549
35390`)
	assert.Equal(t, 21, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`30373
25512
65332
33549
35390`)
	assert.Equal(t, 8, r)
}
