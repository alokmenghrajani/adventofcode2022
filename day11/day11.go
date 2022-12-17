package day11

import (
	"math/big"
	"sort"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type op int

const (
	addition op = iota
	multiplication
	square
)

type monkey struct {
	items             []*big.Int
	inspected         int
	operationType     op
	operationArgument *big.Int
	test              *big.Int
	ifTrue            int
	ifFalse           int
}

func Part1(input string) int {
	return common(input, big.NewInt(3), 20)
}

func Part2(input string) int {
	return common(input, nil, 10000)
}

func common(input string, div *big.Int, rounds int) int {
	monkeys := []*monkey{}
	var curMonkey *monkey

	// parse input
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "Monkey") {
			curMonkey = &monkey{items: []*big.Int{}}
			monkeys = append(monkeys, curMonkey)
		} else if strings.HasPrefix(line, "  Starting items: ") {
			s := strings.TrimPrefix(line, "  Starting items: ")
			for _, piece := range strings.Split(s, ", ") {
				t := utils.MustAtoi(piece)
				curMonkey.items = append(curMonkey.items, big.NewInt(int64(t)))
			}
		} else if strings.HasPrefix(line, "  Operation: new = ") {
			s := strings.TrimPrefix(line, "  Operation: new = ")
			if strings.HasPrefix(s, "old + ") {
				s2 := strings.TrimPrefix(s, "old + ")
				curMonkey.operationType = addition
				t := utils.MustAtoi(s2)
				curMonkey.operationArgument = big.NewInt(int64(t))
			} else if s == "old * old" {
				curMonkey.operationType = square
			} else if strings.HasPrefix(s, "old * ") {
				s2 := strings.TrimPrefix(s, "old * ")
				curMonkey.operationType = multiplication
				t := utils.MustAtoi(s2)
				curMonkey.operationArgument = big.NewInt(int64(t))
			}
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			s := strings.TrimPrefix(line, "  Test: divisible by ")
			t := utils.MustAtoi(s)
			curMonkey.test = big.NewInt(int64(t))
		} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
			s := strings.TrimPrefix(line, "    If true: throw to monkey ")
			curMonkey.ifTrue = utils.MustAtoi(s)
		} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
			s := strings.TrimPrefix(line, "    If false: throw to monkey ")
			curMonkey.ifFalse = utils.MustAtoi(s)
		} else if line == "" {
			// do nothing
		} else {
			panic("unknown input")
		}
	}

	// compute common divisor to optimize computation for part 2
	// note: all the monkey.test are prime numbers, I should have simply
	// multiplied all the values together.
	commonDivisor := big.NewInt(1)
	for _, monkey := range monkeys {
		t := &big.Int{}
		t.GCD(nil, nil, commonDivisor, monkey.test)
		commonDivisor.Mul(commonDivisor, monkey.test)
		commonDivisor.Div(commonDivisor, t)
	}

	for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			monkey.inspect(monkeys, commonDivisor, div)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})

	return monkeys[0].inspected * monkeys[1].inspected
}

func (m *monkey) inspect(monkeys []*monkey, commonDivisor *big.Int, div *big.Int) {
	for len(m.items) > 0 {
		item := m.items[0]
		m.items = m.items[1:]
		m.inspected++
		switch m.operationType {
		case addition:
			item.Add(item, m.operationArgument)
		case multiplication:
			item.Mul(item, m.operationArgument)
		case square:
			item.Mul(item, item)
		}
		item.Mod(item, commonDivisor)
		if div != nil {
			item.Div(item, div)
		}
		t := &big.Int{}
		t.Mod(item, m.test)
		if t.Sign() == 0 {
			monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, item)
		} else {
			monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, item)
		}
	}
}
