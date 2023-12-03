package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func parseLine(line string) (int, int, int, int) {
	maxColors := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	strId := line[strings.Index(line, " ")+1 : strings.Index(line, ":")]
	id, err := strconv.Atoi(strId)
	if err != nil {
		fmt.Println(err)
		return 0, 0, 0, 0
	}

	line = line[strings.Index(line, ":")+2:]
	pulls := strings.Split(line, "; ")

	for _, pull := range pulls {
		colors := strings.Split(pull, ", ")

		for _, color := range colors {
			cube := strings.Split(color, " ")

			count, err := strconv.Atoi(cube[0])
			if err != nil {
				fmt.Println(err)
				count = 0
			}
			color := cube[1]

			if count > maxColors[color] {
				maxColors[color] = count
			}
		}
	}

	return id, maxColors["red"], maxColors["green"], maxColors["blue"]
}

func possible(red, green, blue int) bool {
	const maxRed int = 12
	const maxGreen int = 13
	const maxBlue int = 14

	return red <= maxRed && green <= maxGreen && blue <= maxBlue
}

type Game struct {
	id    int
	red   int
	green int
	blue  int
}

func main() {
	in, err := readLines("./input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	games := []Game{}
	sum := 0
	for _, line := range in {
		id, red, green, blue := parseLine(line)

		games = append(games, Game{id: id, red: red, green: green, blue: blue})
		if possible(red, green, blue) {
			sum += id
		}
	}

	fmt.Println(sum)

	powerSum := 0
	for _, game := range games {
		powerSum += game.red * game.green * game.blue
	}

	fmt.Println(powerSum)
}
