package day8

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	left  string
	right string
}

func Run() {
	data, err := os.ReadFile("day8/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines := strings.Split(string(data), "\n")

	instuctions := lines[0]
	fmt.Println(len(instuctions), instuctions)
	directions := make(map[string]node)
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		directions[line[0:3]] = node{
			left:  line[7:10],
			right: line[12:15],
		}
	}
	step := 0
	curNode := "AAA"
	for {
		if curNode == "ZZZ" {
			break
		}
		i := step % len(instuctions)
		fmt.Println(step, i, curNode, instuctions[i:i+1])
		if instuctions[i:i+1] == "L" {
			curNode = directions[curNode].left
		} else {
			curNode = directions[curNode].right
		}
		step++
	}
	fmt.Println("STEPS", step)
}

func Run2() {
	data, err := os.ReadFile("day8/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	lines := strings.Split(string(data), "\n")

	instuctions := lines[0]
	fmt.Println(len(instuctions), instuctions)
	directions := make(map[string]node)
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		directions[line[0:3]] = node{
			left:  line[7:10],
			right: line[12:15],
		}
		fmt.Println(line[0:3], line[7:10], line[12:15])
	}
	var curNodes []string

	for n := range directions {
		if strings.HasSuffix(n, "A") {
			curNodes = append(curNodes, n)
		}
	}
	fmt.Println(curNodes)

	var steps []int
	for _, v := range curNodes {
		step := 0
		curNode := v
		for {
			if strings.HasSuffix(curNode, "Z") {
				break
			}
			i := step % len(instuctions)
			// fmt.Println(step, i, curNode, instuctions[i:i+1])
			if instuctions[i:i+1] == "L" {
				curNode = directions[curNode].left
			} else {
				curNode = directions[curNode].right
			}
			step++
		}
		steps = append(steps, step)
		fmt.Println("D", step%len(instuctions))
	}
	fmt.Println("STEPS", steps, LCM(1, 1, steps...))
}

// func atEnd(curNodes []string) bool {
// 	for _, v := range curNodes {
// 		if !strings.HasSuffix(v, "Z") {
// 			return false
// 		}
// 	}
// 	return true
// }

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
