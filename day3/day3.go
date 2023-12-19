package day3

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Run() {

	data, err := os.ReadFile("day3/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines := strings.Split(string(data), "\n")
	lines = append(lines, strings.Repeat(".", len(lines[0])))
	lines = append([]string{strings.Repeat(".", len(lines[0]))}, lines...)

	for i, v := range lines {
		lines[i] = "." + v + "."
	}

	sum := 0
	gearRatios := make(map[string][]int)

	for y, line := range lines {
		fmt.Println(y, line, len(line))
		inNumber := false
		adjacent := false
		start := -1
		gears := []string{}
		for x, s := range line {
			if unicode.IsNumber(s) && !inNumber {
				inNumber = true
				start = x
				if isAdjacent(lines, x-1, y, &gears) {
					adjacent = true

				}
			}
			if inNumber {
				if isAdjacent(lines, x, y, &gears) {
					adjacent = true
				}
			}
			if !unicode.IsNumber(s) && inNumber {
				// adjacent = false
				number, _ := strconv.Atoi(line[start:x])
				fmt.Println(number, adjacent, gears)
				if adjacent {
					sum += number
				}
				for _, g := range gears {
					gearRatios[g] = append(gearRatios[g], number)
				}

				inNumber = false
				adjacent = false
				gears = []string{}
			}
		}
	}

	fmt.Println(sum)
	fmt.Println(gearRatios)

	sum2 := 0
	for _, v := range gearRatios {
		if len(v) == 2 {
			sum2 += v[0] * v[1]
		}
	}

	fmt.Println(sum2)

}

func isAdjacent(lines []string, px, py int, gears *[]string) bool {
	adjacent := false

	for y := py - 1; y <= py+1; y++ {
		c := lines[y][px]
		if !unicode.IsNumber(rune(c)) && c != '.' {
			adjacent = true
		}
		if c == '*' {
			id := fmt.Sprintf("%d-%d", px, y)
			if !slices.Contains(*gears, id) {
				*gears = append(*gears, id)
			}
		}
	}
	return adjacent
}
