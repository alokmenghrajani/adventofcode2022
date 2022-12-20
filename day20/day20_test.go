package day20

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`1
2
-3
3
-2
0
4`)
	require.Equal(t, 3, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`1
2
-3
3
-2
0
4`)
	require.Equal(t, 1623178306, r)
}
