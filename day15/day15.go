package day15

import (
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2022/utils"
)

type sensor struct {
	sensorX  int
	sensorY  int
	beaconX  int
	beaconY  int
	distance int
}

func Part1(input string, row int) int {
	disallowed := map[int]bool{}
	re := regexp.MustCompile(`Sensor at x=(-?[0-9]+), y=(-?[0-9]+): closest beacon is at x=(-?[0-9]+), y=(-?[0-9]+)`)
	for _, line := range strings.Split(input, "\n") {
		pieces := re.FindStringSubmatch(line)
		process(disallowed, row,
			utils.MustAtoi(pieces[1]),
			utils.MustAtoi(pieces[2]),
			utils.MustAtoi(pieces[3]),
			utils.MustAtoi(pieces[4]))
	}
	return len(disallowed)
}

// The key insight for part2 is that the only possible position for the distress beacon is
// going to be distance+1 to one of the sensors. So we can simply check the border at distance+1
// for each sensor and remove the parts which are included in other sensors' disallow area.
func Part2(input string, rowcol int) int {
	sensors := []sensor{}
	re := regexp.MustCompile(`Sensor at x=(-?[0-9]+), y=(-?[0-9]+): closest beacon is at x=(-?[0-9]+), y=(-?[0-9]+)`)
	for _, line := range strings.Split(input, "\n") {
		pieces := re.FindStringSubmatch(line)
		s := sensor{}
		s.sensorX = utils.MustAtoi(pieces[1])
		s.sensorY = utils.MustAtoi(pieces[2])
		s.beaconX = utils.MustAtoi(pieces[3])
		s.beaconY = utils.MustAtoi(pieces[4])
		s.distance = utils.Abs(s.sensorX-s.beaconX) + utils.Abs(s.sensorY-s.beaconY)
		sensors = append(sensors, s)
	}

	x, y := findXY(sensors, rowcol)
	return x*4000000 + y
}

func process(disallowed map[int]bool, row, sensorX, sensorY, beaconX, beaconY int) {
	distance := utils.Abs(sensorX-beaconX) + utils.Abs(sensorY-beaconY)
	leftOver := distance - utils.Abs(row-sensorY)
	if leftOver < 0 {
		return
	}
	disallowed[sensorX] = true
	for i := 1; i <= leftOver; i++ {
		disallowed[sensorX+i] = true
		disallowed[sensorX-i] = true
	}
	if beaconY == row {
		delete(disallowed, beaconX)
	}
}

func findXY(sensors []sensor, rowcol int) (int, int) {
	for _, sensor := range sensors {
		for row := sensor.sensorY - (sensor.distance + 1); row <= sensor.sensorY+(sensor.distance+1); row++ {
			if row < 0 || row > rowcol {
				continue
			}
			leftOver := (sensor.distance + 1) - utils.Abs(row-sensor.sensorY)
			col := sensor.sensorX - leftOver
			if col >= 0 && col <= rowcol {
				if !included(sensors, col, row) {
					return col, row
				}
			}
			col = sensor.sensorX + leftOver
			if col >= 0 && col <= rowcol {
				if !included(sensors, col, row) {
					return col, row
				}
			}
		}
	}
	panic("not found")
}

func included(sensors []sensor, x, y int) bool {
	for _, sensor := range sensors {
		distance := utils.Abs(sensor.sensorX-x) + utils.Abs(sensor.sensorY-y)
		if distance <= sensor.distance {
			return true
		}
	}
	return false
}
