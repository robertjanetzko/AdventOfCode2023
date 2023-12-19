package day14

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var lines []string

func Run() {
	data, err := os.ReadFile("day14/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines = strings.Split(string(data), "\n")

	for _, v := range lines {
		fmt.Println(v)
	}

	cache := make(map[string]int)

	cycleStart := 0
	cycleLength := 0
	for i := 0; i < 1000000000; i++ {
		north()
		west()
		south()
		east()
		state := strings.Join(lines, "")
		if idx, ok := cache[state]; ok {
			cycleStart = idx
			cycleLength = i - idx
			fmt.Println("cycle after", cycleStart, " len ", cycleLength, " ->", i, ":", idx)
			for {
				if i+cycleLength >= 1000000000 {
					break
				} else {
					i += cycleLength
				}
			}
		} else {
			cache[state] = i
		}
	}
	// remaining := (1000000000 - cycleStart) % cycleLength
	// fmt.Println("REM", remaining)
	// for i := 0; i < remaining; i++ {
	// 	north()
	// 	west()
	// 	south()
	// 	east()
	// }

	fmt.Println("-E-----------------------")
	for _, v := range lines {
		fmt.Println(v)
	}

	sum := 0
	for idx, line := range lines {
		for _, v := range line {
			if v == 'O' {
				sum += len(lines) - idx
			}
		}
	}
	fmt.Println(sum)
}

func north() {
	for x := 0; x < len(lines[0]); x++ {
		free := 0
		for y := 0; y < len(lines); y++ {
			c := lines[y][x : x+1]
			if c == "#" {
				free = y + 1
			} else if c == "O" {
				if free < y {
					lines[free] = lines[free][:x] + "O" + lines[free][x+1:]
					lines[y] = lines[y][:x] + "." + lines[y][x+1:]
					free++
				} else {
					free = y + 1
				}
			}
		}
	}
}

func south() {
	for x := 0; x < len(lines[0]); x++ {
		free := len(lines) - 1
		for y := len(lines) - 1; y >= 0; y-- {
			c := lines[y][x : x+1]
			if c == "#" {
				free = y - 1
			} else if c == "O" {
				if free > y {
					lines[free] = lines[free][:x] + "O" + lines[free][x+1:]
					lines[y] = lines[y][:x] + "." + lines[y][x+1:]
					free--
				} else {
					free = y - 1
				}
			}
		}
	}
}

func west() {
	for y := 0; y < len(lines); y++ {
		free := 0
		for x := 0; x < len(lines[0]); x++ {
			c := lines[y][x : x+1]
			if c == "#" {
				free = x + 1
			} else if c == "O" {
				if free < x {
					lines[y] = lines[y][:free] + "O" + lines[y][free+1:]
					lines[y] = lines[y][:x] + "." + lines[y][x+1:]
					free++
				} else {
					free = x + 1
				}
			}
		}
	}
}

func east() {
	for y := 0; y < len(lines); y++ {
		free := len(lines[0]) - 1
		for x := len(lines[0]) - 1; x >= 0; x-- {
			c := lines[y][x : x+1]
			if c == "#" {
				free = x - 1
			} else if c == "O" {
				if free > x {
					lines[y] = lines[y][:free] + "O" + lines[y][free+1:]
					lines[y] = lines[y][:x] + "." + lines[y][x+1:]
					free--
				} else {
					free = x - 1
				}
			}
		}
	}
}
