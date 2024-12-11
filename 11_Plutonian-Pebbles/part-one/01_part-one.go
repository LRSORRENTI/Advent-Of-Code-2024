package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"plutonianpebbles/utils"
)

func countStonesAfterBlinks(stones []int, blinks int) int {
	queue := make([]int, len(stones))
	copy(queue, stones)

	for i := 0; i < blinks; i++ {
		nextQueue := []int{}
		for _, stone := range queue {
			if stone == 0 {
				nextQueue = append(nextQueue, 1)
			} else if len(strconv.Itoa(stone))%2 == 0 {
				left, right := utils.SplitNumber(stone)
				nextQueue = append(nextQueue, left, right)
			} else {
				nextQueue = append(nextQueue, stone*2024)
			}
		}
		queue = nextQueue
	}
	return len(queue)
}

func main() {
	// Read input from file
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	input := strings.TrimSpace(string(data))
	stoneStrs := strings.Fields(input)
	stones := make([]int, len(stoneStrs))
	for i, s := range stoneStrs {
		stones[i], _ = strconv.Atoi(s)
	}

	// Number of blinks for part two
	blinks := 75

	// Count stones after blinks
	count := countStonesAfterBlinks(stones, blinks)

	// Output the result
	fmt.Println("Number of stones after", blinks, "blinks:", count)
}
