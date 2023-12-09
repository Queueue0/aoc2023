package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	// Part 1
	hists := parseLines(lines)
	total := 0
	for _, h := range hists {
		diffs := [][]int{}
		diffs = append(diffs, h)
		prediction := 0 
		for i := 1; !allSame(diffs[len(diffs)-1]); i++ {
			diffs = append(diffs, []int{})
			for j := 1; j < len(diffs[i-1]); j++ {
				diffs[i] = append(diffs[i], diffs[i-1][j] - diffs[i-1][j-1])
			}
			
		}

		for _, d := range diffs {
			prediction += d[len(d)-1]
		}

		total += prediction
	}

	fmt.Println(total)

	// Part 2
	total = 0
	for _, h := range hists {
		diffs := [][]int{}
		diffs = append(diffs, h)
		prediction := 0
		for i := 1; !allSame(diffs[len(diffs)-1]); i++ {
			diffs = append(diffs, []int{})
			for j := 1; j < len(diffs[i-1]); j++ {
				diffs[i] = append(diffs[i], diffs[i-1][j] - diffs[i-1][j-1])
			}
		}

		// This loop is the only difference between the 2 solutions
		for i := len(diffs)-1; i >= 0; i-- {
			prediction = diffs[i][0] - prediction
		}

		total += prediction
	}

	fmt.Println(total)
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

func parseLines(lines []string) [][]int {
	output := make([][]int, len(lines))

	for i, l := range lines {
		sints := strings.Fields(l)
		ints := make([]int, len(sints))
		for j, s := range sints {
			ints[j], _ = strconv.Atoi(s)
		}
		output[i] = ints
	}

	return output
}

func allSame(ints []int) bool {
	for i := 1; i < len(ints); i++ {
		if ints[i] != ints[i-1] {
			return false
		}
	}
	return true
}
