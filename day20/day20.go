package day20

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type module struct {
	name    string
	mode    string
	targets []string

	state  bool
	states map[string]bool
}

type signal struct {
	src, dst string
	state    bool
}

func MakeModule(s string) module {
	data := strings.Split(s, " -> ")

	mode := ""
	if strings.HasPrefix(s, "%") {
		mode = "%"
	} else if strings.HasPrefix(s, "&") {
		mode = "&"
	}

	return module{
		name:    strings.TrimLeft(data[0], "%&"),
		mode:    mode,
		targets: strings.Split(data[1], ", "),
		states:  map[string]bool{},
	}
}

var modules = make(map[string]*module)
var signals []signal

var low, high int

func Run() {
	data, err := os.ReadFile("day20/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	sc := bufio.NewScanner(strings.NewReader(string(data)))
	for sc.Scan() {
		line := sc.Text()
		module := MakeModule(line)
		modules[module.name] = &module
	}
	// update inputs
	for _, module := range modules {
		for _, target := range module.targets {
			if m, ok := modules[target]; ok && m.mode == "&" {
				modules[target].states[module.name] = false
			}
		}
	}
	for _, module := range modules {
		fmt.Println(module)
	}

	cnt := 0
	done := false
	doneM := make(map[string]int)
	for k := range modules["lx"].states {
		doneM[k] = -1
	}
	for {
		cnt++
		emit(signal{"button", "broadcaster", false})
		for len(signals) > 0 {
			sig := signals[0]
			signals = signals[1:]

			if sig.dst == "lx" && sig.state {
				fmt.Println("DONE", sig.src, cnt)
				if doneM[sig.src] == -1 {
					doneM[sig.src] = cnt
				}
				done = true
				for _, v := range doneM {
					if v == -1 {
						done = false
						break
					}
				}
				break
			}

			if m, ok := modules[sig.dst]; ok {
				m.process(sig)
			}
		}
		if done {
			break
		}
		if cnt%100000 == 0 {
			fmt.Println(cnt)
		}
	}
	fmt.Println("cnt", cnt)
	fmt.Println(low * high)

	var cycles []int
	for _, v := range doneM {
		cycles = append(cycles, v)
	}
	fmt.Println(LCM(1, 1, cycles...))
}

func emit(sig signal) {
	// fmt.Println("SIG", sig)
	signals = append(signals, sig)
	if sig.state {
		high++
	} else {
		low++
	}
}

func (m *module) process(sig signal) {
	switch m.mode {
	case "":
		for _, t := range m.targets {
			emit(signal{m.name, t, sig.state})
		}
	case "%":
		if !sig.state {
			m.state = !m.state
			for _, t := range m.targets {
				emit(signal{m.name, t, m.state})
			}
		}
	case "&":
		m.states[sig.src] = sig.state
		r := true
		for _, v := range m.states {
			r = r && v
		}
		for _, t := range m.targets {
			emit(signal{m.name, t, !r})
		}
	}
}

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
