package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	MaxTime int
	MinDist int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing input file arg")
		os.Exit(1)
	}

	lines, err := readLines(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Part 1 (original)
	//	races := parseInput(lines)
	//
	//	total := 1
	//	for _, race := range races {
	//		half := race.MaxTime / 2
	//		n := half / 2
	//		mini := 0
	//		maxi := half
	//		fmt.Println(race, mini, maxi, n)
	//
	//		for !race.IsFirst(n) {
	//			if race.GetDist(n) > race.MinDist {
	//				maxi = n - 1
	//			} else {
	//				mini = n + 1
	//			}
	//			n = (mini + maxi) / 2
	//			fmt.Println(race, mini, maxi, n)
	//		}
	//
	//		ways := (half - n + 1) * 2
	//		if race.MaxTime%2 == 0 {
	//			ways--
	//		}
	//
	//		total *= ways
	//	}

	// Part 1 (Quadratic)
	races := parseInput(lines)

	total := 1
	for _, race := range races {
		ways := race.GetWays()
		total *= int(ways)
	}

	fmt.Println(total)

	// Part 2 (original)
	//	race := parseInput2(lines)
	//
	//	half := race.MaxTime / 2
	//	n := half / 2
	//	mini := 0
	//	maxi := half
	//	fmt.Println(race, mini, maxi, n)
	//
	//	for !race.IsFirst(n) {
	//		if race.GetDist(n) > race.MinDist {
	//			maxi = n - 1
	//		} else {
	//			mini = n + 1
	//		}
	//		n = (mini + maxi) / 2
	//		fmt.Println(race, mini, maxi, n)
	//	}
	//
	//	ways := (half - n + 1) * 2
	//	if race.MaxTime%2 == 0 {
	//		ways--
	//	}

	// Part 2 (quadratic)
	race := parseInput2(lines)

	ways := race.GetWays()
	fmt.Println(int(ways))
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

func parseInput(lines []string) []Race {
	races := []Race{}

	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])

	for i := range times {
		time, _ := strconv.Atoi(times[i])
		dist, _ := strconv.Atoi(distances[i])

		races = append(races, Race{MaxTime: time, MinDist: dist})
	}

	return races
}

func parseInput2(lines []string) Race {
	time := strings.Join(strings.Fields(strings.Split(lines[0], ":")[1]), "")
	distance := strings.Join(strings.Fields(strings.Split(lines[1], ":")[1]), "")

	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(distance)

	return Race{MaxTime: t, MinDist: d}
}

func (r *Race) GetDist(n int) int {
	return (r.MaxTime - n) * n
}

func (r *Race) IsFirst(n int) bool {
	// If n beats the distance and n-1 doesn't
	return (r.GetDist(n) > r.MinDist && r.GetDist(n-1) <= r.MinDist)
}

func (r *Race) GetWays() int {
	b := float64(-r.MaxTime)
	c := float64(r.MinDist + 1)
	first := math.Ceil((-b - math.Sqrt((b*b)-(4*c))) / 2)
	last := math.Floor((-b + math.Sqrt((b*b)-(4*c))) / 2)
	ways := last - first + 1

	return int(ways)
}
