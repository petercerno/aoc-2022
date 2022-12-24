package day15

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type interval [2]int

func (i interval) clip(c interval) interval {
	if c[0] == c[1] {
		return i
	}

	if i[1] <= c[0] {
		return interval{c[0], c[0]}
	}

	if i[0] >= c[1] {
		return interval{c[1], c[1]}
	}

	if c[0] > i[0] {
		i[0] = c[0]
	}
	if c[1] < i[1] {
		i[1] = c[1]
	}
	return i
}

type intervals [][2]int

func (is intervals) add(i interval) intervals {
	n := len(is)
	k := 0
	js := make(intervals, 0, n+1)
	for ; k < n && is[k][1] < i[0]; k++ {
		js = append(js, is[k])
	}
	if k == n {
		js = append(js, i)
		return js
	}

	if is[k][0] < i[0] {
		i[0] = is[k][0]
	}
	for ; k < n && is[k][0] <= i[1]; k++ {
		if is[k][1] > i[1] {
			i[1] = is[k][1]
		}
	}
	js = append(js, i)
	for ; k < n; k++ {
		js = append(js, is[k])
	}
	return js
}

func (is intervals) sum() int {
	sum := 0
	for _, i := range is {
		sum += i[1] - i[0]
	}
	return sum
}

type point [2]int

func abs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}

func dist(u, v point) int {
	return abs(u[0]-v[0]) + abs(u[1]-v[1])
}

type sensor struct {
	p point
	b point
	d int
}

func parse(s []string) []int {
	n := len(s)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i], _ = strconv.Atoi(s[i])
	}
	return t
}

func newSensor(l []string) sensor {
	p := parse(l)
	s := sensor{
		p: point{p[0], p[1]},
		b: point{p[2], p[3]},
	}
	s.d = dist(s.p, s.b)
	return s
}

func (s *sensor) scan(y int) interval {
	d := s.d - abs(s.p[1]-y)
	if d < 0 {
		return interval{s.p[0], s.p[0]}
	}

	return interval{s.p[0] - d, s.p[0] + d + 1}
}

type sensors []sensor

func (ss sensors) scan(y int, c interval) intervals {
	is := intervals{}
	for _, s := range ss {
		i := s.scan(y).clip(c)
		if i[0] < i[1] {
			is = is.add(i)
		}
	}
	return is
}

func (ss sensors) beacons(y int) int {
	m := make(map[point]bool)
	for _, s := range ss {
		if s.b[1] == y {
			m[s.b] = true
		}
	}
	return len(m)
}

func RunPart1(ss sensors, y int) {
	c := interval{}
	fmt.Println(ss.scan(y, c).sum() - ss.beacons(y))
}

func RunPart2(ss sensors, m int) {
	c := interval{0, m + 1}
	M := 10000
	d := make(chan bool, m/M+1)
	for i := 0; i <= m/M; i++ {
		go func(y0, y1 int) {
			defer func() { d <- true }()
			for y := y0; y < y1 && y <= m; y++ {
				is := ss.scan(y, c)
				if len(is) == 2 {
					x := is[0][1]
					fmt.Println(x, y)
					fmt.Println(int64(x)*4000000 + int64(y))
				}
			}
		}(M*i, M*(i+1))
	}
	for i := 0; i <= m/M; i++ {
		<-d
	}
}

func Run(s *bufio.Scanner, part, m int) {
	ss := sensors{}
	re := regexp.MustCompile(`-?[0-9]+`)
	for s.Scan() {
		l := re.FindAllString(s.Text(), -1)
		ss = append(ss, newSensor(l))
	}
	switch part {
	case 1:
		RunPart1(ss, m)
	case 2:
		RunPart2(ss, m)
	}
}
