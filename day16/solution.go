package day16

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/gammazero/deque"
)

type input struct {
	u string
	v []string
	r int
}

func parse(s string) input {
	in := input{}
	ss := strings.Split(s, "; ")
	in.u = ss[0][6:8]
	in.r, _ = strconv.Atoi(ss[0][23:])
	if ss[1][:7] == "tunnels" {
		in.v = strings.Split(ss[1][23:], ", ")
	} else {
		in.v = strings.Split(ss[1][22:], ", ")
	}
	return in
}

type edge struct {
	v int
	t int
}

type cave struct {
	m map[string]int
	t [][]int
	e [][]edge
	r []int
	i []int
	n int
}

func newCave(inputs []input) *cave {
	n := len(inputs)
	c := &cave{
		m: make(map[string]int),
		t: make([][]int, n),
		e: make([][]edge, n),
		r: make([]int, n),
		i: make([]int, n),
	}
	i := 0
	for _, in := range inputs {
		c.m[in.u] = i
		i++
	}
	for _, in := range inputs {
		u := c.m[in.u]
		for _, w := range in.v {
			c.t[u] = append(c.t[u], c.m[w])
			c.r[u] = in.r
		}
		if in.r > 0 {
			c.i[u] = c.n
			c.n++
		}
	}
	c.bfs(c.m["AA"])
	for u := 0; u < n; u++ {
		if c.r[u] > 0 {
			c.bfs(u)
		}
	}
	return c
}

func (c *cave) bfs(u0 int) {
	p := make(map[int]int)
	d := deque.New[int]()
	p[u0] = 0
	d.PushBack(u0)
	for d.Len() > 0 {
		u := d.PopFront()
		for _, v := range c.t[u] {
			if _, ok := p[v]; ok {
				continue
			}

			p[v] = p[u] + 1
			d.PushBack(v)
		}
	}
	for v, t := range p {
		if c.r[v] > 0 {
			c.e[u0] = append(c.e[u0], edge{v, t})
		}
	}
}

type state struct {
	p int
	o int
	r int
	t int
}

func (c *cave) start(o, t int) state {
	return state{
		p: c.m["AA"],
		o: o,
		r: 0,
		t: t,
	}
}

func (s state) next(c *cave) []state {
	ns := make([]state, 0, c.n)
	for _, e := range c.e[s.p] {
		if e.t >= s.t {
			continue
		}

		r := c.r[e.v]
		i := c.i[e.v]
		if s.o&(1<<i) == 0 {
			ns = append(ns, state{
				p: e.v,
				t: s.t - e.t - 1,
				o: s.o | (1 << i),
				r: s.r + r,
			})
		}
	}
	return ns
}

func (s state) search(c *cave, acc int,
	visit map[state][2]int,
	cache map[state]int) int {
	if s.o == (1<<c.n)-1 {
		return s.t * s.r
	}

	if u, ok := cache[s]; ok {
		return u
	}

	s0 := state{p: s.p, o: s.o}
	if v, ok := visit[s0]; ok &&
		v[0] > s.t && v[1]+s.r*(v[0]-s.t) >= acc {
		return 0
	}

	visit[s0] = [2]int{s.t, acc}

	top := s.t * s.r
	for _, w := range s.next(c) {
		if cur := (s.t-w.t)*s.r + w.search(
			c, acc+(s.t-w.t)*s.r, visit, cache); cur > top {
			top = cur
		}
	}

	cache[s] = top
	return top
}

func RunPart1(c *cave) {
	visit := make(map[state][2]int)
	cache := make(map[state]int)
	fmt.Println(c.start(0, 30).search(c, 0, visit, cache))
}

func RunPart2(c *cave) {
	upp := (1 << c.n) - 1
	sol := make(map[int]int)
	out := make(chan [2]int, 1000)
	for o := 1; o < upp; o++ {
		go func(o int) {
			visit := make(map[state][2]int)
			cache := make(map[state]int)
			out <- [2]int{o, c.start(o, 26).search(c, 0, visit, cache)}
		}(o)
	}
	for o := 1; o < upp; o++ {
		z := <-out
		sol[z[0]] = z[1]
	}
	top := 0
	for o1 := 1; o1 <= upp/2; o1++ {
		o2 := upp - o1
		if cur := sol[o1] + sol[o2]; cur > top {
			top = cur
		}
	}
	fmt.Println(top)
}

func Run(s *bufio.Scanner, part int) {
	inputs := make([]input, 0)
	for s.Scan() {
		inputs = append(inputs, parse(s.Text()))
	}
	c := newCave(inputs)
	switch part {
	case 1:
		RunPart1(c)
	case 2:
		RunPart2(c)
	}
}
