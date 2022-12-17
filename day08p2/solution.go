package day08p2

import (
	"bufio"
	"fmt"
)

type dist struct {
	s [][2]int
}

func (d *dist) empty() bool {
	return d.s == nil || len(d.s) == 0
}

func (d *dist) top() [2]int {
	return d.s[len(d.s)-1]
}

func (d *dist) pop() {
	d.s = d.s[:len(d.s)-1]
}

func (d *dist) push(i, h int) int {
	if d.empty() {
		d.s = append(d.s, [2]int{i, h})
		return 0
	}

	for !d.empty() && d.top()[1] < h {
		d.pop()
	}
	j := 0
	if !d.empty() {
		j = d.top()[0]
	}
	d.s = append(d.s, [2]int{i, h})
	return i - j
}

type grid struct {
	h          [][]byte
	m, n       int
	l, r, u, d [][]int
}

func newGrid() *grid {
	return &grid{
		h: make([][]byte, 0),
	}
}

func (g *grid) new2D() [][]int {
	a := make([][]int, g.m)
	for i := 0; i < g.m; i++ {
		a[i] = make([]int, g.n)
	}
	return a
}

func (g *grid) append(l string) {
	g.h = append(g.h, []byte(l))
	g.n = len(l)
	g.m++
}

func (g *grid) height(i, j int) int {
	return int(g.h[i][j]-'0') + 1
}

func (g *grid) build() {
	g.l = g.new2D()
	g.r = g.new2D()
	g.u = g.new2D()
	g.d = g.new2D()
	ld := make([]dist, g.m)
	rd := make([]dist, g.m)
	ud := make([]dist, g.n)
	dd := make([]dist, g.n)
	for i := 0; i < g.m; i++ {
		for j := 0; j < g.n; j++ {
			h := g.height(i, j)
			g.l[i][j] = ld[i].push(j, h)
			g.u[i][j] = ud[j].push(i, h)
		}
		for j := g.n - 1; j >= 0; j-- {
			h := g.height(i, j)
			g.r[i][j] = rd[i].push(g.n-1-j, h)
		}
	}
	for i := g.m - 1; i >= 0; i-- {
		for j := 0; j < g.n; j++ {
			h := g.height(i, j)
			g.d[i][j] = dd[j].push(g.m-1-i, h)
		}
	}
}

func (g *grid) topScenicScore() int {
	top := 0
	for i := 0; i < g.m; i++ {
		for j := 0; j < g.n; j++ {
			cur := g.l[i][j] * g.r[i][j] * g.u[i][j] * g.d[i][j]
			if cur > top {
				top = cur
			}
		}
	}
	return top
}

func Run(s *bufio.Scanner) {
	g := newGrid()
	for s.Scan() {
		g.append(s.Text())
	}
	g.build()
	fmt.Println(g.topScenicScore())
}
