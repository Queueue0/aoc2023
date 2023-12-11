package main

import (
	"bufio"
	"fmt"
	"os"
)

type Galaxy struct {
	X int
	Y int
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

	gals, eRows, eCols := parseLines(lines)

	sum := 0
	for i := 0; i < len(gals); i++ {
		for j := i + 1; j < len(gals); j++ {
			dist := 0
			lowX := 0
			highX := 0
			lowY := 0
			highY := 0

			if gals[i].X > gals[j].X {
				lowX = gals[j].X
				highX = gals[i].X
				dist += gals[i].X - gals[j].X
			} else {
				lowX = gals[i].X
				highX = gals[j].X
				dist += gals[j].X - gals[i].X
			}

			if gals[i].Y > gals[j].Y {
				lowY = gals[j].Y
				highY = gals[i].Y
				dist += gals[i].Y - gals[j].Y
			} else {
				lowY = gals[i].Y
				highY = gals[j].Y
				dist += gals[j].Y - gals[i].Y
			}
			
			// Solution to part 1 just replaced dist += 999999 with dist++ in these 2 loops
			for i := lowX; i < highX; i++ {
				if eRows[i] {
					dist += 999999
				}
			}

			for i := lowY; i < highY; i++ {
				if eCols[i] {
					dist += 999999
				}
			}

			sum += dist
		}
	}

	fmt.Println(sum)
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

func parseLines(lines []string) ([]Galaxy, map[int]bool, map[int]bool) {
	rlines := make([][]rune, len(lines))
	galaxies := []Galaxy{}
	emptyRows := map[int]bool{}
	emptyCols := map[int]bool{}

	for x, l := range lines {
		rlines[x] = []rune(l)
		emptyRows[x] = rowEmpty(rlines[x])

		if !emptyRows[x] {
			for y, r := range rlines[x] {
				if r == '#' {
					galaxies = append(galaxies, Galaxy{X: x, Y: y})
				}
			}
		}
	}

	for col := range rlines[0] {
		emptyCols[col] = colEmpty(rlines, col)
	}

	return galaxies, emptyRows, emptyCols
}

func rowEmpty(row []rune) bool {
	for _, r := range row {
		if r == '#' {
			return false
		}
	}

	return true
}

func colEmpty(chart [][]rune, col int) bool {
	for _, r := range chart {
		if r[col] == '#' {
			return false
		}
	}

	return true
}
