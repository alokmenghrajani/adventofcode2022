package day20

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type ll struct {
	multiplier int
	size       int
	first      *node
	zero       *node
}

type node struct {
	v    int
	modv int
	next *node
	prev *node
}

func Part1(input string) int {
	ll := NewLL(1)
	for _, line := range strings.Split(input, "\n") {
		ll.add(utils.MustAtoi(line))
	}

	ll.mix(1)

	n1 := ll.advance(ll.zero, 1000)
	n2 := ll.advance(n1, 1000)
	n3 := ll.advance(n2, 1000)

	return n1.v + n2.v + n3.v
}

func Part2(input string) int {
	ll := NewLL(811589153)
	for _, line := range strings.Split(input, "\n") {
		ll.add(utils.MustAtoi(line))
	}

	ll.mix(10)

	n1 := ll.advance(ll.zero, 1000)
	n2 := ll.advance(n1, 1000)
	n3 := ll.advance(n2, 1000)

	return n1.v + n2.v + n3.v
}

func NewLL(multiplier int) *ll {
	return &ll{multiplier: multiplier}
}

func (ll *ll) add(v int) {
	node := &node{v: v * ll.multiplier, prev: nil, next: nil}
	if ll.first == nil {
		ll.first = node
		node.prev = node
		node.next = node
	} else {
		oldLast := ll.first.prev
		node.prev = oldLast
		node.next = ll.first
		ll.first.prev = node
		oldLast.next = node
	}
	if v == 0 {
		if ll.zero != nil {
			panic("multiple zeros?")
		}
		ll.zero = node
	}
	ll.size++
}

func (ll *ll) mix(rounds int) {
	originalList := make([]*node, 0, ll.size)
	c := ll.first
	for i := 0; i < ll.size; i++ {
		originalList = append(originalList, c)

		// I'm sure there's a cleaner way to calculate this but my brain is fried.
		if c.v > 0 {
			c.modv = c.v
			for c.modv >= ll.size {
				c.modv = (c.modv % ll.size) + c.modv/ll.size
			}
		} else if c.v < 0 {
			c.modv = -c.v
			for c.modv >= ll.size {
				c.modv = (c.modv % ll.size) + c.modv/ll.size
			}
			c.modv = -c.modv
		}
		c = c.next
	}

	for round := 0; round < rounds; round++ {
		for i := 0; i < ll.size; i++ {
			c := originalList[i]
			if c.modv == 0 {
				// nop
				// fmt.Printf("%d does not move\n", c.v)
				// ll.debug()
				// fmt.Println()
				continue
			}

			// remove c
			next := c.next
			prev := c.prev
			prev.next = next
			next.prev = prev

			// move by c.modv
			var t1 *node
			var t2 *node
			switch utils.Sign(c.modv) {
			case 1:
				t1 = c
				for j := 0; j < c.modv; j++ {
					t1 = t1.next
				}
				t2 = t1.next
			case -1:
				t2 = c
				for j := 0; j > c.modv; j-- {
					t2 = t2.prev
				}
				t1 = t2.prev
			default:
				panic("unreachable")
			}

			// add it in the new spot
			c.next = t2
			c.prev = t1
			t1.next = c
			t2.prev = c

			// fmt.Printf("%d moves between %d and %d\n", c.v, c.prev.v, c.next.v)
			// ll.debug()
			// fmt.Println()
		}
		// fmt.Printf("after %d rounds:\n", round+1)
		// ll.debug()
		// fmt.Println()
	}
}

func (ll *ll) advance(n *node, amount int) *node {
	amount = amount % ll.size
	for i := 0; i < amount; i++ {
		n = n.next
	}
	return n
}

// func (ll *ll) debug() {
// 	c := ll.first
// 	for i := 0; i < ll.size; i++ {
// 		if i > 0 {
// 			fmt.Printf(", ")
// 		}
// 		fmt.Printf("%d", c.v)
// 		c = c.next
// 	}
// 	fmt.Println()
// }
