package day10

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func Run() {
	data, err := os.ReadFile("day10/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines := strings.Split(string(data), "\n")

	var sx, sy int
	for y, line := range lines {
		sx = strings.Index(line, "S")
		if sx != -1 {
			sy = y
			break
		}
	}
	fmt.Println(sx, sy)

	start := coords{x: sx, y: sy}
	var starts []coords
	for x := max(sx-1, 0); x <= sx+1; x++ {
		for y := max(sy-1, 0); y <= sy+1; y++ {
			if x != sx || y != sy {
				pipe := lines[y][x : x+1]
				if pipe == "." {
					continue
				}
				fmt.Println(start, ">", next(lines, coords{x, y}, start))
				if opt := next(lines, coords{x, y}, start); opt != noCoords {
					starts = append(starts, coords{x, y})
				}
			}
		}
	}
	fmt.Println(starts)

	var paths [][]coords
	paths = append(paths, []coords{start, start})
	paths = append(paths, starts)
	var onPath []coords
	onPath = append(onPath, start)
	onPath = append(onPath, starts...)
	fmt.Println("PATHS", paths)
	for {
		prevs := paths[len(paths)-2]
		starts := paths[len(paths)-1]
		if starts[0] == starts[1] {
			break
		}
		nexts := []coords{next(lines, starts[0], prevs[0]), next(lines, starts[1], prevs[1])}
		paths = append(paths, nexts)
		onPath = append(onPath, nexts...)
		fmt.Println("STEP", len(paths), nexts)
	}
	fmt.Println("LEN", len(paths)-1)

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if !slices.Contains(onPath, coords{x, y}) {
				lines[y] = lines[y][:x] + "." + lines[y][x+1:]
			}
		}
		// fmt.Println(lines[y])
	}
	sum := 0
	for y := 0; y < len(lines); y++ {
		cnt := 0
		for x := 0; x < len(lines[0]); x++ {
			pipe := lines[y][x : x+1]
			if pipe == "|" || pipe == "L" || pipe == "J" {
				cnt++
			}
			if pipe == "." {

				if cnt%2 == 0 {
					lines[y] = lines[y][:x] + " " + lines[y][x+1:]
				} else {
					lines[y] = lines[y][:x] + "X" + lines[y][x+1:]
					sum++
				}
			}
		}
		fmt.Println(lines[y])
	}
	fmt.Println("AREA", sum)
}

type coords struct {
	x, y int
}

var noCoords = coords{-1, -1}

var dirs = map[string][]coords{
	"|": {{0, -1}, {0, 1}},
	"-": {{-1, 0}, {1, 0}},
	"L": {{0, -1}, {1, 0}},
	"J": {{0, -1}, {-1, 0}},
	"7": {{0, 1}, {-1, 0}},
	"F": {{0, 1}, {1, 0}},
}

func next(lines []string, at, from coords) coords {
	pipe := lines[at.y][at.x : at.x+1]
	opts := dirs[pipe]
	// fmt.Println("N", at, from, pipe, add(at, opts[0]), add(at, opts[1]), add(at, opts[0]) == from, add(at, opts[1]) == from)
	if opt := add(at, opts[0]); opt == from {
		return add(at, opts[1])
	} else if opt := add(at, opts[1]); opt == from {
		return add(at, opts[0])
	} else {
		return noCoords
	}
}

func add(a, b coords) coords {
	return coords{x: a.x + b.x, y: a.y + b.y}
}
