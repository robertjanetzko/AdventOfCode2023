package day7

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type betData struct {
	tRank string
	bet   int
	hand  string
}

func Run() {
	const cards = "AKQJT98765432"
	const cardRank = "EDCBA98765432"

	tRank := map[string]string{
		"11111": "1",
		"1112":  "2",
		"122":   "3",
		"113":   "4",
		"23":    "5",
		"14":    "6",
		"5":     "7",
	}

	data, err := os.ReadFile("day7/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	var results []betData

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		data := strings.Split(line, " ")
		bet, _ := strconv.Atoi(data[1])

		r := ""
		t := ""
		s := strings.Split(data[0]+"Z", "")
		sort.Strings(s)
		strings.Join(s, "")
		c := 1
		for i := 1; i <= 5; i++ {
			if s[i] == s[i-1] {
				c++
			} else {
				t += fmt.Sprintf("%d", c)
				c = 1
			}
			idx := strings.Index(cards, data[0][i-1:i])
			r += cardRank[idx : idx+1]
			fmt.Println(s[i-1], cardRank[idx:idx+1])
		}

		ta := strings.Split(t, "")
		sort.Strings(ta)
		t2 := strings.Join(ta, "")
		r = tRank[t2] + "-" + r
		fmt.Println(data[0], s, t2, bet, r)

		results = append(results, betData{
			tRank: r,
			bet:   bet,
			hand:  data[0],
		})

	}

	slices.SortFunc(results, func(a, b betData) int {
		return strings.Compare(a.tRank, b.tRank)
	})

	fmt.Println(results)

	sum := 0
	for i, v := range results {
		sum += (i + 1) * v.bet
		fmt.Println(i+1, v.bet, v.hand, v.tRank)
	}
	fmt.Println(sum)
}

const cards = "AKQJT98765432"
const cardsRepl = "AKQT98765432"
const cardRank2 = "EDC0A98765432"

var tRank = map[string]string{
	"11111": "1",
	"1112":  "2",
	"122":   "3",
	"113":   "4",
	"23":    "5",
	"14":    "6",
	"5":     "7",
}

func Run2() {

	data, err := os.ReadFile("day7/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	var results []betData

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		data := strings.Split(line, " ")
		bet, _ := strconv.Atoi(data[1])

		// uniq := uniqCards(data[0])

		maxRank := ""
		for _, newHand := range expandJokers(data[0], strings.Split(cardsRepl, "")) {
			r := rank2(newHand, data[0])
			if r > maxRank {
				maxRank = r
			}
		}

		// fmt.Println(data[0], bet, maxRank, uniq)
		// fmt.Println(expandJokers(data[0], uniq))

		results = append(results, betData{
			tRank: maxRank,
			bet:   bet,
			hand:  data[0],
		})

	}

	slices.SortFunc(results, func(a, b betData) int {
		return strings.Compare(a.tRank, b.tRank)
	})

	fmt.Println(results)

	sum := 0
	for i, v := range results {
		sum += (i + 1) * v.bet
		fmt.Println(i+1, v.bet, v.hand, v.tRank)
	}
	fmt.Println(sum)
}

func rank2(hand, base string) string {
	r := ""
	t := ""
	s := strings.Split(hand+"Z", "")
	sort.Strings(s)
	strings.Join(s, "")
	c := 1
	for i := 1; i <= 5; i++ {
		if s[i] == s[i-1] {
			c++
		} else {
			t += fmt.Sprintf("%d", c)
			c = 1
		}
		idx := strings.Index(cards, base[i-1:i])
		r += cardRank2[idx : idx+1]
		fmt.Println(s[i-1], cardRank2[idx:idx+1])
	}

	ta := strings.Split(t, "")
	sort.Strings(ta)
	t2 := strings.Join(ta, "")
	r = tRank[t2] + "-" + r
	return r
}

func uniqCards(hand string) []string {
	hand = strings.ReplaceAll(hand, "J", "")
	s := strings.Split(hand, "")
	sort.Strings(s)
	return slices.Compact(s)
}

func expandJokers(hand string, cards []string) []string {
	idx := strings.Index(hand, "J")
	if idx == -1 {
		return []string{hand}
	} else {
		var result []string

		for _, c := range cards {
			newHand := hand[:idx] + c + hand[idx+1:]
			result = append(result, expandJokers(newHand, cards)...)
		}

		return result
	}
}
