package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {

	data, err := os.ReadFile("day2/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	maxColor := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	gamePattern := regexp.MustCompile(`^Game (\d+): (.*)`)

	sc := bufio.NewScanner(strings.NewReader(string(data)))
	sum := 0
	sum2 := 0
	for sc.Scan() {
		line := sc.Text()

		matches := gamePattern.FindStringSubmatch(line)

		id, _ := strconv.Atoi(matches[1])

		valid := true
		counts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, set := range strings.Split(matches[2], "; ") {
			for _, cubes := range strings.Split(set, ", ") {
				data := strings.Split(cubes, " ")
				count, _ := strconv.Atoi(data[0])
				if count > maxColor[data[1]] {
					valid = false
				}
				if count > counts[data[1]] {
					counts[data[1]] = count
				}
			}
		}
		fmt.Println(line, "\n  >", id, valid)
		if valid {
			sum += id
		}

		power := counts["red"] * counts["green"] * counts["blue"]
		sum2 += power
	}
	log.Println("sum 1:", sum)
	log.Println("sum 2:", sum2)
}
