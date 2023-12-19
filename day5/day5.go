package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var maps map[string]*[]Mapping
var dsts map[string]string

type Mapping struct {
	Destination uint64
	Source      uint64
	Length      uint64
}

func (r *Mapping) End() uint64 {
	return r.Source + r.Length - 1
}

type Range struct {
	Start  uint64
	Length uint64
}

func (r *Range) End() uint64 {
	return r.Start + r.Length - 1
}

func Run() {
	data, err := os.ReadFile("day5/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	var seeds []uint64

	sc := bufio.NewScanner(strings.NewReader(string(data)))
	inmap := false
	var currentMap *[]Mapping
	maps = make(map[string]*[]Mapping)
	dsts = make(map[string]string)
	for sc.Scan() {
		line := sc.Text()

		if strings.HasPrefix(line, "seeds: ") {
			seeds = list(line[len("seeds: "):])
		} else if strings.HasSuffix(line, " map:") {
			types := strings.Split(line[:len(line)-len(" map:")], "-to-")
			inmap = true
			fmt.Println(line, types)
			currentMap = &[]Mapping{}
			dsts[types[0]] = types[1]
			maps[types[0]] = currentMap
		} else if inmap && line != "" {
			data := list(line)
			fmt.Println(data)
			*currentMap = append(*currentMap, Mapping{
				Destination: data[0],
				Source:      data[1],
				Length:      data[2],
			})
		}
	}
	fmt.Println(seeds)
	for k, v := range maps {
		fmt.Println(k, *v)
	}

	minL := ^uint64(0)
	for _, v := range seeds {
		ranges := mapValue(Range{Start: v, Length: 1}, "seed", "location")
		if minStart(ranges) < minL {
			minL = minStart(ranges)
		}
	}
	fmt.Println("min loc1", minL)

	minL = ^uint64(0)
	for i := 0; i < len(seeds); i += 2 {
		ranges := mapValue(Range{
			Start:  seeds[i],
			Length: seeds[i+1],
		}, "seed", "location")
		if minStart(ranges) < minL {
			minL = minStart(ranges)
		}
	}

	fmt.Println("min loc2", minL)
}

func minStart(ranges []Range) uint64 {
	ms := ^uint64(0)
	for _, v := range ranges {
		if v.Start < ms {
			ms = v.Start
		}
	}
	return ms
}

func list(s string) []uint64 {
	var result []uint64
	for _, v := range strings.Split(s, " ") {
		n, _ := strconv.ParseUint(v, 10, 64)
		result = append(result, n)
	}
	return result
}

func mapValue(r Range, src, dst string) []Range {
	ranges := []Range{r}
	d := src
	for d != dst {
		ranges, d = applyMap(ranges, d)
	}
	return ranges
}

func MakeRangeTo(start, end uint64) Range {
	return Range{Start: start, Length: end - start + 1}
}

func applyMap(ranges []Range, src string) ([]Range, string) {
	var result []Range
	for len(ranges) > 0 {
		r := ranges[0]
		ranges = ranges[1:]
		found := false
		for _, m := range *maps[src] {
			overlap := r.Start <= m.End() && r.End() >= m.Source
			if overlap {
				if r.Start < m.Source {
					ranges = append(ranges, MakeRangeTo(r.Start, m.Source-1))
				}
				if r.End() > m.End() {
					ranges = append(ranges, MakeRangeTo(m.End()+1, r.End()))
				}
				r = MakeRangeTo(max(r.Start, m.Source), min(r.End(), m.End()))
				result = append(result, Range{Start: r.Start - m.Source + m.Destination, Length: r.Length})
				found = true
				break
			}
		}
		if !found {
			result = append(result, r)
		}
	}
	return result, dsts[src]
}
