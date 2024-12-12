package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func (p Point) Add(o Point) Point {
	return Point{p.X + o.X, p.Y + o.Y}
}

type Grid struct {
	Width, Height int
	Data          []rune
}

func (g *Grid) Contains(p Point) bool {
	return p.X >= 0 && p.X < g.Width && p.Y >= 0 && p.Y < g.Height
}

func (g *Grid) At(p Point) rune {
	return g.Data[p.Y*g.Width+p.X]
}

// ORTHOGONAL neighbors for flood filling.
var ORTHOGONAL = []Point{
	{0, -1}, // Up
	{-1, 0}, // Left
	{1, 0},  // Right
	{0, 1},  // Down
}

// DIAGONAL neighbors in order: up_left, up, up_right, left, right, down_left, down, down_right
// This order must match the indexing used when creating the lookup table for sides.
var DIAGONAL = []Point{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func parse(input string) int {
	// Split into lines
	linesRaw := strings.Split(input, "\n")

	// Trim trailing spaces and ignore empty lines
	var lines []string
	for _, line := range linesRaw {
		line = strings.TrimRight(line, " \r")
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	if len(lines) == 0 {
		panic("No input lines found after trimming.")
	}

	height := len(lines)
	width := len(lines[0])
	for i, l := range lines {
		if len(l) != width {
			panic(fmt.Sprintf("Input lines are not of consistent width: line %d has length %d, expected %d", i, len(l), width))
		}
	}

	// Build the grid
	g := Grid{
		Width:  width,
		Height: height,
		Data:   make([]rune, width*height),
	}
	for y, line := range lines {
		for x, ch := range line {
			g.Data[y*width+x] = ch
		}
	}

	// Create a 2D slice to track which region a cell belongs to
	seen := make([][]int, height)
	for i := range seen {
		seen[i] = make([]int, width)
	}

	lut := sidesLUT()

	var region []Point
	totalPrice := 0

	// Flood fill each region
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if seen[y][x] != 0 {
				continue
			}

			point := Point{x, y}
			kind := g.At(point)
			id := y*g.Width + x + 1

			region = region[:0]
			region = append(region, point)
			seen[y][x] = id

			area := 0

			// Flood fill to find all plots in this region
			for area < len(region) {
				p := region[area]
				area++

				// Check orth neighbors for region continuity
				for _, o := range ORTHOGONAL {
					next := p.Add(o)
					if g.Contains(next) && g.At(next) == kind {
						if seen[next.Y][next.X] == 0 {
							seen[next.Y][next.X] = id
							region = append(region, next)
						}
					}
				}
			}

			// Now we need to calculate the number of sides using the lookup table
			sides := 0
			for _, p := range region {
				index := 0
				for _, d := range DIAGONAL {
					n := p.Add(d)
					index <<= 1
					if g.Contains(n) && seen[n.Y][n.X] == id {
						index |= 1
					}
				}
				sides += lut[index]
			}

			// Price for this region = area * sides
			regionPrice := area * sides
			totalPrice += regionPrice
		}
	}

	return totalPrice
}

// sidesLUT precomputes the number of "sides" based on an 8-bit pattern of neighbors.
// This is adapted directly from the original Rust logic you provided.
func sidesLUT() [256]int {
	var lut [256]int
	for neighbours := 0; neighbours < 256; neighbours++ {
		upLeft := (neighbours & (1 << 0)) != 0
		up := (neighbours & (1 << 1)) != 0
		upRight := (neighbours & (1 << 2)) != 0
		left := (neighbours & (1 << 3)) != 0
		right := (neighbours & (1 << 4)) != 0
		downLeft := (neighbours & (1 << 5)) != 0
		down := (neighbours & (1 << 6)) != 0
		downRight := (neighbours & (1 << 7)) != 0

		// Logic from the original code snippet:
		// A corner (side) occurs when there's a "bend" or an outer edge based on the pattern of neighbors.
		ul := (!up && !left) || (up && left && !upLeft)
		ur := (!up && !right) || (up && right && !upRight)
		dl := (!down && !left) || (down && left && !downLeft)
		dr := (!down && !right) || (down && right && !downRight)

		count := 0
		if ul {
			count++
		}
		if ur {
			count++
		}
		if dl {
			count++
		}
		if dr {
			count++
		}
		lut[neighbours] = count
	}
	return lut
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)

	// Calculate total price for all regions under the bulk discount (part two)
	result := parse(input)
	fmt.Println(result)
}
