package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type robot struct {
	x, y   int64
	dx, dy int64
}

func main() {
	robots, err := readRobots("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// We know the pattern repeats every 101*103 = 10403 seconds
	width, height := int64(101), int64(103)
	period := int64(width * height) // 101*103=10403

	timeAtPattern := findEarliestUniqueTime(robots, width, height, period)
	fmt.Printf("The fewest number of seconds that must elapse: %d\n", timeAtPattern)
}

// findEarliestUniqueTime checks each t from 0 to period-1 to find when all robots
// occupy unique positions.
func findEarliestUniqueTime(robots []robot, width, height, period int64) int64 {
	// We'll create a boolean slice for visited positions each time.
	// Positions can be mapped as index = y*width + x.
	// This requires height*width=10403 booleans, which is not big.
	for t := int64(0); t < period; t++ {
		visited := make([]bool, width*height)
		unique := true
		for _, r := range robots {
			finalX := (r.x + r.dx*t) % width
			if finalX < 0 {
				finalX += width
			}
			finalY := (r.y + r.dy*t) % height
			if finalY < 0 {
				finalY += height
			}
			pos := finalY*width + finalX
			if visited[pos] {
				// Collision! Not all unique
				unique = false
				break
			}
			visited[pos] = true
		}
		if unique {
			// Found a time where all positions are unique
			return t
		}
	}

	// If not found, return -1 or handle the error.
	return -1
}

func readRobots(filename string) ([]robot, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Lines like: p=0,4 v=3,-3
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`)

	var robots []robot
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
