package day4

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	data, err := os.ReadFile("day4/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	sum := 0
	sum2 := 0

	counts := make(map[int]int)

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		card := strings.Split(line, ": ")
		id, _ := strconv.Atoi(strings.Trim(strings.Split(card[0], "d")[1], " "))
		numbers := strings.Split(card[1], " | ")
		win := parseList(numbers[0])
		my := parseList(numbers[1])
		count := 0
		for _, v := range my {
			if slices.Contains(win, v) {
				count++
			}
		}
		if count > 0 {
			sum += 1 << (count - 1)
		}
		fmt.Println(id, count)
		counts[id] = counts[id] + 1
		for i := id + 1; i <= id+count; i++ {
			if i <= len(lines) {
				counts[i] = counts[i] + counts[id]
			}
		}
		sum2 += counts[id]
	}
	fmt.Println(sum)
	fmt.Println(counts)
	fmt.Println(sum2)
}

func parseList(list string) []string {
	var result []string
	for i := 0; i < len(list); i += 3 {
		result = append(result, list[i:i+2])
	}
	return result
}
