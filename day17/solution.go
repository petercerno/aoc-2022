package day17

import (
	"bufio"
	"fmt"
)

type shape = [][]byte

var shapes = []shape{
	{[]byte("####")},
	{[]byte(".#."), []byte("###"), []byte(".#.")},
	{[]byte("..#"), []byte("..#"), []byte("###")},
	{[]byte("#"), []byte("#"), []byte("#"), []byte("#")},
	{[]byte("##"), []byte("##")},
}

type rock struct {
	x, y int
	w, h int
	t    int
}

type space struct {
	m      map[[2]int]bool
	c, h   int
	dc, dh int64
}

func newSpace() *space {
	return &space{
		m: make(map[[2]int]bool),
	}
}

func (s *space) newRock(t int) *rock {
	w, h := len(shapes[t][0]), len(shapes[t])
	x, y := 2, s.h+h+3
	return &rock{x, y, w, h, t}
}

func (s *space) move(r *rock, dx, dy int) bool {
	if r.x+r.w+dx > 7 || r.x+dx < 0 || r.y-r.h+dy < 0 {
		return false
	}

	for i := 0; i < r.h; i++ {
		for j := 0; j < r.w; j++ {
			if shapes[r.t][i][j] == '#' {
				p := [2]int{r.x + j + dx, r.y - i + dy}
				if _, ok := s.m[p]; ok {
					return false
				}
			}
		}
	}
	r.x += dx
	r.y += dy
	return true
}

func (s *space) freeze(r *rock) {
	for i := 0; i < r.h; i++ {
		for j := 0; j < r.w; j++ {
			if shapes[r.t][i][j] == '#' {
				p := [2]int{r.x + j, r.y - i}
				s.m[p] = true
			}
		}
	}
	if r.y > s.h {
		s.h = r.y
	}
}

const N = 100

func (s *space) suffix() [N]byte {
	var out [N]byte
	for i := s.h; i > s.h-N && i >= 1; i-- {
		k := s.h - i
		for j := 0; j < 7; j++ {
			out[k] *= 2
			if s.m[[2]int{j, i}] {
				out[k]++
			}
		}
	}
	return out
}

type state struct {
	t, i int
	suff [N]byte
}

type value struct {
	c, h int
}

func (s *space) simulate(jet []byte, rocks int64) {
	cache := make(map[state]value)
	n := len(jet)
	t := 0
	i := 0
	r := s.newRock(t)
	for int64(s.c)+s.dc < rocks {
		switch jet[i] {
		case '<':
			s.move(r, -1, 0)
		case '>':
			s.move(r, 1, 0)
		}
		if !s.move(r, 0, -1) {
			s.freeze(r)
			if cache != nil {
				st := state{t: t, i: i, suff: s.suffix()}
				if v, ok := cache[st]; ok {
					dh := int64(s.h - v.h)
					dc := int64(s.c - v.c)
					m := (rocks - int64(s.c)) / dc
					s.dc += m * dc
					s.dh += m * dh
					cache = nil
				} else {
					cache[st] = value{s.c, s.h}
				}
			}
			t = (t + 1) % 5
			r = s.newRock(t)
			s.c++
		}
		i = (i + 1) % n
	}
}

func Run(s *bufio.Scanner, part int) {
	s.Scan()
	spc := newSpace()
	jet := []byte(s.Text())
	switch part {
	case 1:
		spc.simulate(jet, 2022)
	case 2:
		spc.simulate(jet, 1000000000000)
	}
	fmt.Println(int64(spc.h) + spc.dh)
}
