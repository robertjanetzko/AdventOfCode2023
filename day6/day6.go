package day6

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	data, err := os.ReadFile("day6/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines := strings.Split(string(data), "\n")
	times := list(lines[0])
	distances := list(lines[1])
	fmt.Println("times", times)
	fmt.Println("dist", distances)

	product := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		dist := distances[i]
		count := 0
		for press := 0; press <= time; press++ {
			travel := (time - press) * press
			if travel > dist {
				count++
			}
		}
		fmt.Println(i, "C", count)
		product *= count
	}
	fmt.Println("Prod", product)
}

func list(s string) []int {
	var result []int
	re := regexp.MustCompile(" +")
	split := re.Split(s, -1)
	for _, v := range split[1:] {
		n, _ := strconv.Atoi(v)
		result = append(result, n)
	}
	return result
}
