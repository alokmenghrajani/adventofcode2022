package day23

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`)
	require.Equal(t, 110, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`)
	require.Equal(t, 20, r)
}
