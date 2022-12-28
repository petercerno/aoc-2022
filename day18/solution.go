package day18

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type cube [3]int

func parse(s string) cube {
	p := strings.Split(s, ",")
	x, _ := strconv.Atoi(p[0])
	y, _ := strconv.Atoi(p[1])
	z, _ := strconv.Atoi(p[2])
	return cube{x, y, z}
}

type grid struct {
	m map[cube]bool
	l cube
	u cube
	a int
}

func newGrid() *grid {
	return &grid{
		m: make(map[cube]bool),
		l: [3]int{+1e6, +1e6, +1e6},
		u: [3]int{-1e6, -1e6, -1e6},
	}
}

func (g *grid) add(c cube) {
	g.m[c] = true
	g.a += 6
	for i := 0; i < 3; i++ {
		if c[i] < g.l[i] {
			g.l[i] = c[i]
		}
		if c[i] > g.u[i] {
			g.u[i] = c[i]
		}
		d := c
		d[i] = c[i] + 1
		if _, ok := g.m[d]; ok {
			g.a -= 2
		}
		d[i] = c[i] - 1
		if _, ok := g.m[d]; ok {
			g.a -= 2
		}
	}
}

func (g *grid) outsideArea() int {
	a := 0
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			a += 2 * (g.u[i] - g.l[i] + 3) * (g.u[j] - g.l[j] + 3)
		}
	}
	return a
}

func (g *grid) outsideGrid() *grid {
	o := newGrid()
	s := make([]cube, 0, 10000)
	c := cube{g.l[0] - 1, g.l[1] - 1, g.l[2] - 1}
	o.add(c)
	s = append(s, c)
	for len(s) > 0 {
		c = s[len(s)-1]
		s = s[:len(s)-1]
		for i := 0; i < 3; i++ {
			for j := -1; j <= +1; j += 2 {
				d := c
				d[i] += j
				if d[i] < g.l[i]-1 || d[i] > g.u[i]+1 {
					continue
				}

				if _, ok := g.m[d]; ok {
					continue
				}

				if _, ok := o.m[d]; ok {
					continue
				}

				o.add(d)
				s = append(s, d)
			}
		}
	}
	return o
}

func Run(s *bufio.Scanner) {
	g := newGrid()
	for s.Scan() {
		g.add(parse(s.Text()))
	}
	fmt.Println(g.a)
	fmt.Println(g.outsideGrid().a - g.outsideArea())
}
