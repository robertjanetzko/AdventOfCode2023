package day18

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type step struct {
	dir   string
	count int
	color string
}

type coords struct {
	x, y int
}

const part2 = true
const new = true

func Run() {
	data, err := os.ReadFile("day18/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines := strings.Split(string(data), "\n")

	var steps []step
	for _, lines := range lines {
		data := strings.Split(lines, " ")
		if !part2 {
			cnt, _ := strconv.Atoi(data[1])
			steps = append(steps, step{
				dir:   data[0],
				count: cnt,
				color: data[2],
			})
		} else {
			dir := ""
			switch data[2][7:8] {
			case "0":
				dir = "R"
			case "1":
				dir = "D"
			case "2":
				dir = "L"
			case "3":
				dir = "U"
			}
			cnt, _ := strconv.ParseInt(data[2][2:7], 16, 64)
			// fmt.Println(dir, data[2], "_>", data[2][2:7], cnt)
			steps = append(steps, step{
				dir:   dir,
				count: int(cnt),
				color: data[2],
			})
		}
	}

	if !new {

		x := 0
		y := 0
		minx := 0
		miny := 0

		maxx := 0
		maxy := 0
		for _, s := range steps {
			switch s.dir {
			case "R":
				x += s.count
			case "L":
				x -= s.count
			case "U":
				y -= s.count
			case "D":
				y += s.count
			}
			fmt.Println(x, y)
			minx = min(minx, x)
			miny = min(miny, y)
			maxx = max(maxx, x)
			maxy = max(maxy, y)
		}

		w := maxx - minx + 3
		h := maxy - miny + 3

		fmt.Println("SIZE:", w, h)

		plan := make([][]string, h)
		for y := range plan {
			plan[y] = make([]string, w)
			for x := 0; x < w; x++ {
				plan[y][x] = "."
			}
		}

		x = -minx + 1
		y = -miny + 1

		fmt.Println("START", x, y)

		for idx, s := range steps {
			// ox := x
			// oy := y
			prev := steps[(len(steps)+idx-1)%len(steps)].dir
			switch s.dir {
			case "R":
				if prev == "U" {
					plan[y][x] = "F"
				} else {
					plan[y][x] = "L"
				}
				for j := 1; j < s.count; j++ {
					plan[y][x+j] = "-"
				}
				x += s.count
			case "L":
				if prev == "U" {
					plan[y][x] = "7"
				} else {
					plan[y][x] = "J"
				}
				for j := 1; j < s.count; j++ {
					plan[y][x-j] = "-"
				}
				x -= s.count
			case "U":
				if prev == "R" {
					plan[y][x] = "J"
				} else {
					plan[y][x] = "L"
				}
				for j := 1; j < s.count; j++ {
					plan[y-j][x] = "|"
				}
				y -= s.count
			case "D":
				if prev == "R" {
					plan[y][x] = "7"
				} else {
					plan[y][x] = "F"
				}
				for j := 1; j < s.count; j++ {
					plan[y+j][x] = "|"
				}
				y += s.count
			}
			// fmt.Println(x, y, " < ", ox, oy)
			// for sy := min(y, oy); sy <= max(y, oy); sy++ {
			// 	for sx := min(x, ox); sx <= max(x, ox); sx++ {
			// 		plan[sy][sx] = "#"
			// 	}
			// }
		}

		sum := 0
		for y := 0; y < h; y++ {
			cnt := 0
			for x := 0; x < w; x++ {
				p := plan[y][x]
				if p == "|" || p == "L" || p == "J" {
					cnt++
				}
				if p == "." {
					if cnt%2 == 1 {
						sum++
						plan[y][x] = "#"
					}
				} else {
					sum++
				}
			}
		}

		// for y := 0; y < h; y++ {
		// 	for x := 0; x < w; x++ {
		// 		fmt.Print(plan[y][x])
		// 	}
		// 	fmt.Println()
		// }
		fmt.Println("AREA:", sum)

	} else {

		var shape []coords

		x := 0
		y := 0
		for idx, s := range steps {
			next := steps[(idx+1)%len(steps)].dir
			switch s.dir {
			case "R":
				x += s.count
				if next == "U" {
					// plan[y][x] = "J"
					shape = append(shape, coords{x, y})
				} else {
					// plan[y][x] = "7"
					shape = append(shape, coords{x + 1, y})
				}
			case "L":
				x -= s.count
				if next == "U" {
					// plan[y][x] = "L"
					shape = append(shape, coords{x, y + 1})
				} else {
					// plan[y][x] = "F"
					shape = append(shape, coords{x + 1, y + 1})
				}
			case "U":
				y -= s.count
				if next == "R" {
					// plan[y][x] = "F"
					shape = append(shape, coords{x, y})
				} else {
					// plan[y][x] = "7"
					shape = append(shape, coords{x, y + 1})
				}
			case "D":
				y += s.count
				if next == "R" {
					// plan[y][x] = "L"
					shape = append(shape, coords{x + 1, y})
				} else {
					// plan[y][x] = "J"
					shape = append(shape, coords{x + 1, y + 1})
				}
			}
			// fmt.Println(x, y, s.dir, next, shape[len(shape)-1])
		}
		s1 := 0
		s2 := 0
		for i := 1; i < len(shape); i++ {
			s1 += shape[i].y * shape[i-1].x
			s2 += shape[i].x * shape[i-1].y
		}
		a := (s1 - s2) / 2
		// fmt.Println(shape)
		fmt.Print(a)
	}
}
