package day14

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func minmax(a, b int) (int, int) {
	if a <= b {
		return a, b
	}

	return b, a
}

func parse(s []string) [][2]int {
	n := len(s)
	l := make([][2]int, n)
	for i := 0; i < n; i++ {
		p := strings.Split(s[i], ",")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])
		l[i] = [2]int{x, y}
	}
	return l
}

type cave struct {
	m    map[[2]int]bool
	p, q [2]int
}

func newCave() *cave {
	return &cave{
		m: make(map[[2]int]bool),
		p: [2]int{+1e6, +1e6},
		q: [2]int{-1e6, -1e6},
	}
}

func (c *cave) drawLine(u, v [2]int) {
	i := 0
	if u[1] == v[1] {
		i = 1
	}
	j := 1 - i
	a, b := minmax(u[j], v[j])
	c.p[i], _ = minmax(c.p[i], u[i])
	c.p[j], _ = minmax(c.p[j], a)
	_, c.q[i] = minmax(c.q[i], u[i])
	_, c.q[j] = minmax(c.q[j], b)
	p := [2]int{}
	p[i] = u[i]
	for z := a; z <= b; z++ {
		p[j] = z
		c.m[p] = true
	}
}

func (c *cave) draw(l [][2]int) {
	n := len(l)
	for i := 0; i < n-1; i++ {
		c.drawLine(l[i], l[i+1])
	}
}

func (c *cave) drop() bool {
	p := [2]int{500, 0}
	_, ok := c.m[p]
	if ok {
		return false
	}

	d := [][2]int{{0, 1}, {-1, 1}, {1, 1}}
	for !ok {
		if p[1] > c.q[1] {
			return false
		}

		for i := 0; i < len(d); i++ {
			q := [2]int{p[0] + d[i][0], p[1] + d[i][1]}
			_, ok = c.m[q]
			if !ok {
				p = q
				break
			}
		}
	}
	c.m[p] = true
	return true
}

func Run(s *bufio.Scanner, part2 bool) {
	c := newCave()
	for s.Scan() {
		c.draw(parse(strings.Split(s.Text(), " -> ")))
	}
	if part2 {
		c.drawLine(
			[2]int{500 - c.q[1] - 3, c.q[1] + 2},
			[2]int{500 + c.q[1] + 3, c.q[1] + 2})
	}
	n := 0
	for c.drop() {
		n++
	}
	fmt.Println(n)
}
