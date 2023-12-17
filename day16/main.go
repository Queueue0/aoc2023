package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tile struct {
	Char      rune
	Energized bool
}

type Grid [][]Tile

func (g Grid) Traverse(x, y int, dir rune) {
	for x >= 0 && x < len(g) && y >= 0 && y < len(g[0]) {
		if g[x][y].Char == '|' && (dir == 'L' || dir == 'R') {
			if g[x][y].Energized {
				break
			}
			dir = 'D'
			g.Traverse(x-1, y, 'U')
		}

		if g[x][y].Char == '-' && (dir == 'U' || dir == 'D') {
			if g[x][y].Energized {
				break
			}
			dir = 'L'
			g.Traverse(x, y+1, 'R')
		}

		if g[x][y].Char == '/' {
			switch dir {
			case 'R':
				dir = 'U'
			case 'L':
				dir = 'D'
			case 'U':
				dir = 'R'
			case 'D':
				dir = 'L'
			}
		}

		if g[x][y].Char == '\\' {
			switch dir {
			case 'R':
				dir = 'D'
			case 'L':
				dir = 'U'
			case 'U':
				dir = 'L'
			case 'D':
				dir = 'R'
			}
		}

		g[x][y].Energized = true

		if dir == 'U' {
			x--
		}

		if dir == 'D' {
			x++
		}

		if dir == 'L' {
			y--
		}

		if dir == 'R' {
			y++
		}
	}
}

func (g Grid) CountEnergized() int {
	total := 0
	for x := range g {
		for y := range g[0] {
			if g[x][y].Energized {
				total++
			}
		}
	}

	return total
}

func (g Grid) Reset() {
	for x := range g {
		for y := range g[0] {
			g[x][y].Energized = false
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing input arg")
		os.Exit(1)
	}

	lines, err := readLines(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grid := parseLines(lines)

	me := 0
	for x := range grid {
		grid.Traverse(x, 0, 'R')
		me = max(me, grid.CountEnergized())
		grid.Reset()

		grid.Traverse(x, len(grid[x])-1, 'L')
		me = max(me, grid.CountEnergized())
		grid.Reset()
	}

	for y := range grid[0] {
		grid.Traverse(0, y, 'D')
		me = max(me, grid.CountEnergized())
		grid.Reset()

		grid.Traverse(len(grid)-1, y, 'U')
		me = max(me, grid.CountEnergized())
		grid.Reset()
	}

	fmt.Println(me)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseLines(lines []string) Grid {
	g := make([][]Tile, len(lines))

	for i := range lines {
		for _, r := range []rune(lines[i]) {
			g[i] = append(g[i], Tile{Char: r, Energized: false})
		}
	}

	return g
}
