package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Run() {
	data, err := os.ReadFile("day13/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	sc := bufio.NewScanner(strings.NewReader(string(data) + "\n\n"))
	sum := 0

	var pattern []string
	for sc.Scan() {
		line := sc.Text()
		if line != "" {
			pattern = append(pattern, line)
		} else {
			for _, v := range pattern {
				fmt.Println(v)
			}
			sum += evaluate(pattern)
			fmt.Println("-----------------")
			pattern = []string{}
		}
	}

	fmt.Println(sum)
}

const allowed = 1 // 0 for part1

func evaluate(pattern []string) int {
	for i := 1; i < len(pattern); i++ {
		diffs := 0
		for j := i - 1; j >= max(0, i-(len(pattern)-i)); j-- {
			diffs += rowMatch(pattern, i+(i-1-j), j)
		}
		if diffs == allowed {
			fmt.Println("H MIRROR", i-1, i)
			return i * 100
		}
	}

	for i := 1; i < len(pattern[0]); i++ {
		diffs := 0
		for j := i - 1; j >= max(0, i-(len(pattern[0])-i)); j-- {
			diffs += colMatch(pattern, i+(i-1-j), j)
		}
		if diffs == allowed {
			fmt.Println("V MIRROR", i-1, i)
			return i
		}
	}

	return 0
}

func colMatch(pattern []string, a, b int) int {
	sum := 0
	for i := 0; i < len(pattern); i++ {
		if pattern[i][a] != pattern[i][b] {
			sum++
		}
	}
	return sum
}

func rowMatch(pattern []string, a, b int) int {
	sum := 0
	for i := 0; i < len(pattern[0]); i++ {
		if pattern[a][i] != pattern[b][i] {
			sum++
		}
	}
	return sum
}
