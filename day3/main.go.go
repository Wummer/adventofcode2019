package day3

import (
	"math"
	"strconv"
	"strings"
)

func ClosestIntersection(wire1 [][2]int, wire2 [][2]int) int {
	var intersections [][2]int
	for _, cWire1 := range wire1 {
		for _, cWire2 := range wire2 {
			if cWire1 == cWire2 && cWire1 != [2]int{0, 0} {
				intersections = append(intersections, cWire1)
			}
		}
	}
	closestDistance := math.MaxInt64
	for _, intersection := range intersections {
		distance := ManhattanDistance(intersection)
		if distance < closestDistance {
			closestDistance = distance
		}
	}
	return closestDistance
}

func FindCoordinates(input string) [][2]int {
	var result [][2]int
	paths := strings.Split(input, ",")

	startCoordinate := [2]int{0, 0}
	coordinates := [][2]int{startCoordinate}

	for _, path := range paths {
		result = DrawPath(coordinates[len(coordinates)-1], path)

		coordinates = append(coordinates, result...)
	}
	return coordinates
}

func ManhattanDistance(y [2]int) int {
	return int(math.Abs(float64(y[0])) + math.Abs(float64(y[1])))
}

func DrawPath(start [2]int, input string) [][2]int {
	distance, _ := strconv.Atoi(input[1:])

	var x int
	var y int

	switch input[0] {
	case 'R':
		x = 1
	case 'L':
		x = -1
	case 'U':
		y = 1
	case 'D':
		y = -1
	}
	line := [][2]int{start}
	for i := 0; i < distance; i++ {
		start[0] += x
		start[1] += y
		line = append(line, start)
	}
	return line
}
