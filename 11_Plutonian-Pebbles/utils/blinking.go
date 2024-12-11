package utils

import (
	"strconv"
)

// Function to split a number into two parts
func SplitNumber(num int) (int, int) {
	numStr := strconv.Itoa(num)
	n := len(numStr) / 2
	left, _ := strconv.Atoi(numStr[:n])
	right, _ := strconv.Atoi(numStr[n:])
	return left, right
}

// Function to simulate the blinking process
func BlinkStones(stones []int, blinks int) []int {
	for i := 0; i < blinks; i++ {
		nextStones := []int{}
		for _, stone := range stones {
			if stone == 0 {
				nextStones = append(nextStones, 1)
			} else if len(strconv.Itoa(stone))%2 == 0 {
				left, right := SplitNumber(stone)
				nextStones = append(nextStones, left, right)
			} else {
				nextStones = append(nextStones, stone*2024)
			}
		}
		stones = nextStones
	}
	return stones
}
