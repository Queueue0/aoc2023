package main

import (
	"bufio"
	"fmt"
	"os"
)

type pattern [][]rune

func (p pattern) String() string {
	output := ""
	for _, s := range p {
		output += string(s) + "\n"
	}

	output += fmt.Sprintf("%d", p.HReflections(0)) + "\n"
	output += fmt.Sprintf("%d", p.VReflections(0)) + "\n"

	return output
}

func (p pattern) HReflections(exclude int) int {
	for i := 0; i < len(p[0])-1; i++ {
		delta := 0
		success := true

		for i-delta >= 0 && i+delta+1 < len(p[0]) {
			l, u := i-delta, i+delta+1

			if !p.ColumnsEQ(l, u) {
				success = false
				break
			}

			delta++
		}

		if success {
			if i + 1 != exclude {
				return i + 1
			}
		}
	}
	return 0
}

func (p pattern) VReflections(exclude int) int {
	for i := 0; i < len(p)-1; i++ {
		delta := 0
		success := true

		for i-delta >= 0 && i+delta+1 < len(p) {
			l, u := i-delta, i+delta+1

			if string(p[l]) != string(p[u]) {
				success = false
				break
			}

			delta++
		}

		if success {
			if i + 1 != exclude {
				return i + 1
			}
		}
	}
	return 0
}

func (p pattern) ColumnsEQ(i, j int) bool {
	for _, r := range p {
		if r[i] != r[j] {
			return false
		}
	}

	return true
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

	patterns := parseLines(lines)

	total := 0
	for _, p := range patterns {
		oh, ov := p.HReflections(0), p.VReflections(0)
		ot := total
	Inner:
		for i := range p {
			for j := range p[i] {
				c := make(pattern, len(p))
				for x, line := range p {
					c[x] = []rune(string(line))
				}
				if c[i][j] == '.' {
					c[i][j] = '#'
				} else {
					c[i][j] = '.'
				}
				ch, cv := c.HReflections(oh), c.VReflections(ov)
				if max(ch, cv) != 0 {
					if ch != oh && ch != 0 {
						total += ch
						break Inner
					} else if cv != ov && cv != 0{
						total += 100 * cv
						break Inner
					}
				}
			}
		}

		// This is for debugging.
		// If any of the patterns get printed out then something is wrong
		if ot == total {
			fmt.Println(p)
		}
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

func parseLines(lines []string) []pattern {
	patterns := []pattern{}

	p := pattern{}
	for _, line := range lines {
		if len(line) > 0 {
			p = append(p, []rune(line))
		} else {
			patterns = append(patterns, p)
			p = pattern{}
		}
	}

	patterns = append(patterns, p)

	return patterns
}
