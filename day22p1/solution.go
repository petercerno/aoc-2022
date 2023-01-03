package day22p1

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
	x, f int
}

type path []move

func parsePath(s string) path {
	p := make(path, 0, len(s))
	f := 0
	x := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			x = 10*x + int(s[i]-'0')
		} else {
			p = append(p, move{x, f})
			x = 0
			switch s[i] {
			case 'L':
				f = (f + 3) % 4
			case 'R':
				f = (f + 1) % 4
			}
		}
	}
	p = append(p, move{x, f})
	return p
}

type coord [2]int

func (inp *input) get(p coord) byte {
	if p[1] >= len(inp.l[p[0]]) {
		return ' '
	}

	return inp.l[p[0]][p[1]]
}

type cell struct {
	p coord
	t byte
	b [4]int
	w [4]int
}

type board struct {
	c    map[coord]*cell
	m, n int
	s    *cell
}

func newBoard(inp *input) *board {
	bo := &board{c: make(map[coord]*cell), m: inp.m, n: inp.n}
	for i := 0; i < inp.m; i++ {
		lb := -1
		lw := -1
		for j := 0; j < inp.n; j++ {
			p := coord{i, j}
			t := inp.get(p)
			if t == ' ' {
				lb = -1
				lw = -1
				continue
			} else if lb == -1 {
				lb = j
			}
			if t == '#' {
				lw = j
			}
			c := &cell{
				p: p,
				t: t,
			}
			c.b[2] = lb
			c.w[2] = lw
			bo.c[p] = c
			if bo.s == nil {
				bo.s = c
			}
		}
		rb := -1
		rw := -1
		for j := inp.n - 1; j >= 0; j-- {
			p := coord{i, j}
			t := inp.get(p)
			if t == ' ' {
				rb = -1
				rw = -1
				continue
			} else if rb == -1 {
				rb = j
			}
			if t == '#' {
				rw = j
			}
			c := bo.c[p]
			c.b[0] = rb
			c.w[0] = rw
		}
	}
	for j := 0; j < inp.n; j++ {
		ub := -1
		uw := -1
		for i := 0; i < inp.m; i++ {
			p := coord{i, j}
			t := inp.get(p)
			if t == ' ' {
				ub = -1
				uw = -1
				continue
			} else if ub == -1 {
				ub = i
			}
			if t == '#' {
				uw = i
			}
			c := bo.c[p]
			c.b[3] = ub
			c.w[3] = uw
		}
		db := -1
		dw := -1
		for i := inp.m - 1; i >= 0; i-- {
			p := coord{i, j}
			t := inp.get(p)
			if t == ' ' {
				db = -1
				dw = -1
				continue
			} else if db == -1 {
				db = i
			}
			if t == '#' {
				dw = i
			}
			c := bo.c[p]
			c.b[1] = db
			c.w[1] = dw
		}
	}
	return bo
}

func (bo *board) maxJump(p coord, f int) int {
	d := [4][2]int{{1, 1}, {0, 1}, {1, -1}, {0, -1}}
	c := bo.c[p]
	g := (f + 2) % 4
	i, s := d[f][0], d[f][1]
	k := s * (c.b[f] - p[i])
	if c.w[f] != -1 {
		return s*(c.w[f]-p[i]) - 1
	}
	if c.w[g] != -1 {
		p[i] = c.b[g]
		return k + 1 + bo.maxJump(p, f)
	}
	return s*(c.b[f]-c.b[g]) + 1
}

func (bo *board) move(p coord, m move) coord {
	d := [4][2]int{{1, 1}, {0, 1}, {1, -1}, {0, -1}}
	c := bo.c[p]
	f, g := m.f, (m.f+2)%4
	i, s := d[f][0], d[f][1]
	n := s*(c.b[f]-c.b[g]) + 1
	if m.x >= n {
		m.x = n + (m.x % n)
	}
	k := bo.maxJump(p, f)
	if k > m.x {
		k = m.x
	}
	l := s * (c.b[f] - p[i])
	if k > l {
		k -= l + 1
		p[i] = c.b[g]
	}
	p[i] += s * k
	return p
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
	p := bo.s.p
	var m move
	for _, m = range pa {
		p = bo.move(p, m)
	}
	fmt.Println(1000*(p[0]+1) + 4*(p[1]+1) + m.f)
}
