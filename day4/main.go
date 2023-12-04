package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide input file")
		os.Exit(1)
	}

	lines, err := readLines(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cards := parseInput(lines)

	// Part 1
	sum := 0
	for _, c := range cards {
		points := 0
		for _, n := range c.Have {
			if slices.Contains(c.Winning, n) {
				if points == 0 {
					points += 1
				} else {
					points *= 2
				}
			}
		}

		sum += points
	}

	fmt.Println(sum)

	// Part 2
	for i, c := range cards {
		hits := 0
		for _, n := range c.Have {
			if slices.Contains(c.Winning, n) {
				hits += 1
			}
		}

		for j := i + 1; j <= i+hits; j++ {
			cards[j].Instances += c.Instances
		}
	}

	totalCards := 0
	for _, c := range cards {
		totalCards += c.Instances
	}

	fmt.Println(totalCards)
}

type Card struct {
	Winning   []int
	Have      []int
	Instances int
}

func parseInput(lines []string) []Card {
	cards := []Card{}

	for _, line := range lines {
		card := Card{
			Winning:   []int{},
			Have:      []int{},
			Instances: 1,
		}

		colon := strings.Index(line, ":")

		line = line[colon+1:]

		line = strings.ReplaceAll(line, "  ", " ")

		halves := strings.Split(line, "|")

		halves[0] = strings.TrimSpace(halves[0])
		halves[1] = strings.TrimSpace(halves[1])

		swn := strings.Split(halves[0], " ")
		shn := strings.Split(halves[1], " ")

		for _, s := range swn {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
			}
			card.Winning = append(card.Winning, n)
		}

		for _, s := range shn {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
			}
			card.Have = append(card.Have, n)
		}

		cards = append(cards, card)
	}

	return cards
}
