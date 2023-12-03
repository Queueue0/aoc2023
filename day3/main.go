package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Number struct {
	Value  int
	XPos   int
	YPos   int
	Length int
}

type Symbol struct {
	Value   rune
	XPos    int
	YPos    int
	Numbers []int
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("missing input file")
		os.Exit(1)
	}

	in, err := readLines(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sum := 0
	nums, syms := parseInput(in)

	for _, n := range nums {
		minX := n.XPos - 1
		maxX := n.XPos + n.Length
		minY := n.YPos - 1
		maxY := n.YPos + 1

		for i, s := range syms {
			if (s.XPos >= minX && s.XPos <= maxX) &&
				(s.YPos >= minY && s.YPos <= maxY) {

				syms[i].Numbers = append(s.Numbers, n.Value)
				sum += n.Value
				break
			}
		}
	}

	fmt.Println(sum)

	gearSum := 0
	for _, s := range syms {
		if s.Value != '*' {
			continue
		}

		if len(s.Numbers) != 2 {
			continue
		}

		gearSum += s.Numbers[0] * s.Numbers[1]
	}

	fmt.Println(gearSum)
}

func parseInput(input []string) ([]Number, []Symbol) {
	nums := []Number{}
	syms := []Symbol{}
	num := ""
	for y, s := range input {
		s = s + "."
		for x, c := range s {
			if unicode.IsDigit(c) {
				num += string(c)
			} else {
				if len(num) > 0 {
					l := len(num)
					xPos := x - l
					v, err := strconv.Atoi(num)
					if err != nil {
						fmt.Println(err)
					}

					nums = append(nums, Number{
						Value:  v,
						XPos:   xPos,
						YPos:   y,
						Length: l,
					})
					num = ""
				}

				if c != '.' {
					syms = append(syms, Symbol{
						Value: c,
						XPos: x,
						YPos: y,
						Numbers: []int{},
					})
				}
			}
		}
	}
	return nums, syms
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
