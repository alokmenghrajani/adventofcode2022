package day16

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type node struct {
	rate      int
	neighbors map[string]int
	visited   bool
}

func Part1(input string) int {
	return common(input, 1, 30)
}

func Part2(input string) int {
	return common(input, 2, 26)
}

// original solution for part1
//
// func solvePart1(nodes map[string]*node, current string, flow, totalflow, time int) int {
// 	if time < 0 {
// 		return 0
// 	}
// 	if time == 0 {
// 		return totalflow
// 	}
// 	if current != "AA" {
// 		totalflow += flow
// 		flow += nodes[current].rate
// 		time--
// 	}
// 	if time == 0 {
// 		return totalflow
// 	}
// 	// if we don't do anything
// 	best := totalflow + flow*time
// 	for name, distance := range nodes[current].neighbors {
// 		n := nodes[name]
// 		if n.visited {
// 			continue
// 		}
// 		n.visited = true
// 		t := solve(nodes, name, flow, totalflow+flow*distance, time-distance)
// 		n.visited = false
// 		if t > best {
// 			best = t
// 		}
// 	}
// 	return best
// }

type character struct {
	movingTo string
	timer    int
}

func common(input string, characters int, time int) int {
	nodes := map[string]*node{}

	// parse input
	re := regexp.MustCompile(`Valve (..) has flow rate=([0-9]+); tunnels? leads? to valves? (.*)`)
	for _, line := range strings.Split(input, "\n") {
		pieces := re.FindStringSubmatch(line)
		n := &node{
			rate:      utils.MustAtoi(pieces[2]),
			neighbors: map[string]int{},
		}
		for _, name := range strings.Split(pieces[3], ", ") {
			n.neighbors[name] = 1
		}
		nodes[pieces[1]] = n
	}

	// compute closest distance between all nodes
	for k1, n1 := range nodes {
		// add self as 0
		n1.neighbors[k1] = 0
		for k2 := range nodes {
			if _, ok := n1.neighbors[k2]; !ok {
				n1.neighbors[k2] = utils.MaxInt
			}
		}
	}

	done := false
	for !done {
		done = true
		for _, n1 := range nodes {
			for name1, v1 := range n1.neighbors {
				if v1 == utils.MaxInt {
					continue
				}
				for name2, v2 := range nodes[name1].neighbors {
					if v2 == utils.MaxInt {
						continue
					}
					if n1.neighbors[name2] > v1+v2 {
						n1.neighbors[name2] = v1 + v2
						done = false
					}
				}
			}
		}
	}

	// remove nodes with zero rate but keep "AA" since that's where we start
	nodesToRemove := []string{}
	for name, n := range nodes {
		if n.rate == 0 && name != "AA" {
			nodesToRemove = append(nodesToRemove, name)
		}
	}
	for _, name := range nodesToRemove {
		delete(nodes, name)
	}
	for name1, n := range nodes {
		for _, name2 := range nodesToRemove {
			delete(n.neighbors, name2)
		}
		// remove self
		delete(n.neighbors, name1)
	}

	// Search for the solution. If needed, we can prune states (since we can calculate lower and upper bounds given the current state).
	nodes["AA"].visited = true
	c := []character{}
	for i := 0; i < characters; i++ {
		c = append(c, character{movingTo: "AA", timer: 0})
	}
	r := solvePart2(nodes, c, 0, 0, time, 0)
	return r
}

var bestest int

func solvePart2(nodes map[string]*node, characters []character, flow, totalflow, time, depth int) int {
	debug("entry", characters, flow, totalflow, time, depth)

	// check for movement
	for id, c := range characters {
		if c.timer == 0 {
			currentName := c.movingTo
			flow += nodes[c.movingTo].rate
			best := 0
			for name, distance := range nodes[c.movingTo].neighbors {
				n := nodes[name]
				if n.visited {
					continue
				}
				debug(fmt.Sprintf("%d is going to move to %s", id, name), characters, flow, totalflow, time, depth)
				n.visited = true
				characters[id].movingTo = name
				characters[id].timer = distance + 1
				t := solvePart2(nodes, characters, flow, totalflow, time, depth+1)
				characters[id].movingTo = currentName
				characters[id].timer = 0
				n.visited = false
				if t > best {
					best = t
				}
			}
			// alternative: stay still
			characters[id].timer = utils.MaxInt
			t := solvePart2(nodes, characters, flow, totalflow, time, depth+1)
			characters[id].timer = 0
			if t > best {
				best = t
			}
			debug(fmt.Sprintf("best: %d", best), characters, flow, totalflow, time, depth)
			return best
		}
	}

	// make a copy of the characters array
	characters2 := []character{}
	characters2 = append(characters2, characters...)

	// decrement time by the max amount
	timeDecrement := time
	for _, c := range characters2 {
		timeDecrement = utils.Min(timeDecrement, c.timer)
	}
	for id := range characters2 {
		characters2[id].timer -= timeDecrement
		if characters2[id].timer < 0 {
			panic("neg timer")
		}
	}
	totalflow += flow * timeDecrement
	time -= timeDecrement
	if time == 0 {
		debug("time=0", characters2, flow, totalflow, time, depth)
		if totalflow > bestest {
			bestest = totalflow
			fmt.Printf("bestest: %d\n", bestest)
		}
		return totalflow
	}
	if time < 0 {
		panic("neg time")
	}

	debug("iterating", characters2, flow, totalflow, time, depth)
	r := solvePart2(nodes, characters2, flow, totalflow, time, depth)
	return r
}

func debug(msg string, characters []character, flow, totalflow, time, depth int) {
	// fmt.Printf("%s", strings.Repeat(" ", depth*2))
	// fmt.Printf("%s: ", msg)
	// fmt.Printf("(time: %d, flow: %d, totalflow: %d)\n", time, flow, totalflow)

	// for id, c := range characters {
	// 	fmt.Printf("%s", strings.Repeat(" ", depth*2))
	// 	fmt.Printf("%d: movingTo: %s, timer: %d\n", id, c.movingTo, c.timer)
	// }
	// fmt.Println()
}
