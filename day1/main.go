package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

var wds map[string]rune = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
	"zero":  '0',
}

func firstWordDigit(line string) (rune, int) {
	lowest := -1
	var fwd rune

	for s, r := range wds {
		i := strings.Index(line, s)
		if i != -1 && (i < lowest || lowest == -1) {
			lowest = i
			fwd = r
		}
	}

	return fwd, lowest
}

func lastWordDigit(line string) (rune, int) {
	highest := -1
	var lwd rune

	for s, r := range wds {
		i := strings.LastIndex(line, s)
		if i != -1 && (i > highest || highest == -1) {
			highest = i
			lwd = r
		}
	}

	return lwd, highest
}

func main() {
	in, err := readLines("./input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sum := 0
	for _, s := range in {
		rs := []rune(s)
		var d string

		fnd := len(rs) + 1
		var fndr rune
		for i, r := range rs {
			if unicode.IsDigit(r) {
				fndr = r
				fnd = i
				break
			}
		}

		fwdr, fwd := firstWordDigit(s)

		if fwd < fnd && fwd != -1 {
			d += string(fwdr)
		} else {
			d += string(fndr)
		}

		var lnd int
		var lndr rune
		for i := len(rs) - 1; i >= 0; i-- {
			if unicode.IsDigit(rs[i]) {
				lndr = rs[i]
				lnd = i
				break
			}
		}

		lwdr, lwd := lastWordDigit(s)

		if lwd > lnd && lwd != -1 {
			d += string(lwdr)
		} else {
			d += string(lndr)
		}

		dint, err := strconv.Atoi(d)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sum += dint
	}

	fmt.Println(sum)
}
