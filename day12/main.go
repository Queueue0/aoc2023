package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	Springs string
	Groups  string
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

	records := parseLines(lines)

	sum := 0
	for _, r := range records {
		sum += calculate(r)
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

func parseLines(lines []string) []Record {
	records := []Record{}

	for _, l := range lines {
		fields := strings.Fields(l)

		r := Record{
			Springs: strings.Repeat(fields[0]+"?", 4) + fields[0],
			Groups:  strings.Repeat(fields[1]+",", 4) + fields[1],
		}

		records = append(records, r)
	}

	return records
}

var cache = make(map[Record]int)

func calculate(r Record) int {
	if v, ok := cache[r]; ok {
		return v
	}

	if len(r.Groups) == 0 {
		if !strings.Contains(r.Springs, "#") {
			cache[r] = 1
			return 1
		} else {
			cache[r] = 0
			return 0
		}
	}

	if len(r.Springs) == 0 {
		cache[r] = 0
		return 0
	}

	var result int
	nextRune := []rune(r.Springs)[0]
	nextGroup, _ := strconv.Atoi(strings.Split(r.Groups, ",")[0])

	if nextRune == '#' {
		result = pound(r, nextGroup)
	}

	if nextRune == '.' {
		result = dot(r)
	}

	if nextRune == '?' {
		result = dot(r) + pound(r, nextGroup)
	}

	cache[r] = result
	return result
}

func dot(r Record) int {
	nr := Record{
		Springs: string([]rune(r.Springs)[1:]),
		Groups:  r.Groups,
	}

	return calculate(nr)
}

func pound(r Record, nextGroup int) int {
	currentGroup := string([]rune(r.Springs)[:nextGroup])
	currentGroup = strings.ReplaceAll(currentGroup, "?", "#")

	if currentGroup != strings.Repeat("#", nextGroup) {
		return 0
	}

	if len([]rune(r.Springs)) == nextGroup {
		if len(strings.Split(r.Groups, ",")) == 1 {
			return 1
		} else {
			return 0
		}
	}

	nextRune := []rune(r.Springs)[nextGroup]
	if nextRune == '?' || nextRune == '.' {
		nr := Record{
			Springs: string([]rune(r.Springs)[nextGroup+1:]),
			Groups:  strings.Join(strings.Split(r.Groups, ",")[1:], ","),
		}
		return calculate(nr)
	}

	return 0
}
