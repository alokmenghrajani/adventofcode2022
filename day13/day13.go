package day13

import (
	"fmt"
	"sort"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type v struct {
	values    []v
	i         int
	isDivider bool
}

func Part1(input string) int {
	r := 0
	for i, lines := range strings.Split(input, "\n\n") {
		pairs := strings.Split(lines, "\n")
		v1, lo1 := parse(pairs[0])
		if lo1 != "" {
			panic(fmt.Errorf("left over: %s", lo1))
		}
		v2, lo2 := parse(pairs[1])
		if lo2 != "" {
			panic(fmt.Errorf("left over: %s", lo2))
		}
		t := compare(v1, v2)
		if t == nil {
			panic("compare returned nil")
		}
		if *t {
			r += (i + 1)
		}
	}
	return r
}

func Part2(input string) int {
	packets := []v{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		v1, lo1 := parse(line)
		if lo1 != "" {
			panic(fmt.Errorf("left over: %s", lo1))
		}
		packets = append(packets, v1)
	}

	// add [[2]]
	v1, lo1 := parse("[[2]]")
	if lo1 != "" {
		panic(fmt.Errorf("left over: %s", lo1))
	}
	v1.isDivider = true
	packets = append(packets, v1)

	// add [[6]]
	v1, lo1 = parse("[[6]]")
	if lo1 != "" {
		panic(fmt.Errorf("left over: %s", lo1))
	}
	v1.isDivider = true
	packets = append(packets, v1)

	sort.Slice(packets, func(i, j int) bool {
		t := compare(packets[i], packets[j])
		if t == nil {
			panic("compare returned nil")
		}
		return *t
	})

	r := 1
	for i, v := range packets {
		if v.isDivider {
			r *= (i + 1)
		}
	}
	return r
}

func parse(line string) (v, string) {
	if line[0] == '[' {
		current := v{values: []v{}}
		t := line[1:]
		firstTime := true
		for {
			if t[0] == ']' {
				break
			}
			if !firstTime {
				if t[0] != ',' {
					panic("oops")
				}
				t = t[1:]
			}
			firstTime = false
			v, leftOver := parse(t)
			t = leftOver
			current.values = append(current.values, v)
		}
		return current, t[1:]
	}
	i := 0
	for ; i < len(line); i++ {
		if line[i] < '0' || line[i] > '9' {
			break
		}
	}
	return v{i: utils.MustAtoi(line[0:i])}, line[i:]
}

func compare(v1, v2 v) *bool {
	t := true
	f := false
	if v1.values == nil && v2.values == nil {
		if v1.i < v2.i {
			return &t
		}
		if v1.i > v2.i {
			return &f
		}
		return nil
	}
	if v1.values == nil && v2.values != nil {
		temp := v{values: []v{v1}}
		return compare(temp, v2)
	}
	if v1.values != nil && v2.values == nil {
		temp := v{values: []v{v2}}
		return compare(v1, temp)
	}
	for i := 0; i < len(v1.values) || i < len(v2.values); i++ {
		if i >= len(v1.values) {
			return &t
		}
		if i >= len(v2.values) {
			return &f
		}
		temp := compare(v1.values[i], v2.values[i])
		if temp != nil {
			return temp
		}
	}
	return nil
}
