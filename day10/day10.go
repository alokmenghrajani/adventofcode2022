package day10

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type cpu struct {
	x      int
	cycles int
	r      int
}

type crt struct {
	gunX   int
	gunY   int
	screen [][]bool
}

func Part1(input string) int {
	cpu := &cpu{x: 1, cycles: 0, r: 0}
	for _, line := range strings.Split(input, "\n") {
		if line == "noop" {
			cpu.nextCycle()
		} else {
			cpu.nextCycle()
			cpu.nextCycle()
			t := utils.MustAtoi(line[5:])
			cpu.x += t
		}
	}
	return cpu.r
}

func Part2(input string) string {
	cpu := &cpu{x: 1, cycles: 0, r: 0}
	crt := &crt{gunX: 0, gunY: 0, screen: make([][]bool, 6)}
	for i := 0; i < 6; i++ {
		crt.screen[i] = make([]bool, 40)
	}

	for _, line := range strings.Split(input, "\n") {
		if line == "noop" {
			crt.nextCycle(cpu)
			cpu.nextCycle()
		} else {
			crt.nextCycle(cpu)
			cpu.nextCycle()
			crt.nextCycle(cpu)
			cpu.nextCycle()
			t := utils.MustAtoi(line[5:])
			cpu.x += t
		}
	}
	return crt.String()
}

func (cpu *cpu) nextCycle() {
	cpu.cycles++
	if cpu.cycles == 20 || cpu.cycles == 60 || cpu.cycles == 100 || cpu.cycles == 140 || cpu.cycles == 180 || cpu.cycles == 220 {
		cpu.r += cpu.x * cpu.cycles
	}
}

func (crt *crt) nextCycle(cpu *cpu) {
	if utils.Abs(cpu.x-crt.gunX) <= 1 {
		crt.screen[crt.gunY][crt.gunX] = true
	} else {
		crt.screen[crt.gunY][crt.gunX] = false
	}
	crt.gunX++
	if crt.gunX == 40 {
		crt.gunY++
		crt.gunX = 0
		if crt.gunY == 6 {
			crt.gunY = 0
		}
	}
}

func (crt *crt) String() string {
	var r strings.Builder
	for j := 0; j < 6; j++ {
		for i := 0; i < 40; i++ {
			if crt.screen[j][i] {
				r.WriteByte('#')
			} else {
				r.WriteByte('.')
			}
		}
		r.WriteByte('\n')
	}
	return r.String()
}
