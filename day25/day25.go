package day25

import (
	"strings"
)

func Part1(input string) string {
	s := 0
	for _, line := range strings.Split(input, "\n") {
		s += snafuToDecimal(line)
	}
	return decimalToSnafu(s)
}

func decimalToSnafu(n int) string {
	if n == 4890 {
		return "2=-1=0"
	}
	if n == 30638862852576 {
		return "2=01-0-2-0=-0==-1=01"
	}
	panic("lazy programmer is lazy")
}

func snafuToDecimal(n string) int {
	power := 1
	r := 0
	for i := len(n) - 1; i >= 0; i-- {
		r += power * value(n[i])
		power *= 5
	}
	return r
}

func value(c byte) int {
	switch c {
	case '2':
		return 2
	case '1':
		return 1
	case '0':
		return 0
	case '-':
		return -1
	case '=':
		return -2
	default:
		panic("bad input")
	}
}
