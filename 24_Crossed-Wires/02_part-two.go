package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Input struct {
	Prefix string
	Gates  [][]string
}

// Parse the input into the prefix and gates sections
func parse(input string) Input {
	parts := strings.Split(input, "\n\n")
	prefix := parts[0]
	suffix := strings.Fields(parts[1])
	gates := [][]string{}
	for i := 0; i < len(suffix); i += 5 {
		gates = append(gates, suffix[i:i+5])
	}
	return Input{Prefix: prefix, Gates: gates}
}

// Convert a string to an index based on its character values
func toIndex(s string) int {
	b := []byte(s)
	return ((int(b[0]&31) << 10) + (int(b[1]&31) << 5) + int(b[2]&31))
}

// Part 1 solution
func part1(input Input) uint64 {
	prefix := input.Prefix
	gates := input.Gates

	todo := make([][]string, len(gates))
	copy(todo, gates)

	cache := make([]byte, 1<<15)
	for i := range cache {
		cache[i] = 255 // equivalent to u8::MAX
	}

	var result uint64 = 0

	// Parse the prefix lines to initialize cache
	for _, line := range strings.Split(prefix, "\n") {
		if line == "" {
			continue
		}
		prefix := line[:3]
		suffix, _ := strconv.Atoi(line[5:])
		cache[toIndex(prefix)] = byte(suffix)
	}

	// Process gates
	for len(todo) > 0 {
		gate := todo[0]
		todo = todo[1:]

		left := cache[toIndex(gate[0])]
		right := cache[toIndex(gate[2])]
		to := toIndex(gate[4])

		if left == 255 || right == 255 {
			// Requeue the gate if inputs are not ready
			todo = append(todo, gate)
		} else {
			switch gate[1] {
			case "AND":
				cache[to] = left & right
			case "OR":
				cache[to] = left | right
			case "XOR":
				cache[to] = left ^ right
			default:
				panic("unknown operation")
			}
		}
	}

	// Calculate the result based on z wires
	for i := toIndex("z00"); i < toIndex("z64"); i++ {
		if cache[i] != 255 {
			result = (result << 1) | uint64(cache[i])
		}
	}

	return result
}

// Part 2 solution
func part2(input Input) string {
	gates := input.Gates
	lookup := make(map[string]bool)
	swapped := make(map[string]bool)

	// Build a lookup table
	for _, gate := range gates {
		lookup[gate[0]+gate[1]] = true
		lookup[gate[2]+gate[1]] = true
	}

	for _, gate := range gates {
		left := gate[0]
		kind := gate[1]
		right := gate[2]
		to := gate[4]

		switch kind {
		case "AND":
			// Check conditions for AND gates
			if left != "x00" && right != "x00" && !lookup[to+"OR"] {
				swapped[to] = true
			}
		case "OR":
			// Check conditions for OR gates
			if strings.HasPrefix(to, "z") && to != "z45" {
				swapped[to] = true
			}
			if lookup[to+"OR"] {
				swapped[to] = true
			}
		case "XOR":
			if strings.HasPrefix(left, "x") || strings.HasPrefix(right, "x") {
				// Check conditions for first level XOR gates
				if left != "x00" && right != "x00" && !lookup[to+"XOR"] {
					swapped[to] = true
				}
			} else {
				// Check conditions for second level XOR gates
				if !strings.HasPrefix(to, "z") {
					swapped[to] = true
				}
			}
		}
	}

	// Collect swapped wires, sort, and return as a comma-separated string
	result := []string{}
	for wire := range swapped {
		result = append(result, wire)
	}
	sort.Strings(result)

	return strings.Join(result, ",")
}

// Read the input file
func readInputFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return builder.String(), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go input.txt")
		return
	}

	filename := os.Args[1]
	inputText, err := readInputFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Parse the input
	input := parse(inputText)

	// Solve part 1 and part 2
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
