package main

import (
	"fmt"
	"os"

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
	default:
		panic(fmt.Errorf("no such day: %d", d))
	}
}

// Reads day from os.Args.
func day() int {
	if len(os.Args) == 1 {
		return 12
	}
	day := utils.Atoi(os.Args[1], -1)
	return day
}

// func genTree() {
// 	fmt.Println("<pre>")
// 	n := 30
// 	for i := 0; i < n; i++ {
// 		fmt.Printf("%s", strings.Repeat(" ", n-i))
// 		if i == 0 {
// 			fmt.Println("<span class=\"s\">*</span>")
// 			continue
// 		}
// 		fmt.Printf("&gt;")
// 		for j := 0; j < i*2-1; j++ {
// 			t := rand.Intn(i * 2)
// 			if t < 5 {
// 				if rand.Intn(2) == 0 {
// 					fmt.Printf("<span class=\"c%d\">o</span>", rand.Intn(3))
// 				} else {
// 					fmt.Printf("<span class=\"c%d\">O</span>", rand.Intn(3))
// 				}
// 			} else if t < 7 {
// 				fmt.Printf("<span class=\"s\">*</span>")
// 			} else {
// 				if rand.Intn(2) == 0 {
// 					fmt.Printf("&lt;")
// 				} else {
// 					fmt.Printf("&gt;")
// 				}
// 			}
// 		}
// 		fmt.Println("&lt;")
// 	}
// 	fmt.Printf("<span class=\"t\">")
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("%s", strings.Repeat(" ", n-1))
// 		fmt.Println("===")
// 	}
// 	fmt.Printf("</span>")
// 	fmt.Println("</pre>")
// }
