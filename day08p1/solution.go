package day08p1

import (
	"bufio"
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
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
	for i := 0; i < g.m; i++ {
		for j := 0; j < g.n; j++ {
			if j > 0 {
				g.l[i][j] = max(g.l[i][j-1], g.height(i, j-1))
			}
			if i > 0 {
				g.u[i][j] = max(g.u[i-1][j], g.height(i-1, j))
			}
		}
		for j := g.n - 2; j >= 0; j-- {
			g.r[i][j] = max(g.r[i][j+1], g.height(i, j+1))
		}
	}
	for i := g.m - 2; i >= 0; i-- {
		for j := 0; j < g.n; j++ {
			g.d[i][j] = max(g.d[i+1][j], g.height(i+1, j))
		}
	}
}

func (g *grid) countVisible() int {
	cnt := 0
	for i := 0; i < g.m; i++ {
		for j := 0; j < g.n; j++ {
			h := g.height(i, j)
			if h > g.l[i][j] || h > g.r[i][j] ||
				h > g.u[i][j] || h > g.d[i][j] {
				cnt++
			}
		}
	}
	return cnt
}

func Run(s *bufio.Scanner) {
	g := newGrid()
	for s.Scan() {
		g.append(s.Text())
	}
	g.build()
	fmt.Println(g.countVisible())
}
