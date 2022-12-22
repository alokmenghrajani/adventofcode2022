package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/alokmenghrajani/adventofcode2022/day01"
	"github.com/alokmenghrajani/adventofcode2022/day02"
	"github.com/alokmenghrajani/adventofcode2022/day03"
	"github.com/alokmenghrajani/adventofcode2022/day04"
	"github.com/alokmenghrajani/adventofcode2022/day05"
	"github.com/alokmenghrajani/adventofcode2022/day06"
	"github.com/alokmenghrajani/adventofcode2022/day07"
	"github.com/alokmenghrajani/adventofcode2022/day08"
	"github.com/alokmenghrajani/adventofcode2022/day09"
	"github.com/alokmenghrajani/adventofcode2022/day10"
	"github.com/alokmenghrajani/adventofcode2022/day11"
	"github.com/alokmenghrajani/adventofcode2022/day12"
	"github.com/alokmenghrajani/adventofcode2022/day13"
	"github.com/alokmenghrajani/adventofcode2022/day14"
	"github.com/alokmenghrajani/adventofcode2022/day15"
	"github.com/alokmenghrajani/adventofcode2022/day16"
	"github.com/alokmenghrajani/adventofcode2022/day17"
	"github.com/alokmenghrajani/adventofcode2022/day18"
	"github.com/alokmenghrajani/adventofcode2022/day19"
	"github.com/alokmenghrajani/adventofcode2022/day20"
	"github.com/alokmenghrajani/adventofcode2022/day21"
	"github.com/alokmenghrajani/adventofcode2022/utils"
)

