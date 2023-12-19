package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	prop   string
	op     string
	val    int
	target string
}

type workflow struct {
	name  string
	rules []rule
}

type part struct {
	x, m, a, s int
}

func (p *part) sum() int {
	return p.x + p.m + p.a + p.s
}

func (r *rule) matches(p part) bool {
	if r.op == "" {
		return true
	}

	v := 0
	switch r.prop {
	case "x":
		v = p.x
	case "m":
		v = p.m
	case "a":
		v = p.a
	case "s":
		v = p.s
	}

	if r.op == ">" {
		return v > r.val
	} else {
		return v < r.val
	}
}

func (w workflow) evaluate(p part) string {
	for _, r := range w.rules {
		if r.matches(p) {
			switch r.target {
			case "A", "R":
				return r.target
			default:
				return workflows[r.target].evaluate(p)
			}
		}
	}
	return "?"
}

func MakeRule(s string) rule {
	data := strings.Split(s, ":")
	if len(data) == 1 {
		return rule{target: s}
	}
	v, _ := strconv.Atoi(data[0][2:])
	return rule{
		prop:   data[0][0:1],
		op:     data[0][1:2],
		val:    v,
		target: data[1],
	}
}

func MakeWorkflow(s string) workflow {
	idx := strings.Index(s, "{")
	var rules []rule
	for _, v := range strings.Split(s[idx+1:len(s)-1], ",") {
		rules = append(rules, MakeRule(v))
	}
	return workflow{name: s[0:idx], rules: rules}
}

func MakePart(str string) part {
	data := strings.Split(str[1:len(str)-1], ",")
	x, _ := strconv.Atoi(data[0][2:])
	m, _ := strconv.Atoi(data[1][2:])
	a, _ := strconv.Atoi(data[2][2:])
	s, _ := strconv.Atoi(data[3][2:])
	return part{x, m, a, s}
}

type rng struct {
	min, max int
}

func (r rng) count() int {
	return max(0, r.max-r.min+1)
}

type seg struct {
	x, m, a, s rng
}

func (rn rng) apply(prop string, r rule) rng {
	if r.prop != prop {
		return rn
	}
	if r.op == ">" {
		return rng{
			min: max(rn.min, r.val+1),
			max: rn.max,
		}
	} else {
		return rng{
			min: rn.min,
			max: min(rn.max, r.val-1),
		}
	}
}

func (s seg) apply(r rule) seg {
	return seg{
		x: s.x.apply("x", r),
		m: s.m.apply("m", r),
		a: s.a.apply("a", r),
		s: s.s.apply("s", r),
	}
}

func (s seg) count() int {
	return s.x.count() * s.m.count() * s.a.count() * s.s.count()
}

func (r *rule) inverse() rule {
	if r.op == ">" {
		return rule{prop: r.prop,
			op:     "<",
			val:    r.val + 1,
			target: r.target}
	} else {
		return rule{prop: r.prop,
			op:     ">",
			val:    r.val - 1,
			target: r.target}
	}
}

func (w workflow) segment(s seg) int {
	sum := 0
	for _, r := range w.rules {
		ns := s.apply(r)
		if r.target == "R" {
		} else if r.target == "A" {
			sum += ns.count()
		} else {
			sum += workflows[r.target].segment(ns)
		}
		s = s.apply(r.inverse())
	}
	return sum
}

var workflows = make(map[string]workflow)

func Run() {
	data, err := os.ReadFile("day19/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	sum := 0
	sc := bufio.NewScanner(strings.NewReader(string(data)))
	parseParts := false
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			parseParts = true
			continue
		}
		if !parseParts {
			w := MakeWorkflow(line)
			fmt.Println(w)
			workflows[w.name] = w
		} else {
			part := MakePart(line)
			fmt.Println(part)
			t := workflows["in"].evaluate(part)
			fmt.Println(t)
			if t == "A" {
				sum += part.sum()
			}
		}
	}

	start := seg{rng{1, 4000}, rng{1, 4000}, rng{1, 4000}, rng{1, 4000}}

	fmt.Println(sum)
	fmt.Println(workflows["in"].segment(start))
}
