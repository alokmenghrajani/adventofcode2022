package inputs

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

func ToInts(input string, sep string) []int {
	var r []int
	for _, line := range strings.Split(input, sep) {
		if line != "" {
			r = append(r, utils.MustAtoi(line))
		}
	}
	return r
}
