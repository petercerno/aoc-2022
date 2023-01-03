package day22p2

import (
	"bufio"
	"fmt"
)

type input struct {
	l    []string
	m, n int
}

func (inp *input) append(s string) {
	inp.l = append(inp.l, s)
	if len(s) > inp.n {
		inp.n = len(s)
	}
	inp.m++
}

type move struct {
	x, r int
}

type path []move

func parsePath(s string) path {
	p := make(path, 0, len(s))
	x := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			x = 10*x + int(s[i]-'0')
		} else {
			p = append(p, move{x, 0})
			x = 0
			switch s[i] {
			case 'L':
				p = append(p, move{0, 3})
			case 'R':
				p = append(p, move{0, 1})
			}
		}
	}
	if x > 0 {
		p = append(p, move{x, 0})
	}
	return p
}

type coord struct {
	i, j, s, f int
}

type cube struct {
	n int
	d [4][2]int
	t map[[2]int]func(i, j int) coord
}

func newCube(n int) *cube {
	return &cube{
		n,
		[4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
		map[[2]int]func(i, j int) coord{
			{0, 0}: func(i, j int) coord { return coord{i, 0, 1, 0} },
			{0, 1}: func(i, j int) coord { return coord{0, j, 2, 1} },
			{0, 2}: func(i, j int) coord { return coord{i, n - 1, 3, 2} },
			{0, 3}: func(i, j int) coord { return coord{n - 1, j, 4, 3} },

			{1, 0}: func(i, j int) coord { return coord{n - 1 - i, n - 1, 5, 2} },
			{1, 1}: func(i, j int) coord { return coord{j, n - 1, 2, 2} },
			{1, 2}: func(i, j int) coord { return coord{i, n - 1, 0, 2} },
			{1, 3}: func(i, j int) coord { return coord{n - 1 - j, n - 1, 4, 2} },

			{2, 0}: func(i, j int) coord { return coord{n - 1, i, 1, 3} },
			{2, 1}: func(i, j int) coord { return coord{0, j, 5, 1} },
			{2, 2}: func(i, j int) coord { return coord{n - 1, n - 1 - i, 3, 3} },
			{2, 3}: func(i, j int) coord { return coord{n - 1, j, 0, 3} },

			{3, 0}: func(i, j int) coord { return coord{i, 0, 0, 0} },
			{3, 1}: func(i, j int) coord { return coord{n - 1 - j, 0, 2, 0} },
			{3, 2}: func(i, j int) coord { return coord{n - 1 - i, 0, 5, 0} },
			{3, 3}: func(i, j int) coord { return coord{j, 0, 4, 0} },

			{4, 0}: func(i, j int) coord { return coord{0, n - 1 - i, 1, 1} },
			{4, 1}: func(i, j int) coord { return coord{0, j, 0, 1} },
			{4, 2}: func(i, j int) coord { return coord{0, i, 3, 1} },
			{4, 3}: func(i, j int) coord { return coord{n - 1, j, 5, 3} },

			{5, 0}: func(i, j int) coord { return coord{n - 1 - i, n - 1, 1, 2} },
			{5, 1}: func(i, j int) coord { return coord{0, j, 4, 1} },
			{5, 2}: func(i, j int) coord { return coord{n - 1 - i, 0, 3, 0} },
			{5, 3}: func(i, j int) coord { return coord{n - 1, j, 2, 3} },
		},
	}
}

func (c *cube) move(p coord) coord {
	q := p
	q.i += c.d[p.f][0]
	q.j += c.d[p.f][1]
	if q.i >= 0 && q.i < c.n && q.j >= 0 && q.j < c.n {
		return q
	}

	return c.t[[2]int{p.s, p.f}](p.i, p.j)
}

type cell struct {
	i, j, s int
}

func (p coord) cell() cell {
	return cell{p.i, p.j, p.s}
}

type value struct {
	i, j int
	t    byte
}

type board struct {
	t map[cell]value
	c *cube
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func newBoard(inp *input) *board {
	bo := &board{t: make(map[cell]value), c: newCube(gcd(inp.m, inp.n))}
	j := 0
	for ; j < inp.n; j++ {
		if inp.l[0][j] != ' ' {
			break
		}
	}

	ps := make([][3]int, 0, inp.m*inp.n)
	qs := make([]coord, 0, inp.m*inp.n)
	p0 := [3]int{0, j, 0}
	q0 := coord{}
	bo.t[q0.cell()] = value{p0[0], p0[1], inp.l[p0[0]][p0[1]]}
	ps = append(ps, p0)
	qs = append(qs, q0)
	for len(ps) > 0 {
		p0 = ps[len(ps)-1]
		ps = ps[:len(ps)-1]
		q0 = qs[len(qs)-1]
		qs = qs[:len(qs)-1]
		for i := 0; i < 4; i++ {
			f := (p0[2] + i) % 4
			p1 := [3]int{p0[0] + bo.c.d[f][0], p0[1] + bo.c.d[f][1], f}
			if p1[0] < 0 || p1[0] >= inp.m ||
				p1[1] < 0 || p1[1] >= len(inp.l[p1[0]]) ||
				inp.l[p1[0]][p1[1]] == ' ' {
				continue
			}

			q1 := q0
			q1.f = (q0.f + i) % 4
			q1 = bo.c.move(q1)
			if _, ok := bo.t[q1.cell()]; !ok {
				bo.t[q1.cell()] = value{p1[0], p1[1], inp.l[p1[0]][p1[1]]}
				ps = append(ps, p1)
				qs = append(qs, q1)
			}
		}
	}
	return bo
}

func (bo *board) move(pa path) int {
	q0 := coord{}
	for _, m := range pa {
		for i := 0; i < m.x; i++ {
			q1 := bo.c.move(q0)
			v1, ok := bo.t[q1.cell()]
			if ok && v1.t == '.' {
				q0 = q1
			}
		}
		q0.f = (q0.f + m.r) % 4
	}
	v0 := bo.t[q0.cell()]
	qm0 := coord{bo.c.n / 2, bo.c.n / 2, q0.s, q0.f}
	qm1 := bo.c.move(qm0)
	vm0 := bo.t[qm0.cell()]
	vm1 := bo.t[qm1.cell()]
	f := 0
	for ; f < 4; f++ {
		if vm1.i == vm0.i+bo.c.d[f][0] &&
			vm1.j == vm0.j+bo.c.d[f][1] {
			break
		}
	}
	return 1000*(v0.i+1) + 4*(v0.j+1) + f
}

func Run(s *bufio.Scanner) {
	inp := &input{}
	for s.Scan() {
		l := s.Text()
		if l == "" {
			break
		}

		inp.append(l)
	}
	bo := newBoard(inp)
	s.Scan()
	pa := parsePath(s.Text())
	fmt.Println(bo.move(pa))
}
