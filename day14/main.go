package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rock rune

type Rocks []Rock

func (r Rocks) Less(i, j int) bool {
	if r[i] == 'O' && r[j] == '.' {
		return true
	}

	return false
}

func (r Rocks) Greater(i, j int) bool {
	if r[i] == '.' && r[j] == 'O' {
		return true
	}

	return false
}

func (r Rocks) Len() int {
	return len(r)
}

func (r Rocks) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Rocks) String() string {
	var out string
	for _, rock := range r {
		out += string(rock)
	}

	return out
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

	rs := parseLines(lines)

	rs2 := make([]Rocks, len(rs))
	for i := range rs {
		rs2[i] = make(Rocks, len(rs[i]))
		copy(rs2[i], rs[i])
	}

	for _, r := range rs {
		fmt.Println(r)
	}

	fmt.Println()

	for i := 0; i < 1000; i++ {
		transform(rs, 'N')
		transform(rs, 'W')
		transform(rs, 'S')
		transform(rs, 'E')
	}

	sum := 0
	for i, r := range rs {
		rowScore := len(rs) - i
		fmt.Println(r, rowScore, strings.Count(string(r), "O"), strings.Count(string(r), "O")*rowScore)
		sum += strings.Count(string(r), "O") * rowScore
	}

	fmt.Println(verify(rs2, rs))

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

func parseLines(lines []string) []Rocks {
	rs := make([]Rocks, len(lines))

	for i := range lines {
		rs[i] = Rocks(lines[i])
	}

	return rs
}

func transform(rs []Rocks, dir rune) []Rocks {
	if dir == 'N' || dir == 'S' {
		for i := 0; i < len(rs[0]); i++ {
			col := make(Rocks, len(rs))
			for j := 0; j < len(rs); j++ {
				col[j] = rs[j][i]
			}

			if dir == 'N' {
				// sort.Stable wasn't stable enough when using my weird Less method
				// so I just used insertion sort
				insertionSort(col)
			}

			if dir == 'S' {
				insertionSortReverse(col)
			}

			for j := 0; j < len(rs); j++ {
				rs[j][i] = col[j]
			}
		}
	}

	if dir == 'W' || dir == 'E' {
		for i := range rs {
			if dir == 'W' {
				insertionSort(rs[i])
			}

			if dir == 'E' {
				insertionSortReverse(rs[i])
			}
		}
	}

	return rs
}

func insertionSort(data Rocks) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func insertionSortReverse(data Rocks) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0 && data.Greater(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func verify(r1, r2 []Rocks) bool {
	for i := range r1 {
		for j := range r1[i] {
			if r1[i][j] == '#' && r2[i][j] != '#' {
				fmt.Println(i, j)
				return false
			}
		}
	}

	return true
}
