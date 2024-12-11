package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input from file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Parse input
	stoneNumbers := parseInput(string(input))

	// Part 1: 25 blinks
	fmt.Println(part(stoneNumbers, 25))

	// Part 2: 75 blinks
	fmt.Println(part(stoneNumbers, 75))
}

func parseInput(input string) []uint64 {
	fields := strings.Fields(input)
	stones := make([]uint64, len(fields))
	for i, field := range fields {
		val, err := strconv.ParseUint(field, 10, 64)
		if err != nil {
			panic(err)
		}
		stones[i] = val
	}
	return stones
}

func part(input []uint64, blinks int) uint64 {
	return count(input, blinks)
}

func count(input []uint64, blinks int) uint64 {
	// Map stone numbers to indices
	indices := make(map[uint64]int)
	// Stones transformations and counts
	var stones []struct {
		first, second int
	}
	var current []uint64
	var todo []uint64

	// Initialize stones from input
	for _, number := range input {
		if index, exists := indices[number]; exists {
			current[index]++
		} else {
			indices[number] = len(current)
			todo = append(todo, number)
			current = append(current, 1)
		}
	}

	for i := 0; i < blinks; i++ {
		numbers := todo
		todo = make([]uint64, 0, 200)

		indexOf := func(number uint64) int {
			if index, exists := indices[number]; exists {
				return index
			}
			indices[number] = len(current)
			todo = append(todo, number)
			current = append(current, 0)
			return indices[number]
		}

		// Apply transformation logic to stones
		for _, number := range numbers {
			var first, second int
			if number == 0 {
				first, second = indexOf(1), -1
			} else {
				digits := int(math.Log10(float64(number))) + 1
				if digits%2 == 0 {
					power := uint64(math.Pow10(digits / 2))
					first, second = indexOf(number/power), indexOf(number%power)
				} else {
					first, second = indexOf(number*2024), -1
				}
			}
			stones = append(stones, struct {
				first, second int
			}{first, second})
		}

		// Aggregate counts for next blink
		next := make([]uint64, len(current))
		for i, stone := range stones {
			amount := current[i]
			next[stone.first] += amount
			if stone.second != -1 {
				next[stone.second] += amount
			}
		}
		current = next
	}

	// Sum up the counts
	var total uint64
	for _, amount := range current {
		total += amount
	}
	return total
}
