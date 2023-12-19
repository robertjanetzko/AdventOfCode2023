package day16

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var lines []string
var energized [][]bool
var cache = map[int]bool{}

type coords struct {
	x, y int
}

func Run() {
	data, err := os.ReadFile("day16/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines = strings.Split(string(data), "\n")
	for _, v := range lines {
		fmt.Println(v)
	}

	r := 0
	for i := 0; i < len(lines); i++ {
		r = max(r, calc(coords{-1, i}, coords{1, 0}))
		r = max(r, calc(coords{len(lines), i}, coords{-1, 0}))
		r = max(r, calc(coords{i, -1}, coords{0, 1}))
		r = max(r, calc(coords{i, len(lines)}, coords{0, -1}))
	}
	fmt.Println(r)
}

func calc(curPos, curDir coords) int {
	// fmt.Println(curPos, curDir)
	reset()
	follow(curPos, curDir)
	r := energyLevel()
	// fmt.Println(r)
	return r
}

func reset() {
	cache = map[int]bool{}
	energized = make([][]bool, len(lines))
	for i := range energized {
		energized[i] = make([]bool, len(lines[0]))
	}
}

func energyLevel() int {
	sum := 0
	for _, v := range energized {
		for _, e := range v {
			if e {
				sum++
				// fmt.Print("#")
			} else {
				// fmt.Print(".")
			}
		}
		// fmt.Println()
	}
	return sum
}

func follow(curPos, curDir coords) {
	curPos = coords{curPos.x + curDir.x, curPos.y + curDir.y}

	k := cacheKey(curPos, curDir)
	// fmt.Println(curPos, curDir, k)
	if v, ok := cache[k]; ok && v {
		return
	} else {
		cache[k] = true
	}

	if curPos.x < 0 || curPos.y < 0 || curPos.x >= len(lines[0]) || curPos.y >= len(lines) {
		return
	}

	c := lines[curPos.y][curPos.x : curPos.x+1]
	energized[curPos.y][curPos.x] = true
	if c == "." {
		follow(curPos, curDir)
	} else {
		if c == "/" {
			if curDir.y == 0 {
				follow(curPos, coords{0, -curDir.x})
			} else {
				follow(curPos, coords{-curDir.y, 0})
			}
		} else if c == "\\" {
			if curDir.y == 0 {
				follow(curPos, coords{0, curDir.x})
			} else {
				follow(curPos, coords{curDir.y, 0})
			}
		} else if c == "|" {
			if curDir.y == 0 {
				follow(curPos, coords{0, 1})
				follow(curPos, coords{0, -1})
			} else {
				follow(curPos, curDir)
			}
		} else if c == "-" {
			if curDir.y == 0 {
				follow(curPos, curDir)
			} else {
				follow(curPos, coords{1, 0})
				follow(curPos, coords{-1, 0})
			}
		}
	}
}

func cacheKey(curPos, curDir coords) int {
	return curPos.x*10000000 + curPos.y*1000 + 100 + curDir.x*10 + 10 + curDir.y
}
