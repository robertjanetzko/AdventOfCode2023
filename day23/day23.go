package day23

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/x1m3/priorityQueue"
)

type coords struct {
	x, y int
}

type node map[coords]int

type direction struct {
	at, from coords
}

var lines []string
var graph = make(map[coords]node)
var end coords

func Run() {
	data, err := os.ReadFile("day23/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines = strings.Split(string(data), "\n")

	var open []direction
	closed := make(map[coords]bool)

	end = coords{strings.Index(lines[len(lines)-1], "."), len(lines) - 1}
	fmt.Println("END", end)
	start := coords{1, 0}
	closed[start] = true
	closed[end] = true
	graph[start] = make(node)
	open = append(open, direction{coords{1, 1}, start})

	for len(open) > 0 {
		dir := open[0]
		open = open[1:]
		origin := dir.from
		length := 0
		for {
			length++
			// fmt.Println(" DIR:", dir)
			opts := neighbors(dir)
			// fmt.Println("--->", len(opts), " :: ", opts)
			if len(opts) == 1 {
				dir = opts[0]
				continue
			}
			if len(opts) == 0 {
				break
			}

			graph[origin][dir.at] = length

			if !closed[dir.at] {
				closed[dir.at] = true
				graph[dir.at] = make(node)
				open = append(open, opts...)
				open = append(open, direction{dir.from, dir.at})
			}

			break
		}
	}

	for k, v := range graph {
		fmt.Println(k, "->", v)
	}

	// // DIJKSTRA
	// pq := priorityQueue.New()
	// prevs := make(map[coords]coords)
	// for k := range closed {
	// 	distances[k] = -1
	// 	closed[k] = false
	// }
	// distances[start] = 0

	// pq.Push(start)

	// for {
	// 	n := pq.Pop()
	// 	if n == nil {
	// 		break
	// 	}

	// 	next := n.(coords)
	// 	closed[next] = true
	// 	fmt.Println("CHECK", next)
	// 	if next == end {
	// 		fmt.Println("FOUND!!!")
	// 		break
	// 	}

	// 	for k, v := range graph[next] {
	// 		if !closed[k] {
	// 			fmt.Println("U:", next, "V:", k, " -> ", v)
	// 			alt := distances[next] + v
	// 			if alt > distances[k] {
	// 				fmt.Println("-->", k, alt, "  ---  ", next)
	// 				distances[k] = alt
	// 				prevs[k] = next
	// 				pq.Push(k)
	// 			}
	// 		}
	// 	}
	// }

	// fmt.Println(distances[end] - 1)

	// sum := 0
	// path := []coords{end}
	// u := end
	// for prevs[u] != (coords{0, 0}) {
	// 	sum += graph[prevs[u]][u]
	// 	u = prevs[u]
	// 	path = append(path, u)
	// }
	// slices.Reverse(path)
	// fmt.Println(sum, " . ", path)

	fmt.Println(follow(start, []coords{}, 0))
}

var distances = make(map[coords]int)

func neighbors(dir direction) []direction {
	if dir.at == end {
		return []direction{{dir.at, dir.at}, {dir.at, dir.at}}
	}

	options := []coords{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	// switch lines[dir.at.y][dir.at.x : dir.at.x+1] {
	// case "^":
	// 	options = []coords{{0, -1}}
	// case ">":
	// 	options = []coords{{1, 0}}
	// case "v":
	// 	options = []coords{{0, 1}}
	// case "<":
	// 	options = []coords{{-1, 0}}
	// }
	// fmt.Println(dir.at, lines[dir.at.y][dir.at.x:dir.at.x+1], options)
	var r []direction
	for _, opt := range options {
		pos := coords{dir.at.x + opt.x, dir.at.y + opt.y}
		if pos.y < 0 || pos.y >= len(lines) {
			continue
		}
		if lines[pos.y][pos.x:pos.x+1] == "#" {
			continue
		}
		if pos == dir.from {
			continue
		}
		r = append(r, direction{pos, dir.at})
	}
	return r
}

func (i coords) HigherPriorityThan(other priorityQueue.Interface) bool {
	j := other.(coords)

	return distances[i] > distances[j]
}

func follow(start coords, closed []coords, len int) int {
	if start == end {
		return len
	}
	maxlen := 0
	for k, v := range graph[start] {
		if slices.Contains(closed, k) {
			continue
		}
		l := follow(k, append(closed, start), v+len)
		if l > maxlen {
			maxlen = l
		}
	}
	return maxlen
}
