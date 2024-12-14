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
	Ax, Ay int64
	Bx, By int64
	Px, Py int64
}

func main() {
	machines, err := readMachines("input.txt")
	if err != nil {
		fmt.Println("Error reading machines:", err)
		return
	}

	// Apply the offset to prize coordinates
	offset := int64(10000000000000)
	for i := range machines {
		machines[i].Px += offset
		machines[i].Py += offset
	}

	totalCost := int64(0)
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

// solveMachine uses the direct algebraic method:
// Det = Ax*By - Ay*Bx
// a = (Px*By - Py*Bx)/Det
// b = (Py*Ax - Px*Ay)/Det
// Must have Det != 0, a and b integers and >=0
// Returns (cost, true) if possible, else (_, false).
func solveMachine(m Machine) (int64, bool) {
	Ax, Ay := m.Ax, m.Ay
	Bx, By := m.Bx, m.By
	Px, Py := m.Px, m.Py

	Det := Ax*By - Ay*Bx
	if Det == 0 {
		return 0, false // No unique solution
	}

	aNum := Px*By - Py*Bx
	bNum := Py*Ax - Px*Ay

	// Check divisibility
	if aNum%Det != 0 || bNum%Det != 0 {
		return 0, false
	}

	a := aNum / Det
	b := bNum / Det

	if a < 0 || b < 0 {
		return 0, false
	}

	// cost = 3a + b
	cost := 3*a + b
	return cost, true
}

// readMachines parses the input file expected in the same directory.
// Each machine is described by three lines and then a blank line:
//
// Button A: X+Ax, Y+Ay
// Button B: X+Bx, Y+By
// Prize: X=Px, Y=Py
func readMachines(filename string) ([]Machine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var machines []Machine
	scanner := bufio.NewScanner(file)

	// Regex to extract integers from lines
	reButton := regexp.MustCompile(`X([+\-]?\d+), Y([+\-]?\d+)`)
	rePrize := regexp.MustCompile(`X=(\d+), Y=(\d+)`)

	var lines []string
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			// blank line, process the accumulated lines
			if len(lines) == 3 {
				m, err := parseMachine(lines, reButton, rePrize)
				if err != nil {
					return nil, err
				}
				machines = append(machines, m)
			}
			lines = nil
		} else {
			lines = append(lines, text)
		}
	}
	// In case file doesn't end with a blank line
	if len(lines) == 3 {
		m, err := parseMachine(lines, reButton, rePrize)
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

func parseMachine(lines []string, reButton, rePrize *regexp.Regexp) (Machine, error) {
	// Lines:
	// Button A: X+Ax, Y+Ay
	// Button B: X+Bx, Y+By
	// Prize: X=Px, Y=Py
	m := Machine{}

	// Parse A line
	matchA := reButton.FindStringSubmatch(lines[0])
	if matchA == nil {
		return m, fmt.Errorf("cannot parse A line: %s", lines[0])
	}
	Ax, _ := strconv.ParseInt(matchA[1], 10, 64)
	Ay, _ := strconv.ParseInt(matchA[2], 10, 64)
	m.Ax, m.Ay = Ax, Ay

	// Parse B line
	matchB := reButton.FindStringSubmatch(lines[1])
	if matchB == nil {
		return m, fmt.Errorf("cannot parse B line: %s", lines[1])
	}
	Bx, _ := strconv.ParseInt(matchB[1], 10, 64)
	By, _ := strconv.ParseInt(matchB[2], 10, 64)
	m.Bx, m.By = Bx, By

	// Parse Prize line
	matchP := rePrize.FindStringSubmatch(lines[2])
	if matchP == nil {
		return m, fmt.Errorf("cannot parse Prize line: %s", lines[2])
	}
	Px, _ := strconv.ParseInt(matchP[1], 10, 64)
	Py, _ := strconv.ParseInt(matchP[2], 10, 64)
	m.Px, m.Py = Px, Py

	return m, nil
}
