package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Hand struct {
	Cards    string
	Wager    int
	Strength int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing input argument")
		os.Exit(1)
	}

	lines, err := readLines(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hands := parseInput(lines)
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].lessThan(hands[j])
	})

	sum := 0
	for i, h := range hands {
		fmt.Printf("%+v\n", h)
		sum += (i + 1) * h.Wager
	}

	fmt.Println(sum)
}

func parseInput(lines []string) []Hand {
	hands := make([]Hand, len(lines))
	for i, l := range lines {
		rawHand := strings.Split(l, " ")
		cards := rawHand[0]
		wager, _ := strconv.Atoi(rawHand[1])

		cardset := make(map[rune]int)
		for _, r := range cards {
			cardset[r] += 1
		}

		_, ok := cardset['J']
		if ok && len(cardset) > 1 {
			var c rune
			m := 0
			for k, v := range cardset {
				if v > m && k != 'J' {
					m = v
					c = k
				}
			}

			cards = strings.Replace(cards, "J", string(c), -1)

			cardset = make(map[rune]int)
			for _, r := range cards {
				cardset[r] += 1
			}
		}

		strength := 0
		if len(cardset) == 1 {
			strength = 6
		}

		if len(cardset) == 2 {
			firstCount := cardset[[]rune(cards)[0]]
			if firstCount == 4 || firstCount == 1 {
				strength = 5
			} else {
				strength = 4
			}
		}

		if len(cardset) == 3 {
			i := 0
			for strength == 0 {
				firstCount := cardset[[]rune(cards)[i]]
				if firstCount == 3 {
					strength = 3
				}

				if firstCount == 2 {
					strength = 2
				}

				i++
			}
		}

		if len(cardset) == 4 {
			strength = 1
		}

		hands[i] = Hand{
			Cards:    rawHand[0],
			Wager:    wager,
			Strength: strength,
		}
	}
	return hands
}

var valueMap = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 0,
}

func (h *Hand) lessThan(other Hand) bool {
	if h.Strength != other.Strength {
		return h.Strength < other.Strength
	}

	for i, c := range other.Cards {
		hc := []rune(h.Cards)[i]
		if valueMap[hc] != valueMap[c] {
			return valueMap[hc] < valueMap[c]
		}
	}

	return false
}
