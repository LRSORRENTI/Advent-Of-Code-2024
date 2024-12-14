package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	robots, err := readRobots("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Dimensions
	width := int64(101)
	height := int64(103)
	seconds := int64(100)

	// Count of robots in each quadrant
	Q1, Q2, Q3, Q4 := 0, 0, 0, 0

	for _, r := range robots {
		finalX := (r.x + r.dx*seconds) % width
		finalY := (r.y + r.dy*seconds) % height

		// Handle negative remainders
		if finalX < 0 {
			finalX += width
		}
		if finalY < 0 {
			finalY += height
		}

		// Determine quadrant
		// Center lines: x=50, y=51
		if finalX < 50 && finalY < 51 {
			Q1++
		} else if finalX > 50 && finalY < 51 {
			Q2++
		} else if finalX < 50 && finalY > 51 {
			Q3++
		} else if finalX > 50 && finalY > 51 {
			Q4++
		}
		// If finalX == 50 or finalY == 51, robot is on center line, not in any quadrant.
	}

	safetyFactor := Q1 * Q2 * Q3 * Q4

	fmt.Printf("Quadrants: Q1=%d Q2=%d Q3=%d Q4=%d\n", Q1, Q2, Q3, Q4)
	fmt.Printf("Safety factor: %d\n", safetyFactor)
}

type robot struct {
	x, y   int64
	dx, dy int64
}

func readRobots(filename string) ([]robot, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var robots []robot

	// Example line: p=0,4 v=3,-3
	// Regex or simple parsing can be used
	// Let's use a regex approach:
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		match := re.FindStringSubmatch(line)
		if match == nil {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		x, _ := strconv.ParseInt(match[1], 10, 64)
		y, _ := strconv.ParseInt(match[2], 10, 64)
		dx, _ := strconv.ParseInt(match[3], 10, 64)
		dy, _ := strconv.ParseInt(match[4], 10, 64)

		robots = append(robots, robot{x: x, y: y, dx: dx, dy: dy})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return robots, nil
}
