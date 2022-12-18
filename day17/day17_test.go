package day17

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`)
	require.Equal(t, 3068, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`)
	require.Equal(t, 1514285714288, r)
}
