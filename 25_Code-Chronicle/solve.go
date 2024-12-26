package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// parseHeightsLock converts 7 lines of a LOCK schematic into pin heights.
func parseHeightsLock(lines []string) []int {
	// We assume all lines are the same width.
	width := len(lines[0])
	heights := make([]int, width)

	// For locks: skip row 0 (all '#') and row 6 (all '.'),
	// count consecutive '#' from row 1 down to row 5.
	for col := 0; col < width; col++ {
		count := 0
		for row := 1; row <= 5; row++ {
			if lines[row][col] == '#' {
				count++
			} else {
				break
			}
		}
		heights[col] = count
	}

	return heights
}

// parseHeightsKey converts 7 lines of a KEY schematic into key heights.
func parseHeightsKey(lines []string) []int {
	width := len(lines[0])
	heights := make([]int, width)

	// For keys: skip row 0 (all '.') and row 6 (all '#'),
	// count consecutive '#' from row 5 up to row 1.
	for col := 0; col < width; col++ {
		count := 0
		for row := 5; row >= 1; row-- {
			if lines[row][col] == '#' {
				count++
			} else {
				break
			}
		}
		heights[col] = count
	}

	return heights
}

// isLock checks if the 7-line schematic matches the pattern for a lock:
// top line is all '#' and bottom line is all '.'.
func isLock(lines []string) bool {
	if len(lines) != 7 {
		return false
	}
	top := lines[0]
	bottom := lines[6]
	return all(top, '#') && all(bottom, '.')
}

// isKey checks if the 7-line schematic matches the pattern for a key:
// top line is all '.' and bottom line is all '#'.
func isKey(lines []string) bool {
	if len(lines) != 7 {
		return false
	}
	top := lines[0]
	bottom := lines[6]
	return all(top, '.') && all(bottom, '#')
}

// all checks if an entire string is composed of a given rune.
func all(s string, ch rune) bool {
	for _, r := range s {
		if r != ch {
			return false
		}
	}
	return true
}

// fits returns true if, for every column, lockHeight + keyHeight <= 5.
func fits(lockHeights, keyHeights []int) bool {
	if len(lockHeights) != len(keyHeights) {
		// If the columns differ, you might decide they never fit,
		// or handle partial overlap. Adjust if needed.
		return false
	}
	for i := 0; i < len(lockHeights); i++ {
		if lockHeights[i]+keyHeights[i] > 5 {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	var locks [][]int
	var keys [][]int

	// We'll parse 7-line blocks, ignoring blank lines in between.
	for scanner.Scan() {
		line := scanner.Text()
		// Skip completely blank lines
		if strings.TrimSpace(line) == "" {
			continue
		}
		lines = append(lines, line)

		// Once we have 7 lines, we determine if it's a lock or a key,
		// parse, and reset lines for the next block.
		if len(lines) == 7 {
			if isLock(lines) {
				lockHeights := parseHeightsLock(lines)
				locks = append(locks, lockHeights)
			} else if isKey(lines) {
				keyHeights := parseHeightsKey(lines)
				keys = append(keys, keyHeights)
			} else {
				log.Println("Warning: 7-line block does not match lock or key pattern.")
			}
			lines = nil
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Now we have all locks and keys parsed into integer slices.
	// Count how many unique lock/key pairs fit together.
	count := 0
	for _, lockHeights := range locks {
		for _, keyHeights := range keys {
			if fits(lockHeights, keyHeights) {
				count++
			}
		}
	}

	fmt.Println("Number of unique lock/key pairs that fit:", count)
}
