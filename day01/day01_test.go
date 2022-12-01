package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)
	assert.Equal(t, 24000, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)
	assert.Equal(t, 45000, r)
}
