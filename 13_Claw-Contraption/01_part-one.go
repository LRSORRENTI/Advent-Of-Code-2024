package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	Ax, Ay int
	Bx, By int
	Px, Py int
}

func main() {
	machines, err := readMachines("input.txt")
	if err != nil {
		fmt.Println("Error reading machines:", err)
		return
	}

	totalCost := 0
	winCount := 0
	for _, m := range machines {
		cost, ok := solveMachine(m)
		if ok {
			winCount++
			totalCost += cost
		}
	}

	fmt.Printf("Max prizes won: %d\n", winCount)
	fmt.Printf("Minimum total tokens: %d\n", totalCost)
}

// solveMachine tries to find a and b (0 <= a,b <= 100) that satisfy:
// a*Ax + b*Bx = Px and a*Ay + b*By = Py
// Minimizing cost = 3*a + b
func solveMachine(m Machine) (int, bool) {
	minCost := -1
	for a := 0; a <= 100; a++ {
		dx := m.Px - a*m.Ax
		if m.Bx == 0 {
			continue
		}
		if dx%m.Bx != 0 {
			continue
		}
		bx := dx / m.Bx

		dy := m.Py - a*m.Ay
		if m.By == 0 {
			continue
		}
		if dy%m.By != 0 {
			continue
		}
		by := dy / m.By

		if bx != by {
			continue
		}
		b := bx

		if b < 0 || b > 100 {
			continue
		}

		cost := 3*a + b
		if minCost == -1 || cost < minCost {
			minCost = cost
		}
	}

	return minCost, minCost != -1
}

// readMachines reads the input file and parses the machines' configurations.
func readMachines(filename string) ([]Machine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var machines []Machine
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		// Collect lines until we get 3 lines per machine
		if strings.TrimSpace(line) == "" {
			// Blank line indicates one machine block ended
			if len(lines) == 3 {
				m, err := parseMachine(lines)
				if err != nil {
					return nil, err
				}
				machines = append(machines, m)
			}
			lines = []string{}
		} else {
			lines = append(lines, line)
			if len(lines) == 3 {
				// If there's no blank line after the third line, we still
				// consider this one machine.
				// But we must be prepared for the possibility that the file
				// might not have a trailing blank line at the end.
			}
		}
	}

	// Handle the last machine if the file does not end with a blank line
	if len(lines) == 3 {
		m, err := parseMachine(lines)
		if err != nil {
			return nil, err
		}
		machines = append(machines, m)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return machines, nil
}

// parseMachine takes three lines describing one machine and extracts Ax, Ay, Bx, By, Px, Py.
func parseMachine(lines []string) (Machine, error) {
	// Expected format:
	// Button A: X+45, Y+76
	// Button B: X+84, Y+14
	// Prize: X=9612, Y=4342

	var m Machine
	var err error

	// Regular expressions for parsing
	// For Button lines: "Button A: X+45, Y+76"
	buttonLineRegexp := regexp.MustCompile(`Button [AB]: X([+\-]?\d+), Y([+\-]?\d+)`)
	// For Prize line: "Prize: X=9612, Y=4342"
	prizeLineRegexp := regexp.MustCompile(`Prize: X=([+\-]?\d+), Y=([+\-]?\d+)`)

	// Line 1: Button A
	aMatches := buttonLineRegexp.FindStringSubmatch(lines[0])
	if len(aMatches) != 3 {
		return m, fmt.Errorf("line '%s' does not match expected A button format", lines[0])
	}
	m.Ax, err = strconv.Atoi(aMatches[1])
	if err != nil {
		return m, err
	}
	m.Ay, err = strconv.Atoi(aMatches[2])
	if err != nil {
		return m, err
	}

	// Line 2: Button B
	bMatches := buttonLineRegexp.FindStringSubmatch(lines[1])
	if len(bMatches) != 3 {
		return m, fmt.Errorf("line '%s' does not match expected B button format", lines[1])
	}
	m.Bx, err = strconv.Atoi(bMatches[1])
	if err != nil {
		return m, err
	}
	m.By, err = strconv.Atoi(bMatches[2])
	if err != nil {
		return m, err
	}

	// Line 3: Prize
	pMatches := prizeLineRegexp.FindStringSubmatch(lines[2])
	if len(pMatches) != 3 {
		return m, fmt.Errorf("line '%s' does not match expected Prize format", lines[2])
	}
	m.Px, err = strconv.Atoi(pMatches[1])
	if err != nil {
		return m, err
	}
	m.Py, err = strconv.Atoi(pMatches[2])
	if err != nil {
		return m, err
	}

	return m, nil
}
