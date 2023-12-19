package day12

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const part2 = true

func Run() {
	data, err := os.ReadFile("day12/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	sum := 0
	lines := strings.Split(string(data), "\n")
	for idx, line := range lines {
		data := strings.Split(line, " ")
		pattern := data[0]
		segs := list(data[1])

		if part2 {
			pattern = strings.Repeat("?"+pattern, 5)[1:]
			segs = list(strings.Repeat(","+data[1], 5)[1:])
		}

		minL := -1
		for _, v := range segs {
			minL += v + 1
		}

		possible := 0
		if minL < len(pattern) {
			fmt.Println(idx, pattern, segs)
			possible = recurse(pattern, segs)
			fmt.Println("REC", possible)
			// permutations(len(pattern)-minL, len(segs)+1, []int{}, func(arr []int) {
			// 	str := buildPattern(segs, arr)
			// 	m := matches(pattern, str)
			// 	// fmt.Println(arr, str, pattern, m)
			// 	if m {
			// 		possible++
			// 	}
			// })
		} else {
			possible = 1
		}

		sum += possible
		fmt.Println(pattern, segs, len(pattern), minL, possible)
	}

	fmt.Println(sum)
}

// func buildPattern(segs, spaces []int) string {
// 	p := strings.Repeat(".", spaces[0])
// 	for i := 0; i < len(segs); i++ {
// 		if i > 0 {
// 			p += "."
// 		}
// 		p += strings.Repeat("#", segs[i])
// 		p += strings.Repeat(".", spaces[i+1])
// 	}
// 	return p
// }

// func matches(p, s string) bool {
// 	for i := 0; i < len(p); i++ {
// 		if p[i:i+1] != s[i:i+1] && p[i:i+1] != "?" {
// 			return false
// 		}
// 	}
// 	return true
// }

func list(s string) []int {
	var result []int
	split := strings.Split(s, ",")
	for _, v := range split {
		n, _ := strconv.Atoi(v)
		result = append(result, n)
	}
	return result
}

// func permutations(n, k int, arr []int, check func([]int)) {
// 	if k == 0 {
// 		check(arr)
// 		return
// 	}

// 	if n > 0 && k > 1 {
// 		for i := 0; i <= n; i++ {
// 			permutations(n-i, k-1, append(arr, i), check)
// 		}
// 	} else {
// 		permutations(0, k-1, append(arr, n), check)
// 	}
// }

func cacheKey(open string, segs []int) string {
	r := open
	for _, v := range segs {
		r += fmt.Sprintf(",%d", v)
	}
	return r
}

var cache = map[string]int{}

func recurse(open string, segs []int) int {
	key := cacheKey(open, segs)
	if v, ok := cache[key]; ok {
		return v
	}
	// fmt.Println(pattern, done, open, segs)
	if len(open) == 0 {
		if len(segs) == 0 {
			// fmt.Println("CHECK")
			return 1
		} else {
			return 0
		}
	}
	rem := len(segs) - 1
	for _, v := range segs {
		rem += v
	}
	if rem > len(open) {
		return 0
	}

	count := 0
	next := open[0:1]
	if next == "." {
		count = recurse(open[1:], segs)
	} else if len(segs) > 0 && canPlaceGroup(open, segs[0]) {
		c := 0
		if next == "?" {
			c = recurse(open[1:], segs)
		}
		if len(open) == segs[0] {
			count = recurse(open[segs[0]:], segs[1:]) + c
		} else {
			count = recurse(open[segs[0]+1:], segs[1:]) + c
		}
	} else if next == "#" {
		count = 0
	} else {
		count = recurse(open[1:], segs)
	}
	cache[key] = count
	return count
}

func canPlaceGroup(pattern string, length int) bool {
	if len(pattern) < length {
		return false
	}
	for i := 0; i < length; i++ {
		if pattern[i:i+1] == "." {
			return false
		}
	}
	if len(pattern) == length {
		return true
	}
	return pattern[length:length+1] != "#"
}
