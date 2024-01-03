package day24

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type vec2 struct {
	x, y float64
}

type ray struct {
	origin, dir vec2
}

func MakeVec2(s string) vec2 {
	data := strings.Split(s, ", ")
	x, _ := strconv.ParseFloat(strings.Trim(data[0], " "), 64)
	y, _ := strconv.ParseFloat(strings.Trim(data[1], " "), 64)
	return vec2{x, y}
}

func MakeRay2(s string) ray {
	data := strings.Split(s, " @ ")
	return ray{MakeVec2(data[0]), MakeVec2(data[1])}
}

func rayIntersection(ray1, ray2 ray) (vec2, error) {
	p1, p2 := ray1.origin, ray2.origin
	d1, d2 := ray1.dir, ray2.dir

	det := d1.x*d2.y - d1.y*d2.x

	if det == 0 {
		return vec2{}, fmt.Errorf("rays are parallel or collinear, no intersection")
	}

	t := ((p2.x-p1.x)*d2.y - (p2.y-p1.y)*d2.x) / det
	s := ((p2.x-p1.x)*d1.y - (p2.y-p1.y)*d1.x) / det

	if t < 0 || s < 0 {
		return vec2{}, fmt.Errorf("intersection lies outside the direction of rays")
	}

	intersectionPoint := vec2{p1.x + t*d1.x, p1.y + t*d1.y}
	return intersectionPoint, nil
}

// func (a vec3) sub(b vec3) vec3 {
// 	return vec3{a.x - b.x, a.y - b.y, a.z - b.z}
// }

// func (a vec3) div(b vec3) vec3 {
// 	return vec3{a.x / b.x, a.y / b.y, a.z / b.z}
// }

// func (a vec3) min(b vec3) vec3 {
// 	return vec3{min(a.x, b.x), min(a.y, b.y), min(a.z, b.z)}
// }

// func (a vec3) max(b vec3) vec3 {
// 	return vec3{max(a.x, b.x), max(a.y, b.y), max(a.z, b.z)}
// }

// func (r ray) atX2(x float64) vec3 {
// 	m := r.dir.y / r.dir.x
// 	y := r.origin.y + m*(x-r.origin.x)
// 	return vec3{x, y, 0}
// }

// func (r ray) intersectAABB(boxMin, boxMax vec3) bool {
// 	tMin := boxMin.sub(r.origin).div(r.dir)
// 	tMax := boxMax.sub(r.origin).div(r.dir)
// 	t1 := tMin.min(tMax)
// 	t2 := tMin.max(tMax)
// 	tNear := max(t1.x, t1.y)
// 	tFar := min(t2.x, t2.y)
// 	return tNear <= tFar
// }

func Run() {
	data, err := os.ReadFile("day24/data")
	if err != nil {
		log.Panicln("could not open file", err)
	}

	// start := 7.0
	// end := 27.0
	start := 200000000000000.0
	end := 400000000000000.0

	var rays []ray

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		r := MakeRay2(line)
		fmt.Println(r)

		// if r.intersectAABB(vec3{start, start, 0}, vec3{end, end, 0}) {
		// 	fmt.Println("INT")
		rays = append(rays, r)
		// }
	}

	fmt.Println(".....")
	sum := 0
	for i := 0; i < len(rays); i++ {
		for j := i + 1; j < len(rays); j++ {

			fmt.Println(rays[i], "<>", rays[j])
			p, err := rayIntersection(rays[i], rays[j])
			fmt.Println(p, err)
			if err != nil {
				continue
			}

			if p.x >= start && p.x <= end && p.y >= start && p.y <= end {
				sum++
			}
		}
	}

	fmt.Println(sum)
}
