package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Step struct {
	Label     string
	Operation rune
	FocalLen  int
}

func (s Step) String() string {
	if s.Operation == '=' {
		return fmt.Sprintf("{%s %c %d}", s.Label, s.Operation, s.FocalLen)
	}
	return fmt.Sprintf("{%s %c}", s.Label, s.Operation)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing input arg")
		os.Exit(1)
	}

	steps, err := readInput(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Go's maps aren't ordered, so I had to do some degeneracy to
	// keep track of the positions of lenses
	boxes := make(map[int]map[string][2]int)
	for _, step := range steps {
		box := hash(step.Label)
		if _, ok := boxes[box]; !ok {
			boxes[box] = make(map[string][2]int)
		}
		if step.Operation == '-' {
			val, ok := boxes[box][step.Label]
			if ok {
				pos := val[0]
				delete(boxes[box], step.Label)
				for k, v := range boxes[box] {
					if v[0] > pos {
						boxes[box][k] = [2]int{v[0]-1, v[1]}
					}
				}
			}
		} else {
			v, ok := boxes[box][step.Label]
			if ok {
				boxes[box][step.Label] = [2]int{v[0], step.FocalLen}
			} else {
				boxes[box][step.Label] = [2]int{len(boxes[box]), step.FocalLen}
			}
		}
	}

	total := 0
	for n, box := range boxes {
		fmt.Println(box)
		for _, v := range box {
			total += (n+1) * (v[0]+1) * v[1]
		}
	}

	fmt.Println(total)
}

func readInput(filename string) ([]Step, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Step{}, err
	}

	st := strings.TrimSpace(string(b))
	rawSteps := strings.Split(st, ",")
	steps := make([]Step, len(rawSteps))
	for i, s := range rawSteps {
		if strings.Contains(s, "=") {
			parts := strings.Split(s, "=")
			l := parts[0]
			fl, _ := strconv.Atoi(parts[1])
			steps[i] = Step{
				Label:     l,
				Operation: '=',
				FocalLen:  fl,
			}
		} else {
			dash := strings.Index(s, "-")
			steps[i] = Step{
				Label:     s[:dash],
				Operation: '-',
			}
		}
	}

	return steps, nil
}

func hash(s string) int {
	value := 0
	for _, r := range s {
		value += int(r)
		value *= 17
		value %= 256
	}

	return value
}
