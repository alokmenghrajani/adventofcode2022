package main

import (
	"fmt"
	"os"

	"github.com/alokmenghrajani/adventofcode2022/day01"
	"github.com/alokmenghrajani/adventofcode2022/day02"
	"github.com/alokmenghrajani/adventofcode2022/day03"
	"github.com/alokmenghrajani/adventofcode2022/day04"
	"github.com/alokmenghrajani/adventofcode2022/utils"
)

// Usage: go run main.go <NN>
// assumes input is in day<NN>/input.txt
func main() {
	d := day()
	fmt.Printf("Running day %02d\n", d)

	switch d {
	case 1:
		fmt.Printf("part 1: %d\n", day01.Part1(utils.Readfile(2022, d)))
		fmt.Printf("part 2: %d\n", day01.Part2(utils.Readfile(2022, d)))
	case 2:
		fmt.Printf("part 1: %d\n", day02.Part1(utils.Readfile(2022, d)))
		fmt.Printf("part 2: %d\n", day02.Part2(utils.Readfile(2022, d)))
	case 3:
		fmt.Printf("part 1: %d\n", day03.Part1(utils.Readfile(2022, d)))
		fmt.Printf("part 2: %d\n", day03.Part2(utils.Readfile(2022, d)))
	case 4:
		fmt.Printf("part 1: %d\n", day04.Part1(utils.Readfile(2022, d)))
		fmt.Printf("part 2: %d\n", day04.Part2(utils.Readfile(2022, d)))
	// case 5:
	// 	fmt.Printf("part 1: %d\n", day05.Part1(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day05.Part2(utils.Readfile(2022, d)))
	// case 6:
	// 	fmt.Printf("part 1: %d\n", day06.Part(utils.Readfile(2022, d), 80))
	// 	fmt.Printf("part 2: %d\n", day06.Part(utils.Readfile(2022, d), 256))
	// case 7:
	// 	_, fuel := day07.Part1(utils.Readfile(2022, d))
	// 	fmt.Printf("part 1: %d\n", fuel)
	// 	_, fuel = day07.Part2(utils.Readfile(2022, d))
	// 	fmt.Printf("part 2: %d\n", fuel)
	// case 8:
	// 	fmt.Println("Using z3")
	// 	fmt.Printf("part 1: %d\n", day08.Part1WithZ3(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day08.Part2WithZ3(utils.Readfile(2022, d)))
	// 	fmt.Println("Using gophersat")
	// 	fmt.Printf("part 1: %d\n", day08.Part1WithGophersat(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day08.Part2WithGophersat(utils.Readfile(2022, d)))
	// case 9:
	// 	fmt.Printf("part 1: %d\n", day09.Part1(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day09.Part2(utils.Readfile(2022, d)))
	// case 10:
	// 	fmt.Printf("part 1: %d\n", day10.Part1(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day10.Part2(utils.Readfile(2022, d)))
	// case 11:
	// 	fmt.Printf("part 1: %d\n", day11.Part1(utils.Readfile(2022, d), 100))
	// 	fmt.Printf("part 2: %d\n", day11.Part2(utils.Readfile(2022, d)))
	// case 12:
	// 	fmt.Printf("part 1: %d\n", day12.Part1(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day12.Part2(utils.Readfile(2022, d)))
	// case 13:
	// 	fmt.Printf("part 1: %d\n", day13.Part1(utils.Readfile(2022, d)))
	// 	fmt.Println("part 2")
	// 	day13.Part2(utils.Readfile(2022, d))
	// case 14:
	// 	fmt.Printf("part 1: %d\n", day14.Part(utils.Readfile(2022, d), 10))
	// 	fmt.Printf("part 2: %d\n", day14.Part(utils.Readfile(2022, d), 40))
	// case 15:
	// 	fmt.Printf("part 1: %d\n", day15.Part(utils.Readfile(2022, d), 1))
	// 	fmt.Printf("part 2: %d\n", day15.Part(utils.Readfile(2022, d), 5))
	// case 16:
	// 	fmt.Printf("part 1: %d\n", day16.Part1(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day16.Part2(utils.Readfile(2022, d)))
	// case 17:
	// 	fmt.Printf("part 1: %d\n", day17.Part1(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day17.Part2(utils.Readfile(2022, d)))
	// case 18:
	// 	fmt.Printf("part 1: %d\n", day18.Part1(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day18.Part2(utils.Readfile(2022, d)))
	// case 19:
	// 	fmt.Printf("part 1: %d\n", day19.Part1(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day19.Part2(utils.Readfile(2022, d)))
	// case 20:
	// 	fmt.Printf("part 1: %d\n", day20.Part(utils.Readfile(2022, d), 2))
	// 	fmt.Printf("part 1: %d\n", day20.Part(utils.Readfile(2022, d), 50))
	// case 21:
	// 	fmt.Printf("part 1: %d\n", day21.Part1(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day21.Part2(utils.Readfile(2022, d)))
	// case 22:
	// 	fmt.Printf("part 1: %d\n", day22.Part1(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day22.Part2(utils.Readfile(2022, d)))
	// case 23:
	// 	fmt.Printf("part 1: %d\n", day23.Run1(utils.Readfile(2022, d)))
	// case 24:
	// 	fmt.Printf("part 1: %d\n", day24.Part1(utils.Readfile(2022, d)))
	// 	fmt.Printf("part 2: %d\n", day24.Part2(utils.Readfile(2022, d)))
	// case 25:
	// 	fmt.Printf("part 1: %d\n", day25.Part1(utils.Readfile(2022, d)))

	default:
		panic(fmt.Errorf("no such day: %d", d))
	}
}

// Reads day from os.Args.
func day() int {
	if len(os.Args) == 1 {
		return 4
	}
	day := utils.Atoi(os.Args[1], -1)
	return day
}
