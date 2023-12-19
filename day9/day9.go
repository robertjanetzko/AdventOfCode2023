package day9

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	data, err := os.ReadFile("day9/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	sum := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		numbers := list(line)
		fmt.Println(numbers)
		var diffs [][]int
		diffs = append(diffs, numbers)
		for {
			var zero bool
			numbers, zero = calcDiffs(numbers)
			fmt.Println(numbers, zero)
			diffs = append(diffs, numbers)
			if zero {
				break
			}
		}
		fmt.Println(line, "DDD", diffs)
		diffs[len(diffs)-1] = append(diffs[len(diffs)-1], 0)
		for i := len(diffs) - 2; i >= 0; i-- {
			diffs[i] = append(diffs[i], diffs[i][len(diffs[i])-1]+diffs[i+1][len(diffs[i+1])-1])
			fmt.Println(diffs[i])
		}
		fmt.Println("----", diffs[0][len(diffs[0])-1])
		sum += diffs[0][len(diffs[0])-1]
	}
	fmt.Println("SUM", sum)
}

func Run2() {
	data, err := os.ReadFile("day9/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	sum := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		numbers := list(line)
		fmt.Println(numbers)
		var diffs [][]int
		diffs = append(diffs, numbers)
		for {
			var zero bool
			numbers, zero = calcDiffs(numbers)
			fmt.Println(numbers, zero)
			diffs = append(diffs, numbers)
			if zero {
				break
			}
		}
		fmt.Println(line, "DDD", diffs)
		diffs[len(diffs)-1] = append([]int{0}, diffs[len(diffs)-1]...)
		for i := len(diffs) - 2; i >= 0; i-- {
			diffs[i] = append([]int{diffs[i][0] - diffs[i+1][0]}, diffs[i]...)
			fmt.Println(diffs[i])
		}
		fmt.Println("----", diffs[0][0])
		sum += diffs[0][0]
	}
	fmt.Println("SUM", sum)
}

func list(s string) []int {
	var result []int
	re := regexp.MustCompile(" +")
	split := re.Split(s, -1)
	for _, v := range split {
		n, _ := strconv.Atoi(v)
		result = append(result, n)
	}
	return result
}

func calcDiffs(in []int) ([]int, bool) {
	var r []int
	b := true
	for i := 0; i < len(in)-1; i++ {
		d := in[i+1] - in[i]
		r = append(r, d)
		if d != 0 {
			b = false
		}
	}
	return r, b
}
