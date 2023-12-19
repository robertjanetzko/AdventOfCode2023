package day11

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type coords struct {
	x, y int
}

func (c *coords) MinDist(list []coords) (int, int) {
	dist := int(^uint(0) >> 1)
	other := 0
	for i, v := range list {
		if d := c.Dist(v); d != 0 && d < dist {
			other = i
			dist = d
		}
	}
	return dist, other
}

func (c *coords) Dist(c2 coords) int {
	return Abs(c.x-c2.x) + Abs(c.y-c2.y)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

const part2 = true
const mult = 1000000

func Run() {
	data, err := os.ReadFile("day11/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines := strings.Split(string(data), "\n")

	// expand
	for i := len(lines) - 1; i >= 0; i-- {
		if !strings.Contains(lines[i], "#") {
			if !part2 {
				lines = slices.Insert(lines, i, lines[i])
			} else {
				lines[i] = strings.ReplaceAll(lines[i], ".", "*")
			}
		}
	}
	for i := len(lines[0]) - 1; i >= 0; i-- {
		var colHasGalaxy = false
		for _, v := range lines {
			if v[i:i+1] == "#" {
				colHasGalaxy = true
				break
			}
		}
		if !colHasGalaxy {
			for j := len(lines) - 1; j >= 0; j-- {
				if !part2 {
					lines[j] = lines[j][:i] + "." + lines[j][i:]
				} else {
					lines[j] = lines[j][:i] + "*" + lines[j][i+1:]
				}
			}
		}
	}

	var galaxies []coords
	ev := 0
	for y, line := range lines {
		if !strings.Contains(line, "#") {
			ev++
			continue
		}
		eh := 0
		for x, v := range line {
			if v == '*' {
				eh++
			}
			if v == '#' {
				galaxies = append(galaxies, coords{x + ((mult - 1) * eh), y + ((mult - 1) * ev)})
			}
		}
	}

	for _, v := range lines {
		fmt.Println(v)
	}

	sum := 0
	for _, a := range galaxies {
		for _, b := range galaxies {
			sum += a.Dist(b)
		}
	}
	fmt.Println(sum / 2)

}