// Usage: go run main.go <NN>
// assumes input is in day<NN>/input.txt
func main() {
	d := day()
	fmt.Printf("Running day %02d\n", d)

	switch d {
	case 1:
		fmt.Printf("part 1: %d\n", day01.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day01.Part2(utils.Readfile(d)))
	case 2:
		fmt.Printf("part 1: %d\n", day02.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day02.Part2(utils.Readfile(d)))
	case 3:
		fmt.Printf("part 1: %d\n", day03.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day03.Part2(utils.Readfile(d)))
	case 4:
		fmt.Printf("part 1: %d\n", day04.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day04.Part2(utils.Readfile(d)))
	case 5:
		fmt.Printf("part 1: %s\n", day05.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %s\n", day05.Part2(utils.Readfile(d)))
	case 6:
		fmt.Printf("part 1: %d\n", day06.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day06.Part2(utils.Readfile(d)))
	case 7:
		fmt.Printf("part 1: %d\n", day07.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day07.Part2(utils.Readfile(d)))
	case 8:
		fmt.Printf("part 1: %d\n", day08.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day08.Part2(utils.Readfile(d)))
	case 9:
		fmt.Printf("part 1: %d\n", day09.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day09.Part2(utils.Readfile(d)))
	case 10:
		fmt.Printf("part 1: %d\n", day10.Part1(utils.Readfile(d)))
		fmt.Printf("part 2:\n%s\n", day10.Part2(utils.Readfile(d)))
	case 11:
		fmt.Printf("part 1: %d\n", day11.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day11.Part2(utils.Readfile(d)))
	case 12:
		fmt.Printf("part 1: %d\n", day12.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day12.Part2(utils.Readfile(d)))
	case 13:
		fmt.Printf("part 1: %d\n", day13.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day13.Part2(utils.Readfile(d)))
	case 14:
		fmt.Printf("part 1: %d\n", day14.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day14.Part2(utils.Readfile(d)))
	case 15:
		fmt.Printf("part 1: %d\n", day15.Part1(utils.Readfile(d), 2000000))
		fmt.Printf("part 2: %d\n", day15.Part2(utils.Readfile(d), 4000000))
	case 16:
		fmt.Printf("part 1: %d\n", day16.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day16.Part2(utils.Readfile(d)))
	case 17:
		fmt.Printf("part 1: %d\n", day17.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day17.Part2(utils.Readfile(d)))
	case 18:
		fmt.Printf("part 1: %d\n", day18.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day18.Part2(utils.Readfile(d)))
	case 19:
		fmt.Printf("part 1: %d\n", day19.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day19.Part2(utils.Readfile(d)))
	case 20:
		fmt.Printf("part 1: %d\n", day20.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day20.Part2(utils.Readfile(d)))
	case 21:
		fmt.Printf("part 1: %d\n", day21.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %s\n", day21.Part2(utils.Readfile(d)))
	default:
		panic(fmt.Errorf("no such day: %d", d))
	}
}

// Reads day from os.Args.
func day() int {
	latest := 21
	if len(os.Args) == 1 {
		return latest
	}
	if os.Args[1] == "tree" {
		genTree()
		os.Exit(0)
	}
	if os.Args[1] == "next" {
		genNext(latest + 1)
		os.Exit(0)
	}
	day := utils.MustAtoi(os.Args[1])
	return day
}

func genTree() {
	rand.Seed(time.Now().Unix())
	f, err := os.Create("tree.svg")
	utils.PanicOnErr(err)
	defer f.Close()
	fmt.Fprintln(f, `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 500 600" width="500" height="600">
  <foreignObject width="100%" height="100%">
    <div xmlns="http://www.w3.org/1999/xhtml">
      <style>
* {
  font-family: monospace;
  color: #0f0;
  text-shadow: 0 0 5px #fff;
  background-color: #000;
}
.s {
  color: #ff0;
}
.c0 {
  animation: anim 3s infinite;
}
.c1 {
  animation: anim 3s infinite 1.2s;
}
.c2 {
  animation: anim 3s infinite 2.3s;
}
@keyframes anim {
  0%   { color: #f00; }
  32% { color: #f00; }

  33% { color: #f70; }
  65% {color: #f70; }

  66% { color: #00f;}
  100% { color: #00f;}
}
.t {
  color: #a00;
}
pre {
  padding: 2em;
}
      </style><pre>`)
	n := 27
	len := -1
	for i := 0; i < n; i++ {
		len += 2
		if i%7 == 6 {
			len -= 6
		}
		fmt.Fprintf(f, "%s", strings.Repeat(" ", n-len/2+1))
		if i == 0 {
			fmt.Fprintln(f, "<span class=\"s\">*</span>")
			continue
		}
		fmt.Fprintf(f, "/")
		len2 := len
		if i%7 == 5 {
			len2 -= 4
			fmt.Fprintf(f, "^^")
		}
		for j := 0; j < len2-2; j++ {
			t := rand.Intn(5 * i)
			if t < i {
				fmt.Fprintf(f, "<span class=\"c%d\">•</span>", rand.Intn(3))
			} else if t < int(1.5*float64(i)) {
				fmt.Fprintf(f, "<span class=\"s\">*</span>")
			} else {
				if i == n-1 {
					fmt.Fprintf(f, "^")
				} else {
					fmt.Fprintf(f, " ")
				}
			}
		}
		if i%7 == 5 {
			fmt.Fprintf(f, "^^")
		}
		fmt.Fprintln(f, "\\")
	}
	fmt.Fprintf(f, "<span class=\"t\">")
	for i := 0; i < 5; i++ {
		fmt.Fprintf(f, "%s", strings.Repeat(" ", n))
		fmt.Fprintln(f, "===")
	}
	fmt.Fprintf(f, "</span>")
	fmt.Fprintln(f, "</pre></div></foreignObject></svg>")
	fmt.Printf("wrote tree.svg")
}

func genNext(n int) {
	os.Mkdir(fmt.Sprintf("day%02d", n), 0755)
	f, err := os.Create(fmt.Sprintf("day%02d/day%02d.go", n, n))
	utils.PanicOnErr(err)
	defer f.Close()
	f.WriteString(fmt.Sprintf(`package day%02d

func Part1(input string) int {
	return 0
}

func Part2(input string) int {
	return 0
}
`, n))
	fmt.Printf("wrote day%02d/day%02d.go\n", n, n)

	f, err = os.Create(fmt.Sprintf("day%02d/day%02d_test.go", n, n))
	utils.PanicOnErr(err)
	defer f.Close()
	f.WriteString(fmt.Sprintf(`package day%02d

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1("")
	require.Equal(t, 0, r)
}

func TestPart2(t *testing.T) {
	r := Part2("")
	require.Equal(t, 0, r)
}
`, n))
	fmt.Printf("wrote day%02d/day%02d_test.go\n", n, n)

}
