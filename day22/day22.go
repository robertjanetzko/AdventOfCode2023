package day22

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coords struct {
	x, y, z int
}

type brick struct {
	a, b coords

	above []int
	below []int
}

func (b brick) Overlaps(a brick) bool {
	return a.a.x <= b.b.x && a.b.x >= b.a.x &&
		a.a.y <= b.b.y && a.b.y >= b.a.y
}

// ...AAAA......
// .....BBBB....

func MakeCoords(s string) coords {
	data := strings.Split(s, ",")
	x, _ := strconv.Atoi(data[0])
	y, _ := strconv.Atoi(data[1])
	z, _ := strconv.Atoi(data[2])
	return coords{x, y, z}
}

func MakeBrick(s string) brick {
	data := strings.Split(s, "~")
	return brick{a: MakeCoords(data[0]), b: MakeCoords(data[1])}
}

func Run() {
	data, err := os.ReadFile("day22/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines := strings.Split(string(data), "\n")
	var bricks []brick
	for _, v := range lines {
		brick := MakeBrick(v)
		fmt.Println(brick)
		bricks = append(bricks, brick)
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].a.z < bricks[j].a.z
	})

	fmt.Println("--------")
	for i := range bricks {
		minz := 0
		for j := i - 1; j >= 0; j-- {
			if bricks[j].Overlaps(bricks[i]) {
				// fmt.Println("OVERLAP", i, j)
				if bricks[j].b.z > minz {
					minz = bricks[j].b.z
				}
			}
		}
		minz++
		bricks[i].b.z = bricks[i].b.z - bricks[i].a.z + minz
		bricks[i].a.z = minz
		fmt.Println(bricks[i])
	}

	fmt.Println("--------")
	for i := range bricks {
		for j := i - 1; j >= 0; j-- {
			if bricks[i].a.z == bricks[j].b.z+1 && bricks[j].Overlaps(bricks[i]) {
				bricks[i].below = append(bricks[i].below, j)
				bricks[j].above = append(bricks[j].above, i)
			}
		}
	}
	sum := 0
	sum2 := 0
	for i, b := range bricks {
		fmt.Println(i, b, "--------------")

		stable := true
		cnt := 0
		falling := []int{i}
		desintegrated := make(map[int]bool)
		desintegrated[i] = true
		for len(falling) > 0 {
			b := bricks[falling[0]]
			falling = falling[1:]
			for _, a := range b.above {
				belowCount := 0
				for _, v := range bricks[a].below {
					if !desintegrated[v] {
						belowCount++
					}
				}
				if belowCount < 1 {
					stable = false
					falling = append(falling, a)
					if !desintegrated[a] {
						cnt++
					}
					desintegrated[a] = true
				}
			}
		}
		if stable {
			sum++
		} else {
			sum2 += cnt
		}
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

// 1,0,1~1,2,1
// 0,0,2~2,0,2
// 0,2,3~2,2,3
// 0,0,4~0,2,4
// 2,0,5~2,2,5
// 0,1,6~2,1,6
// 1,1,8~1,1,9

// .A.
// .A.
// .A.

// BB.
// ...
// ...

// ...
// ...
// CC.
