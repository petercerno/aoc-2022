package day23

import (
	"bufio"
	"fmt"
)

type coord [2]int

func (p coord) add(q coord) coord {
	return coord{p[0] + q[0], p[1] + q[1]}
}

type elf struct {
	p coord
	q coord
}

type grove struct {
	e    []*elf
	m    map[coord]*elf
	d    [4][3]coord
	u, v coord
}

func newGrove() *grove {
	return &grove{
		e: make([]*elf, 0),
		m: make(map[coord]*elf),
		d: [4][3]coord{
			{{-1, 0}, {-1, 1}, {-1, -1}},
			{{1, 0}, {1, 1}, {1, -1}},
			{{0, -1}, {-1, -1}, {1, -1}},
			{{0, 1}, {-1, 1}, {1, 1}},
		},
	}
}

func (e *elf) alone(g *grove) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if _, ok := g.m[e.p.add(coord{i, j})]; ok {
				return false
			}
		}
	}
	return true
}

func (e *elf) propose(g *grove, k int) {
	e.q = e.p
	if e.alone(g) {
		return
	}

	for i := 0; i < 4; i++ {
		o := (i + k) % 4
		empty := true
		for j := 0; j < 3; j++ {
			if _, ok := g.m[e.p.add(g.d[o][j])]; ok {
				empty = false
				break
			}
		}
		if empty {
			e.q = e.p.add(g.d[o][0])
			break
		}
	}
}

func (g *grove) move(k int) bool {
	m := false
	c := make(map[coord]int)
	for _, e := range g.e {
		e.propose(g, k)
		c[e.q]++
	}
	for _, e := range g.e {
		if e.q != e.p && c[e.q] == 1 {
			delete(g.m, e.p)
			e.p = e.q
			g.m[e.p] = e
			m = true
		}
	}
	return m
}

func (g *grove) bounds() {
	g.u = coord{1e9, 1e9}
	g.v = coord{-1e9, -1e9}
	for _, e := range g.e {
		for i := 0; i < 2; i++ {
			if e.p[i] < g.u[i] {
				g.u[i] = e.p[i]
			}
			if e.p[i] > g.v[i] {
				g.v[i] = e.p[i]
			}
		}
	}
}

func (g *grove) area() int {
	return (g.v[0] - g.u[0] + 1) * (g.v[1] - g.u[1] + 1)
}

func RunPart1(g *grove) {
	for i := 0; i < 10; i++ {
		g.move(i)
	}
	g.bounds()
	fmt.Println(g.area() - len(g.e))
}

func RunPart2(g *grove) {
	i := 0
	for {
		if !g.move(i) {
			break
		}
		i++
	}
	fmt.Println(i + 1)
}

func Run(s *bufio.Scanner, part int) {
	g := newGrove()
	i := 0
	for s.Scan() {
		l := s.Text()
		for j := 0; j < len(l); j++ {
			if l[j] == '#' {
				e := &elf{p: coord{i, j}}
				g.e = append(g.e, e)
				g.m[e.p] = e
			}
		}
		i++
	}
	switch part {
	case 1:
		RunPart1(g)
	case 2:
		RunPart2(g)
	}
}
