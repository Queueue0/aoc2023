package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Value string
	Left  *Node
	Right *Node
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

	// Part 1
	instructs, cn := parseLines(lines)
	steps := 0
	for i := 0; cn.Value != "ZZZ"; i++ {
		i %= len(instructs)
		steps++

		if instructs[i] == 'L' {
			cn = cn.Left
		}

		if instructs[i] == 'R' {
			cn = cn.Right
		}
	}

	fmt.Println(steps, cn.Value)

	// Part 2
	instructs, aNodes := parseLines2(lines)
	allSteps := make([]int, len(aNodes))
	for i := 0; i < len(aNodes); i++ {
		steps = 0
		for j := 0; strings.LastIndex(aNodes[i].Value, "Z") != 2; j++ {
			j %= len(instructs)
			steps++

			if instructs[j] == 'L' {
				aNodes[i] = aNodes[i].Left
			}

			if instructs[j] == 'R' {
				aNodes[i] = aNodes[i].Right
			}
		}
		allSteps[i] = steps
	}


	fmt.Println(LCM(allSteps[0], allSteps[1], allSteps[2:]...))
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

func parseLines(lines []string) ([]rune, *Node) {
	instructions := lines[0]
	nodeLines := lines[2:]
	AAA := 0
	nodes := make([]*Node, len(nodeLines))
	children := make([]string, len(nodeLines))
	
	for i, l := range nodeLines {
		halves := strings.Split(l, " = ")
		nodes[i] = &Node{Value: halves[0]}
		if nodes[i].Value == "AAA" {
			AAA = i
		}
		children[i] = strings.Replace(strings.Replace(halves[1], ")", "", -1), "(", "", -1)
	}

	for i, n := range nodes {
		lr := strings.Split(children[i], ", ")
		for _, o := range nodes {
			if o.Value == lr[0] {
				n.Left = o
			}
			
			if o.Value == lr[1] {
				n.Right = o
			}
		}
	}

	return []rune(instructions), nodes[AAA]
}

func parseLines2(lines []string) ([]rune, []*Node) {
	instructions := lines[0]
	nodeLines := lines[2:]
	nodes := make([]*Node, len(nodeLines))
	aNodes := []*Node{}
	children := make([]string, len(nodeLines))
	
	for i, l := range nodeLines {
		halves := strings.Split(l, " = ")
		nodes[i] = &Node{Value: halves[0]}
		if strings.LastIndex(nodes[i].Value, "A") == 2 {
			aNodes = append(aNodes, nodes[i])
		}
		children[i] = strings.Replace(strings.Replace(halves[1], ")", "", -1), "(", "", -1)
	}

	for i, n := range nodes {
		lr := strings.Split(children[i], ", ")
		for _, o := range nodes {
			if o.Value == lr[0] {
				n.Left = o
			}
			
			if o.Value == lr[1] {
				n.Right = o
			}
		}
	}

	return []rune(instructions), aNodes
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func LCM(a, b int, ints ...int) int {
	result := a * b / GCD(a, b)

	for _, i := range ints {
		result = LCM(result, i)
	}

	return result
}
