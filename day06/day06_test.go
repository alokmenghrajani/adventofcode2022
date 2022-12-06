package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 7, Part1(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`))
	assert.Equal(t, 5, Part1(`bvwbjplbgvbhsrlpgdmjqwftvncz`))
	assert.Equal(t, 6, Part1(`nppdvjthqldpwncqszvftbrmjlhg`))
	assert.Equal(t, 10, Part1(`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`))
	assert.Equal(t, 11, Part1(`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 19, Part2(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`))
	assert.Equal(t, 23, Part2(`bvwbjplbgvbhsrlpgdmjqwftvncz`))
	assert.Equal(t, 23, Part2(`nppdvjthqldpwncqszvftbrmjlhg`))
	assert.Equal(t, 29, Part2(`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`))
	assert.Equal(t, 26, Part2(`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`))
}
