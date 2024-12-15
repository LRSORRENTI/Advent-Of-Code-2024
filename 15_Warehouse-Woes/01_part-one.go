package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cellType int

const (
	Wall cellType = iota
	Empty
	Box
	Robot
)

func main() {
	// Open input file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var warehouse []string
	// Read until the warehouse map is fully parsed (in Advent of Code style puzzles,
	// you often know the dimensions or see a line of hyphens or a blank line signifying end)
	// For this puzzle, you should know your input format. Assuming we read until a certain separator line.
	// If you do not have a clear separator, you might know how many lines form the map.
	// Adjust according to your input format. For demonstration, let's assume the map ends at a line of '#' repeated:

	// NOTE: Adjust this logic based on your actual input format or known map size.
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// Possibly end of map or moves input starts after a blank line
			break
		}
		warehouse = append(warehouse, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Now warehouse is a slice of strings representing the grid.
	// Find the robot position, store boxes, etc.
	rows := len(warehouse)
	cols := len(warehouse[0])

	grid := make([][]cellType, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]cellType, cols)
	}

	var robotR, robotC int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			switch warehouse[r][c] {
			case '#':
				grid[r][c] = Wall
			case '.':
				grid[r][c] = Empty
			case 'O':
				grid[r][c] = Box
			case '@':
				grid[r][c] = Robot
				robotR, robotC = r, c
			}
		}
	}

	// Now read moves - the puzzle states moves come after the map, possibly multiline.
	// Let's say everything else in the file is moves. We'll concatenate them:
	movesBuilder := strings.Builder{}
	for scanner.Scan() {
		movesBuilder.WriteString(strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	moves := movesBuilder.String()

	// Direction vectors
	dirMap := map[rune][2]int{
		'^': {-1, 0},
		'v': {1, 0},
		'<': {0, -1},
		'>': {0, 1},
	}

	// Process moves
	for _, move := range moves {
		d := dirMap[move]
		newR := robotR + d[0]
		newC := robotC + d[1]

		// If next cell is wall, no move
		if grid[newR][newC] == Wall {
			continue
		}

		// If next cell is empty, move robot
		if grid[newR][newC] == Empty {
			// Move robot
			grid[robotR][robotC] = Empty
			grid[newR][newC] = Robot
			robotR, robotC = newR, newC
			continue
		}

		// If next cell is box, attempt to push
		if grid[newR][newC] == Box {
			// Follow the chain of boxes
			chainCells := [][2]int{{newR, newC}}
			cr, cc := newR, newC
			for {
				nr := cr + d[0]
				nc := cc + d[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					// Out of bounds - no move
					chainCells = nil
					break
				}
				if grid[nr][nc] == Box {
					// Another box in chain
					chainCells = append(chainCells, [2]int{nr, nc})
					cr, cc = nr, nc
					continue
				} else {
					// Reached a non-box cell
					// If wall, no move
					if grid[nr][nc] == Wall {
						chainCells = nil
					} else if grid[nr][nc] == Empty {
						// We can push into this cell
						// chainCells describes all boxes in a line
					}
					break
				}
			}

			if chainCells == nil {
				// Can't push
				continue
			}

			// The cell after the last box:
			lastBox := chainCells[len(chainCells)-1]
			targetR := lastBox[0] + d[0]
			targetC := lastBox[1] + d[1]

			if grid[targetR][targetC] == Empty {
				// We can push
				// Move all boxes forward
				// Start from the last box in the chain to avoid overwriting
				grid[robotR][robotC] = Empty
				for i := len(chainCells) - 1; i >= 0; i-- {
					br, bc := chainCells[i][0], chainCells[i][1]
					grid[br][bc] = Empty
					grid[br+d[0]][bc+d[1]] = Box
				}
				// Move robot
				grid[robotR][robotC] = Empty
				grid[newR][newC] = Robot
				robotR, robotC = newR, newC
			} else {
				// Can't push into that cell
				continue
			}
		}
	}

	// After all moves, compute sum of box GPS coordinates
	sum := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == Box {
				// GPS = r*100 + c (assuming top-left corner is (0,0))
				gps := r*100 + c
				sum += gps
			}
		}
	}

	fmt.Println(sum)
}
