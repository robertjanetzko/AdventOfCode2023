package day17

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"slices"
	"strconv"
	"strings"

	"github.com/x1m3/priorityQueue"
)

type coords struct {
	x, y int
}
type step struct {
	x, y, staight int
}

var heatlevels [][]uint64
var end coords

func Run() {
	data, err := os.ReadFile("day17/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines := strings.Split(string(data), "\n")

	for _, v := range lines {
		heatlevels = append(heatlevels, list(v))
	}
	end = coords{x: len(heatlevels) - 1, y: len(heatlevels) - 1}

	for _, v := range heatlevels {
		fmt.Println(v)
	}

	// a := follow(coords{-1, 0}, coords{1, 0}, 0, []step{})
	// b := follow(coords{0, -1}, coords{0, 1}, 0, []step{})

	// fmt.Println(min(a, b) - heatlevels[0][0])

	dijkstra()
}

func list(s string) []uint64 {
	var result []uint64
	for _, v := range strings.Split(s, "") {
		n, _ := strconv.ParseUint(v, 10, 64)
		result = append(result, n)
	}
	return result
}

var cache = map[int]uint64{}

func follow(curPos, curDir coords, straight int, closed []step) uint64 {
	// fmt.Println(curPos, curDir, straight, "HEAT:", heat)
	k := cacheKey(curPos, curDir, straight)
	curStep := step{curPos.x, curPos.y, straight}

	curPos = coords{curPos.x + curDir.x, curPos.y + curDir.y}
	if curPos.x < 0 || curPos.y < 0 || curPos.x >= len(heatlevels[0]) || curPos.y >= len(heatlevels) {
		return ^uint64(0)
	}

	// fmt.Println(curPos, curDir, k)
	// if v, ok := cache[k]; ok {
	// 	return v
	// }

	if curPos == end {
		fmt.Println("FOUND", closed)
		for y := 0; y < len(heatlevels); y++ {
			for x := 0; x < len(heatlevels[0]); x++ {
				if slices.ContainsFunc(closed, func(s step) bool { return s.x == x && s.y == y }) {
					fmt.Print("X")
				} else {
					fmt.Print(".")
					// fmt.Print(heatlevels[y][x])
				}
			}
			fmt.Println()
		}
		return heatlevels[curPos.y][curPos.x]
	}

	if slices.Contains(closed, curStep) {
		return ^uint64(0)
	}

	newClosed := append(closed, curStep)

	a := ^uint64(0)
	b := ^uint64(0)
	c := ^uint64(0)

	if straight <= 2 {
		a = follow(curPos, curDir, straight+1, newClosed)
	}
	if curDir.x == 0 {
		b = follow(curPos, coords{1, 0}, 1, newClosed)
		c = follow(curPos, coords{-1, 0}, 1, newClosed)
	} else {
		b = follow(curPos, coords{0, 1}, 1, newClosed)
		c = follow(curPos, coords{0, -1}, 1, newClosed)
	}
	r := min(a, b, c) + heatlevels[curPos.y][curPos.x]

	if v, ok := cache[k]; ok {
		cache[k] = min(r, v)
	} else {
		cache[k] = r
	}

	return cache[k]
}

func cacheKey(curPos, curDir coords, straight int) int {
	return (((curPos.x*1000+curPos.y)*1000+(5+curDir.x))*10+(5+curDir.y))*10 + straight
}

func (n *node) key() int {
	return (((n.pos.x*1000+n.pos.y)*1000+(5+n.dir.x))*10+(5+n.dir.y))*100 + n.straight
}

var nodes = map[int]node{}
var closed = map[int]bool{}
var dists = map[int]uint64{}
var prev = map[int]node{}

type node struct {
	pos      coords
	dir      coords
	straight int
}

func (i node) HigherPriorityThan(other priorityQueue.Interface) bool {
	j := other.(node)
	d1 := ^uint64(0)
	if d, ok := dists[i.key()]; ok {
		d1 = d
	}
	d2 := ^uint64(0)
	if d, ok := dists[j.key()]; ok {
		d2 = d
	}
	return d1 < d2
}

func minDistNode() (node, bool) {
	minV := ^uint64(0)
	var minK node
	for k, v := range dists {
		if !closed[k] && v < minV {
			minV = v
			minK = nodes[k]
		}
	}
	return minK, minV < ^uint64(0)
}

var pq *priorityQueue.Queue

func dijkstra() {
	f, err := os.Create("day17.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	pq = priorityQueue.New()

	count := 0

	start := node{coords{0, 0}, coords{0, 0}, 0}
	end := coords{len(heatlevels[0]) - 1, len(heatlevels) - 1}
	dists[start.key()] = 0
	nodes[start.key()] = start
	pq.Push(start)
	for {
		count++
		// if count > 10000 {
		// 	break
		// }

		n := pq.Pop()
		if n == nil {
			break
		}
		next := n.(node)
		if next.pos == end {
			fmt.Println("FOUND!!!")
			break
		}
		closed[next.key()] = true
		for _, v := range neighbors2(next) {
			if cl, ok := closed[v.key()]; !ok || !cl {
				updateDistance(next, v)
			}
		}
	}

	minL := ^uint64(0)
	for _, v := range nodes {
		if v.pos == end {
			result := v

			fmt.Println(dists[result.key()])
			if dists[result.key()] < minL {
				minL = dists[result.key()]
			}

			path := []node{result}
			for {
				if p, ok := prev[result.key()]; ok {
					result = p
					path = append(path, p)
				} else {
					break
				}
			}
			slices.Reverse(path)
			fmt.Println(path)

			for y := 0; y < len(heatlevels); y++ {
				for x := 0; x < len(heatlevels[0]); x++ {
					if slices.ContainsFunc(path, func(s node) bool { return s.pos.x == x && s.pos.y == y }) {
						fmt.Print("X")
					} else {
						fmt.Print(".")
						// fmt.Print(heatlevels[y][x])
					}
				}
				fmt.Println()
			}
		}
	}
	fmt.Println("MIN", minL)
}

func neighbors(n node) []node {
	r := []node{}

	if n.dir == (coords{0, 0}) {
		r = append(r, n.move(1, 0))
		r = append(r, n.move(-1, 0))
		r = append(r, n.move(0, 1))
		r = append(r, n.move(0, -1))
	} else {
		if n.straight < 3 {
			r = append(r, n.move(n.dir.x, n.dir.y))
		}
		if n.dir.x == 0 {
			r = append(r, n.move(1, 0))
			r = append(r, n.move(-1, 0))
		} else {
			r = append(r, n.move(0, 1))
			r = append(r, n.move(0, -1))
		}
	}

	return r
}

func neighbors2(n node) []node {
	r := []node{}

	if n.dir == (coords{0, 0}) {
		r = append(r, n.move(1, 0))
		r = append(r, n.move(-1, 0))
		r = append(r, n.move(0, 1))
		r = append(r, n.move(0, -1))
	} else {
		if n.straight < 10 {
			r = append(r, n.move(n.dir.x, n.dir.y))
		}
		if n.straight >= 4 {
			if n.dir.x == 0 {
				r = append(r, n.move(1, 0))
				r = append(r, n.move(-1, 0))
			} else {
				r = append(r, n.move(0, 1))
				r = append(r, n.move(0, -1))
			}
		}
	}

	return r
}

func updateDistance(u, v node) {
	if v.pos.x < 0 || v.pos.y < 0 || v.pos.x >= len(heatlevels[0]) || v.pos.y >= len(heatlevels) {
		return
	}

	nd := dists[u.key()] + heatlevels[v.pos.y][v.pos.x]
	od, ok := dists[v.key()]
	if !ok {
		od = ^uint64(0)
	}
	if nd < od {
		dists[v.key()] = nd
		nodes[v.key()] = v
		prev[v.key()] = u
		pq.Push(v)
	}
}

func (c *coords) move(dir coords) coords {
	return coords{c.x + dir.x, c.y + dir.y}
}

func (n *node) move(dx, dy int) node {
	dir := coords{dx, dy}
	straight := 1
	if dir == n.dir {
		straight = n.straight + 1
	}
	return node{n.pos.move(dir), dir, straight}
}
