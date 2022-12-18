package day09p2

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n >= 0 {
		return n
	}

	return -n
}

func sign(n int) int {
	switch {
	case n == 0:
		return 0
	case n > 0:
		return 1
	default:
		return -1
	}
}

const N = 10

type rope struct {
	p [N][2]int
	v map[[2]int]bool
	c int
}

func newRope() *rope {
	return &rope{
		v: map[[2]int]bool{{0, 0}: true},
		c: 1,
	}
}

func (r *rope) move(d [2]int) {
	r.p[0][0] += d[0]
	r.p[0][1] += d[1]
	for k := 1; k < N; k++ {
		if abs(r.p[k-1][0]-r.p[k][0]) > 1 || abs(r.p[k-1][1]-r.p[k][1]) > 1 {
			for i := 0; i < 2; i++ {
				r.p[k][i] += sign(r.p[k-1][i] - r.p[k][i])
			}
		}
	}
	v := r.v[r.p[N-1]]
	if !v {
		r.v[r.p[N-1]] = true
		r.c++
	}
}

func Run(s *bufio.Scanner) {
	m := map[string][2]int{
		"L": {-1, 0},
		"R": {+1, 0},
		"U": {0, -1},
		"D": {0, +1},
	}
	r := newRope()
	for s.Scan() {
		l := strings.Split(s.Text(), " ")
		d := m[l[0]]
		n, _ := strconv.Atoi(l[1])
		for i := 0; i < n; i++ {
			r.move(d)
		}
	}
	fmt.Println(r.c)
}
