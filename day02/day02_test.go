package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(`A Y
B X
C Z`)
	assert.Equal(t, 15, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`A Y
B X
C Z`)
	assert.Equal(t, 12, r)
}
