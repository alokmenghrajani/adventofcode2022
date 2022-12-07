package day07

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type dir struct {
	name   string
	files  map[string]int
	subdir map[string]*dir
	parent *dir
	size   int
}

func Part1(input string) int {
	root := parse(input)
	sum := 0
	root.recurse(&sum)
	return sum
}

func Part2(input string) int {
	root := parse(input)
	sum := 0
	root.recurse(&sum)

	smallest := utils.MaxInt
	spaceNeeded := 30000000 - (70000000 - root.size)
	root.recurseAgain(spaceNeeded, &smallest)
	return smallest
}

func parse(input string) *dir {
	root := &dir{name: "/", files: map[string]int{}, subdir: map[string]*dir{}, parent: nil}
	c := root
	for _, line := range strings.Split(input, "\n") {
		if c == nil {
			panic("oops")
		}

		pieces := strings.Split(line, " ")
		if pieces[0] == "$" {
			if pieces[1] == "cd" {
				if pieces[2] == ".." {
					c = c.parent
				} else if pieces[2] == "/" {
					c = root
				} else {
					c = c.addDirectoryIfMissing(pieces[2])
				}
			} else if pieces[1] == "ls" {
				// no need to do anything
			} else {
				panic("oops")
			}
		} else if pieces[0] == "dir" {
			c.addDirectoryIfMissing(pieces[1])
		} else {
			c.addFileIfMissing(pieces[1], utils.MustAtoi(pieces[0]))
		}
	}
	return root
}

func (d *dir) addDirectoryIfMissing(name string) *dir {
	t, ok := d.subdir[name]
	if !ok {
		t = &dir{
			name:   name,
			files:  map[string]int{},
			subdir: map[string]*dir{},
			parent: d,
		}
		d.subdir[name] = t
	}
	return t
}

func (d *dir) addFileIfMissing(name string, size int) {
	t, ok := d.files[name]
	if ok {
		if t != size {
			panic("oops")
		}
	} else {
		d.files[name] = size
	}
}

func (d *dir) recurse(sum *int) int {
	if d.size != 0 {
		panic("oops")
	}

	// add all the files
	for _, v := range d.files {
		d.size += v
	}

	// recurse into subdir
	for _, subdir := range d.subdir {
		d.size += subdir.recurse(sum)
	}

	// contribute to sum
	if d.size <= 100000 {
		*sum += d.size
	}

	return d.size
}

func (d *dir) recurseAgain(target int, smallest *int) {
	if (d.size < *smallest) && (d.size >= target) {
		*smallest = d.size
	}
	for _, subdir := range d.subdir {
		subdir.recurseAgain(target, smallest)
	}
}
