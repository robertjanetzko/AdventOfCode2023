package day15

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type stepData struct {
	step  string
	label string
	op    string
	lens  int
	hash  int
}

type lens struct {
	label string
	lens  int
}

func Run() {
	data, err := os.ReadFile("day15/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines := strings.Split(string(data), "\n")
	line := lines[0]

	var steps []stepData
	for _, v := range strings.Split(line, ",") {
		steps = append(steps, parseStep(v))
	}
	fmt.Println(steps)

	sum := 0
	for _, step := range steps {
		sum += step.hash
	}
	fmt.Println(sum)

	boxes := make([][]lens, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = []lens{}
	}

	for _, step := range steps {
		if step.op == "-" {
			boxes[step.hash] = slices.DeleteFunc(boxes[step.hash], func(l lens) bool { return l.label == step.label })
		} else {
			idx := slices.IndexFunc(boxes[step.hash], func(l lens) bool { return l.label == step.label })
			if idx == -1 {
				boxes[step.hash] = append(boxes[step.hash], lens{label: step.label, lens: step.lens})
			} else {
				boxes[step.hash][idx].lens = step.lens
			}
		}
	}

	sum = 0
	for i, box := range boxes {
		if len(box) > 0 {
			fmt.Println(i+1, box)
		}

		for j, lens := range box {
			sum += (1 + i) * (j + 1) * lens.lens
		}
	}
	fmt.Println(sum)
}

func parseStep(s string) stepData {
	idx := max(strings.Index(s, "="), strings.Index(s, "-"))
	lens, _ := strconv.Atoi(s[idx+1:])
	return stepData{
		step:  s,
		label: s[:idx],
		op:    s[idx : idx+1],
		lens:  lens,
		hash:  hash(s[:idx]),
	}
}

func hash(s string) int {
	v := 0
	for _, c := range s {
		v = ((v + int(c)) * 17) % 256
	}
	return v
}
