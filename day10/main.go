package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Pipe struct {
	X int
	Y int
}

func (p Pipe) String() string {
	return fmt.Sprintf("X: %d Y: %d", p.X, p.Y)
}

var norths = []rune{'|', 'J', 'L'}

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

	pipes := parseLines(lines)
	s := Pipe{}
Outer:
	for x, row := range pipes {
		for y, r := range row {
			if r == 'S' {
				s.X, s.Y = x, y
				break Outer
			}
		}
	}

	fmt.Println(s)

	length := 1
	x, y := s.X, s.Y
	loop := map[Pipe]rune{}

	// I really didn't want to write code to figure out where to go
	// from the start so I cheated a bit and checked my input for initial direction
	dir := 'S'
	x = s.X + 1
	loop[s] = '7'

	for pipes[x][y] != 'S' {
		loop[Pipe{X: x, Y: y}] = pipes[x][y]
		length++
		if pipes[x][y] == 'L' {
			if dir == 'S' {
				dir = 'E'
			}

			if dir == 'W' {
				dir = 'N'
			}
		}

		if pipes[x][y] == 'J' {
			if dir == 'E' {
				dir = 'N'
			}

			if dir == 'S' {
				dir = 'W'
			}
		}

		if pipes[x][y] == '7' {
			if dir == 'E' {
				dir = 'S'
			}

			if dir == 'N' {
				dir = 'W'
			}
		}

		if pipes[x][y] == 'F' {
			if dir == 'N' {
				dir = 'E'
			}

			if dir == 'W' {
				dir = 'S'
			}
		}

		if dir == 'N' {
			x--
		}

		if dir == 'S' {
			x++
		}

		if dir == 'E' {
			y++
		}

		if dir == 'W' {
			y--
		}
	}

	fmt.Println(length/2)

	pipes[s.X][s.Y] = '7'

	isInside := false
	insideCount := 0
	for x, row := range pipes {
		isInside = false
		for y := range row {
			r, ok := loop[Pipe{X: x, Y: y}]

			if ok && slices.Contains(norths, r) {
				isInside = !isInside
			}

			if !ok && isInside {
				insideCount++
			}
		}
	}

	fmt.Println(insideCount)
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

func parseLines(lines []string) [][]rune {
	result := make([][]rune, len(lines))

	for i, l := range lines {
		result[i] = []rune(l)
	}

	return result
}
