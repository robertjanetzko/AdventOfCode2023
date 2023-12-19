package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Run() {
	data, err := os.ReadFile("day1/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	sc := bufio.NewScanner(strings.NewReader(string(data)))
	sum := 0
	for sc.Scan() {
		line := sc.Text()

		log.Println(line)

		numbers := strings.TrimFunc(line, func(r rune) bool {
			return !unicode.IsNumber(r)
		})

		firstandlast := numbers[:1] + numbers[len(numbers)-1:]
		log.Println(firstandlast)
		value, err := strconv.Atoi(firstandlast)
		if err != nil {
			log.Panicln("not a number", firstandlast)
		}
		sum += value
	}
	log.Println(sum)

}

func RunPart2() {
	data, err := os.ReadFile("day1/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	numbers := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	sc := bufio.NewScanner(strings.NewReader(string(data)))
	sum := 0
	for sc.Scan() {
		line := sc.Text()

		firstIndex := -1
		lastIndex := -1
		firstValue := 0
		lastValue := 0
		for key, value := range numbers {
			idx := strings.Index(line, key)
			if idx != -1 && (firstIndex == -1 || idx < firstIndex) {
				firstIndex = idx
				firstValue = value
			}
			idx = strings.LastIndex(line, key)
			if idx != -1 && (lastIndex == -1 || idx > lastIndex) {
				lastIndex = idx
				lastValue = value
			}
		}
		code := (firstValue * 10) + lastValue
		log.Println(line, code)
		sum += code

	}
	log.Println(sum)

}
