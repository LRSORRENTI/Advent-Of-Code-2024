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

var ORTHOGONAL = []Point{
	{0, -1}, // Up
	{-1, 0}, // Left
	{1, 0},  // Right
	{0, 1},  // Down
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

	// seen grid to mark visited plots
	seen := make([][]int, height)
	for i := range seen {
		seen[i] = make([]int, width)
	}

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
			perimeter := 0

			// Flood fill to find all plots in this region
			for area < len(region) {
				p := region[area]
				area++

				// Check orth neighbors
				for _, o := range ORTHOGONAL {
					next := p.Add(o)
					if g.Contains(next) && g.At(next) == kind {
						if seen[next.Y][next.X] == 0 {
							seen[next.Y][next.X] = id
							region = append(region, next)
						}
					} else {
						// Out of bounds or different kind -> increment perimeter
						perimeter++
					}
				}
			}

			// Price for this region
			regionPrice := area * perimeter
			totalPrice += regionPrice
		}
	}

	return totalPrice
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)

	// Calculate total price for all regions
	result := parse(input)
	fmt.Println(result)
}
