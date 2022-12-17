package day16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`)
	require.Equal(t, 1651, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`)
	require.Equal(t, 1707, r)
}

// func TestDebug1(t *testing.T) {
// 	r := common(`Valve AA has flow rate=0; tunnel leads to valve BB
// Valve BB has flow rate=10; tunnel leads to valves AA`, 1, 10)
// 	require.Equal(t, 80, r)
// }

// func TestDebug2(t *testing.T) {
// 	r := common(`Valve AA has flow rate=0; tunnel leads to valve BB
// Valve BB has flow rate=10; tunnel leads to valves AA`, 2, 10)
// 	require.Equal(t, 80, r)
// }

// func TestDebug3(t *testing.T) {
// 	r := common(`Valve AA has flow rate=0; tunnels lead to valves BB, CC
// Valve BB has flow rate=10; tunnel leads to valves AA
// Valve CC has flow rate=1; tunnel leads to valves AA`, 1, 10)
// 	require.Equal(t, 85, r)
// }

func TestDebug4(t *testing.T) {
	r := common(`Valve AA has flow rate=0; tunnels lead to valves BB, CC
Valve BB has flow rate=10; tunnel leads to valves AA
Valve CC has flow rate=1; tunnel leads to valves AA`, 2, 10)
	require.Equal(t, 88, r)
}

// time
// 1 move to BB and CC
// 2 open valve BB and CC
// 3, 4, 5, ... 10 release 11
// 4 move to CC, release 10
// 5 open valve CC, release 10
// 6, 7, 8, 9, 10 => release 11
// 3 release 10
// 4, 5, ... 10 release 10
// total: 80
