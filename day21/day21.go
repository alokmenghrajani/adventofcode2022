package day21

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type op int

const (
	add op = iota
	sub
	mul
	div
	con
	processing
	eq
	unknown
)

type node struct {
	op   op
	arg1 string
	arg2 string
	v    int
}

func Part1(input string) int {
	data := parse(input)
	return eval(data, "root")
}

func Part2(input string) string {
	data := parse(input)
	data["root"].op = eq
	data["humn"].op = unknown
	data["humn"].arg1 = "x"
	n := solvePart2(data, "root")
	return n.String()
}

func parse(input string) map[string]*node {
	data := map[string]*node{}
	reAdd := regexp.MustCompile(`([a-z]+): ([a-z]+) [+] ([a-z]+)`)
	reSub := regexp.MustCompile(`([a-z]+): ([a-z]+) [-] ([a-z]+)`)
	reMul := regexp.MustCompile(`([a-z]+): ([a-z]+) [*] ([a-z]+)`)
	reDiv := regexp.MustCompile(`([a-z]+): ([a-z]+) [/] ([a-z]+)`)
	reCon := regexp.MustCompile(`([a-z]+): ([0-9]+)`)

	for _, line := range strings.Split(input, "\n") {
		match := reAdd.FindStringSubmatch(line)
		if len(match) > 0 {
			key := match[1]
			n := &node{
				op:   add,
				arg1: match[2],
				arg2: match[3],
			}
			data[key] = n
			continue
		}

		match = reSub.FindStringSubmatch(line)
		if len(match) > 0 {
			key := match[1]
			n := &node{
				op:   sub,
				arg1: match[2],
				arg2: match[3],
			}
			data[key] = n
			continue
		}

		match = reMul.FindStringSubmatch(line)
		if len(match) > 0 {
			key := match[1]
			n := &node{
				op:   mul,
				arg1: match[2],
				arg2: match[3],
			}
			data[key] = n
			continue
		}

		match = reDiv.FindStringSubmatch(line)
		if len(match) > 0 {
			key := match[1]
			n := &node{
				op:   div,
				arg1: match[2],
				arg2: match[3],
			}
			data[key] = n
			continue
		}

		match = reCon.FindStringSubmatch(line)
		if len(match) > 0 {
			key := match[1]
			n := &node{
				op: con,
				v:  utils.MustAtoi(match[2]),
			}
			data[key] = n
			continue
		}

		panic("invalid input")
	}
	return data
}

func eval(data map[string]*node, key string) int {
	n := data[key]
	switch n.op {
	case add:
		n.op = processing
		v1 := eval(data, n.arg1)
		v2 := eval(data, n.arg2)
		n.op = con
		n.v = v1 + v2
		return n.v
	case sub:
		n.op = processing
		v1 := eval(data, n.arg1)
		v2 := eval(data, n.arg2)
		n.op = con
		n.v = v1 - v2
		return n.v
	case mul:
		n.op = processing
		v1 := eval(data, n.arg1)
		v2 := eval(data, n.arg2)
		n.op = con
		n.v = v1 * v2
		return n.v
	case div:
		n.op = processing
		v1 := eval(data, n.arg1)
		v2 := eval(data, n.arg2)
		n.op = con
		n.v = v1 / v2
		return n.v
	case con:
		return n.v
	default:
		panic("meh")
	}
}

// Solve part2 by outputting an equation
func solvePart2(data map[string]*node, key string) *node {
	n := data[key]
	switch n.op {
	case eq:
		n.op = processing
		n1 := solvePart2(data, n.arg1)
		n2 := solvePart2(data, n.arg2)
		n.op = unknown
		n.arg1 = fmt.Sprintf("%s=%s", n1.String(), n2.String())
		return n
	case add:
		n.op = processing
		n1 := solvePart2(data, n.arg1)
		n2 := solvePart2(data, n.arg2)
		if n1.op == con && n2.op == con {
			n.op = con
			n.v = n1.v + n2.v
		} else {
			n.op = unknown
			n.arg1 = fmt.Sprintf("(%s+%s)", n1.String(), n2.String())
		}
		return n
	case sub:
		n.op = processing
		n1 := solvePart2(data, n.arg1)
		n2 := solvePart2(data, n.arg2)
		if n1.op == con && n2.op == con {
			n.op = con
			n.v = n1.v - n2.v
		} else {
			n.op = unknown
			n.arg1 = fmt.Sprintf("(%s-%s)", n1.String(), n2.String())
		}
		return n
	case mul:
		n.op = processing
		n1 := solvePart2(data, n.arg1)
		n2 := solvePart2(data, n.arg2)
		if n1.op == con && n2.op == con {
			n.op = con
			n.v = n1.v * n2.v
		} else {
			n.op = unknown
			n.arg1 = fmt.Sprintf("(%s*%s)", n1.String(), n2.String())
		}
		return n
	case div:
		n.op = processing
		n1 := solvePart2(data, n.arg1)
		n2 := solvePart2(data, n.arg2)
		if n1.op == con && n2.op == con {
			n.op = con
			n.v = n1.v / n2.v
		} else {
			n.op = unknown
			n.arg1 = fmt.Sprintf("(%s/%s)", n1.String(), n2.String())
		}
		return n
	case con:
		return n
	case unknown:
		return n
	default:
		panic("meh")
	}
}

func (n *node) String() string {
	switch n.op {
	case con:
		return fmt.Sprintf("%d", n.v)
	case unknown:
		return n.arg1
	default:
		panic("meh")
	}
}
