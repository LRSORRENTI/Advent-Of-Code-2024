package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func (p Point) Add(o Point) Point {
	return Point{p.X + o.X, p.Y + o.Y}
}

var (
	UP    = Point{0, -1}
	DOWN  = Point{0, 1}
	LEFT  = Point{-1, 0}
	RIGHT = Point{1, 0}
)

// Grid holds a 2D array of bytes.
type Grid struct {
	Width, Height int
	Data          []byte
}

func NewGrid(width, height int, fill byte) Grid {
	g := Grid{
		Width:  width,
		Height: height,
		Data:   make([]byte, width*height),
	}
	for i := range g.Data {
		g.Data[i] = fill
	}
	return g
}

func (g Grid) Index(p Point) int {
	return p.Y*g.Width + p.X
}

func (g Grid) InBounds(p Point) bool {
	if p.X < 0 || p.X >= g.Width {
		return false
	}
	if p.Y < 0 || p.Y >= g.Height {
		return false
	}
	return true
}

func (g Grid) Get(p Point) byte {
	return g.Data[g.Index(p)]
}

func (g *Grid) Set(p Point, val byte) {
	g.Data[g.Index(p)] = val
}

func (g Grid) Find(val byte) (Point, bool) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.Data[y*g.Width+x] == val {
				return Point{x, y}, true
			}
		}
	}
	return Point{}, false
}

func (g Grid) Clone() Grid {
	copyData := make([]byte, len(g.Data))
	copy(copyData, g.Data)
	return Grid{
		Width:  g.Width,
		Height: g.Height,
		Data:   copyData,
	}
}

func (g Grid) SameSizeWith(initVal int) [][]int {
	seen := make([][]int, g.Height)
	for i := 0; i < g.Height; i++ {
		row := make([]int, g.Width)
		for j := 0; j < g.Width; j++ {
			row[j] = initVal
		}
		seen[i] = row
	}
	return seen
}

// Parse the input: first part is the map, blank line, then moves
func parseInput() (Grid, string) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	// Read until blank line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	// lines now contain the map
	grid := NewGrid(len(lines[0]), len(lines), 0)
	for y, l := range lines {
		copy(grid.Data[y*grid.Width:(y+1)*grid.Width], []byte(l))
	}

	// Now read moves
	movesBuilder := strings.Builder{}
	for scanner.Scan() {
		movesBuilder.WriteString(strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	moves := movesBuilder.String()

	return grid, moves
}

// Narrow move logic (part1 and also used in part2 for horizontal moves)
func narrow(grid *Grid, start *Point, direction Point) {
	position := start.Add(direction)
	size := 2

	// Search next open space or wall
	for grid.Get(position) != '.' && grid.Get(position) != '#' {
		position = position.Add(direction)
		size++
	}

	// If open space found, push items one space in direction
	if grid.Get(position) == '.' {
		prev := byte('.')
		curPos := *start

		for i := 0; i < size; i++ {
			idx := grid.Index(curPos)
			tmp := grid.Data[idx]
			grid.Data[idx] = prev
			prev = tmp
			curPos = curPos.Add(direction)
		}

		// Move robot
		*start = start.Add(direction)
	}
}

// Wide move logic for vertical moves in part2
func wide(grid *Grid, start *Point, direction Point, todo *[]Point, seen [][]int, id int) {
	position := *start
	next := position.Add(direction)

	// If next is empty, just move robot
	if grid.Get(next) == '.' {
		grid.Set(position, '.')
		grid.Set(next, '@')
		*start = next
		return
	}

	// BFS to find all boxes to move
	// Clear old search data
	*todo = (*todo)[:0]
	(*todo) = append(*todo, *start)

	for index := 0; index < len(*todo); index++ {
		cur := (*todo)[index]
		n := cur.Add(direction)
		val := grid.Get(n)

		if val == '#' {
			// Wall in the way, cancel move
			return
		}

		var otherDir Point
		if val == '[' {
			otherDir = RIGHT
		} else if val == ']' {
			otherDir = LEFT
		} else {
			// open space, no boxes to enqueue
			continue
		}

		// enqueue first half of box
		first := n
		if seen[first.Y][first.X] != id {
			seen[first.Y][first.X] = id
			*todo = append(*todo, first)
		}

		// enqueue second half of box
		second := n.Add(otherDir)
		if seen[second.Y][second.X] != id {
			seen[second.Y][second.X] = id
			*todo = append(*todo, second)
		}
	}

	// If we reach here, no wall blocked us
	// Move boxes in reverse order
	for i := len(*todo) - 1; i >= 0; i-- {
		p := (*todo)[i]
		dst := p.Add(direction)
		grid.Set(dst, grid.Get(p))
		grid.Set(p, '.')
	}

	// Move robot
	*start = start.Add(direction)
}

// Stretch the grid for part2
func stretch(grid Grid) Grid {
	next := NewGrid(grid.Width*2, grid.Height, '.')

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			ch := grid.Data[y*grid.Width+x]
			var left, right byte
			switch ch {
			case '#':
				left, right = '#', '#'
			case 'O':
				left, right = '[', ']'
			case '@':
				left, right = '@', '.'
			case '.':
				// already filled, do nothing
				continue
			default:
				// Just in case
				continue
			}

			next.Set(Point{2 * x, y}, left)
			next.Set(Point{2*x + 1, y}, right)
		}
	}

	return next
}

// Compute GPS sum for boxes
func gps(grid Grid, needle byte) int {
	result := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if grid.Data[y*grid.Width+x] == needle {
				result += y*100 + x
			}
		}
	}
	return result
}

func part2(input Grid, moves string) int {
	// Pass `input` directly, not `&input`.
	grid := stretch(input)
	position, _ := grid.Find('@')

	todo := make([]Point, 0)
	seen := grid.SameSizeWith(-1)

	for id, b := range moves {
		switch b {
		case '<':
			narrow(&grid, &position, LEFT)
		case '>':
			narrow(&grid, &position, RIGHT)
		case '^':
			wide(&grid, &position, UP, &todo, seen, id)
		case 'v':
			wide(&grid, &position, DOWN, &todo, seen, id)
		}
	}

	return gps(grid, '[')
}

func main() {
	grid, moves := parseInput()
	fmt.Println("Part 1:", part1(grid, moves))
	fmt.Println("Part 2:", part2(grid, moves))
}
