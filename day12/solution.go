package day12

import (
	"bufio"
	"fmt"

	"github.com/gammazero/deque"
)

type heightmap struct {
	h    [][]byte
	m, n int
	s, e [2]int
}
type validFunc func(u, v [2]int) bool
type exitFunc func(u [2]int) bool

func add(u, v [2]int) [2]int {
	return [2]int{u[0] + v[0], u[1] + v[1]}
}

func (hm *heightmap) append(l string) {
	r := []byte(l)
	hm.n = len(r)
	for j := 0; j < hm.n; j++ {
		switch r[j] {
		case 'S':
			r[j] = 'a'
			hm.s = [2]int{hm.m, j}
		case 'E':
			r[j] = 'z'
			hm.e = [2]int{hm.m, j}
		}
	}
	hm.h = append(hm.h, r)
	hm.m++
}

func (hm *heightmap) outside(u [2]int) bool {
	return u[0] < 0 || u[0] >= hm.m || u[1] < 0 || u[1] >= hm.n
}

func (hm *heightmap) height(u [2]int) byte {
	return hm.h[u[0]][u[1]]
}

func (hm *heightmap) valid1(u, v [2]int) bool {
	return hm.height(v) <= hm.height(u)+1
}

func (hm *heightmap) valid2(u, v [2]int) bool {
	return hm.valid1(v, u)
}

func (hm *heightmap) exit1(u [2]int) bool {
	return u == hm.e
}

func (hm *heightmap) exit2(u [2]int) bool {
	return hm.h[u[0]][u[1]] == 'a'
}

func (hm *heightmap) solve(s [2]int, valid validFunc, exit exitFunc) int {
	d := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	m := map[[2]int]int{}
	q := deque.New[[2]int]()
	q.PushBack(s)
	m[s] = 1
	for q.Len() > 0 {
		u := q.PopFront()
		for i := 0; i < len(d); i++ {
			v := add(u, d[i])
			if hm.outside(v) {
				continue
			}

			if valid(u, v) && m[v] == 0 {
				q.PushBack(v)
				m[v] = m[u] + 1
				if exit(v) {
					return m[v] - 1
				}
			}
		}
	}
	return -1
}

func Run(s *bufio.Scanner, part int) {
	hm := &heightmap{}
	for s.Scan() {
		hm.append(s.Text())
	}
	steps := 0
	switch part {
	case 1:
		steps = hm.solve(hm.s, hm.valid1, hm.exit1)
	case 2:
		steps = hm.solve(hm.e, hm.valid2, hm.exit2)
	}
	fmt.Println(steps)
}
