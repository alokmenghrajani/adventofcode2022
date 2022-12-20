package day19

import (
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type material int

const (
	ore material = iota
	clay
	obsidian
	geode
)

type blueprint struct {
	id    int
	rules map[material][]rule
	max   map[material]int
}

type rule struct {
	needed material
	amount int
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	r := 0
	for _, line := range lines {
		// parse input
		blueprint := parse(line)

		// compute solution
		s := search(blueprint, 24)
		r += blueprint.id * s
	}
	return r
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	if len(lines) > 3 {
		lines = lines[:3]
	}
	r := 1
	for _, line := range lines {
		// parse input
		blueprint := parse(line)

		// compute solution
		s := search(blueprint, 32)
		r *= s
	}
	return r
}

func parse(line string) blueprint {
	re1 := regexp.MustCompile(`Blueprint (\d+):`)
	re2 := regexp.MustCompile(`Each ([a-z]+) robot costs (\d+) ([a-z]+)[.]`)
	re3 := regexp.MustCompile(`Each ([a-z]+) robot costs (\d+) ([a-z]+) and (\d+) ([a-z]+)[.]`)
	id := re1.FindStringSubmatch(line)
	b := blueprint{
		id:    utils.MustAtoi(id[1]),
		rules: map[material][]rule{},
		max:   map[material]int{},
	}
	pieces := re2.FindAllStringSubmatch(line, -1)
	for _, piece := range pieces {
		r := []rule{{amount: utils.MustAtoi(piece[2]), needed: parseMaterial(piece[3])}}
		b.rules[parseMaterial(piece[1])] = r
	}
	pieces = re3.FindAllStringSubmatch(line, -1)
	for _, piece := range pieces {
		r := []rule{{amount: utils.MustAtoi(piece[2]), needed: parseMaterial(piece[3])},
			{amount: utils.MustAtoi(piece[4]), needed: parseMaterial(piece[5])}}
		b.rules[parseMaterial(piece[1])] = r
	}

	return b
}

func parseMaterial(n string) material {
	switch n {
	case "ore":
		return ore
	case "clay":
		return clay
	case "obsidian":
		return obsidian
	case "geode":
		return geode
	default:
		panic("bad input")
	}
}

type state struct {
	robots     [4]int
	production [4]int
}

// DFS search. Keep track of best known solution to prune search space. Also
// don't revisit states and prune states with too many robots of the same kind.
func search(blueprint blueprint, minutes int) int {
	best := 0
	robots := [4]int{1, 0, 0, 0}
	production := [4]int{0, 0, 0, 0}
	seen := map[state]int{}

	for _, rules := range blueprint.rules {
		for _, rule := range rules {
			blueprint.max[rule.needed] = utils.Max(blueprint.max[rule.needed], rule.amount)
		}
	}

	search2(blueprint, minutes, robots, production, seen, &best)
	return int(best)
}

func search2(blueprint blueprint, minutes int, robots [4]int, production [4]int, seen map[state]int, best *int) {
	if minutes == 0 {
		if production[geode] > *best {
			*best = production[geode]
		}
		return
	}

	// check if this search needs to be pruned. We severly overestimate (since it's not possible)
	// to build a geode robot every minute. A better estimation will speed things up!
	r := robots[geode]
	p := production[geode]
	for m := 0; m < minutes; m++ {
		r++
		p += r
		if p > *best {
			break
		}
	}
	if p <= *best {
		return
	}

	// reject wasteful states
	for i := ore; i <= obsidian; i++ {
		if robots[i] > blueprint.max[i] {
			return
		}
	}

	// cache states we have already visited
	t := state{robots, production}
	v, ok := seen[t]
	if ok && v >= minutes {
		return
	}
	seen[t] = minutes

	// search all possibilities
outer:
	for m := geode; m >= ore; m-- {
		rules := blueprint.rules[m]
		for _, rule := range rules {
			if production[rule.needed] < rule.amount {
				continue outer
			}
		}

		newProduction := [4]int{}
		copy(newProduction[:], production[:])
		for _, rule := range rules {
			newProduction[rule.needed] -= rule.amount
		}
		for i := 0; i < 4; i++ {
			newProduction[i] += robots[i]
		}
		newRobots := [4]int{}
		copy(newRobots[:], robots[:])
		newRobots[m]++

		search2(blueprint, minutes-1, newRobots, newProduction, seen, best)
	}

	newProduction := [4]int{}
	copy(newProduction[:], production[:])
	for i := 0; i < 4; i++ {
		newProduction[i] += robots[i]
	}

	search2(blueprint, minutes-1, robots, newProduction, seen, best)
}
