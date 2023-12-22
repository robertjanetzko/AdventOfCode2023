package day21

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var lines []string

type coords struct {
	x, y, d int
}

var open, closed []coords

func Run() {
	data, err := os.ReadFile("day21/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines = strings.Split(string(data), "\n")
	for _, v := range lines {
		fmt.Println(v)
	}

	start := findStart()
	fmt.Println(start)

	open = append(open, start)

	for len(open) > 0 {
		next := open[0]
		open = open[1:]
		closed = append(closed, next)
		if next.d < 26501365 {
			check(coords{next.x - 1, next.y, next.d + 1})
			check(coords{next.x + 1, next.y, next.d + 1})
			check(coords{next.x, next.y - 1, next.d + 1})
			check(coords{next.x, next.y + 1, next.d + 1})
		}
	}
	sum := 0
	for _, v := range closed {
		if v.d%1 == 0 {
			fmt.Println(v)
			sum++
		}
	}
	fmt.Println(sum)
}

func check(c coords) {
	// if c.x < 0 || c.y < 0 || c.x >= len(lines[0]) || c.y >= len(lines) {
	// 	return
	// }
	x := mod(c.x, len(lines[0]))
	y := mod(c.y, len(lines))
	if lines[y][x:x+1] == "#" {
		return
	}
	if slices.ContainsFunc(closed, func(c2 coords) bool { return x == c2.x && y == c2.y }) {
		return
	}
	if slices.ContainsFunc(open, func(c2 coords) bool { return x == c2.x && y == c2.y }) {
		return
	}
	open = append(open, c)
}

func mod(x, m int) int {
	return (x%m + m) % m
}

func findStart() coords {
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if lines[y][x:x+1] == "S" {
				return coords{x, y, 0}
			}
		}
	}
	return coords{-1, -1, 0}
}
