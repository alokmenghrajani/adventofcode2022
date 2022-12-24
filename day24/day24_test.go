package day24

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`)
	require.Equal(t, 18, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`)
	require.Equal(t, 54, r)
}
