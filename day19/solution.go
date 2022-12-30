package day19

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

const dim = 4

type amount [dim]int

func (a amount) add(b amount) (c amount) {
	for i := 0; i < dim; i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

func (a amount) sub(b amount) (c amount) {
	for i := 0; i < dim; i++ {
		c[i] = a[i] - b[i]
	}
	return c
}

func (a amount) max(b amount) (c amount) {
	for i := 0; i < dim; i++ {
		c[i] = a[i]
		if b[i] > a[i] {
			c[i] = b[i]
		}
	}
	return c
}

func (a amount) ge(b amount) bool {
	for i := 0; i < dim; i++ {
		if a[i] < b[i] {
			return false
		}
	}
	return true
}

type blueprint [dim]amount

func parse(l []string) (b blueprint) {
	b[0][0], _ = strconv.Atoi(l[1])
	b[1][0], _ = strconv.Atoi(l[2])
	b[2][0], _ = strconv.Atoi(l[3])
	b[2][1], _ = strconv.Atoi(l[4])
	b[3][0], _ = strconv.Atoi(l[5])
	b[3][2], _ = strconv.Atoi(l[6])
	return
}

type state struct {
	r amount
	a amount
	t int
}

type states []state

func start(t int) state {
	return state{
		r: amount{1, 0, 0, 0},
		t: t,
	}
}

func (u state) betterThan(v state) bool {
	return (u.a.ge(v.a) && u.r.ge(v.r)) ||
		(u.a[3] >= v.a[3]+2 && u.r[3] >= v.r[3]) ||
		(u.a[3] >= v.a[3] && u.r[3] >= v.r[3]+2)
}

func (ss states) evolve(b *blueprint, u amount) states {
	if len(ss) == 0 || ss[0].t == 0 {
		return nil
	}

	m := make(map[state]bool)
	add := func(s state) {
		if s.r[0] > u[0] || s.r[1] > u[1] || s.r[2] > u[2] {
			return
		}

		for z := range m {
			if z.betterThan(s) {
				return
			}
			if s.betterThan(z) {
				delete(m, z)
			}
		}
		m[s] = true
	}
	for _, s0 := range ss {
		for i := 0; i <= dim; i++ {
			s1 := s0
			if i < dim && s0.a.ge(b[i]) {
				s1.a = s0.a.sub(b[i])
				s1.r[i]++
			}
			s1.a = s1.a.add(s0.r)
			s1.t--
			add(s1)
		}
	}
	ns := make([]state, 0, len(m))
	for s := range m {
		ns = append(ns, s)
	}
	return ns
}

func (s state) evolve(b *blueprint) states {
	u := b[0].max(b[1]).max(b[2]).max(b[3])
	ss := states{s}
	for i := 0; i < s.t; i++ {
		ss = ss.evolve(b, u)
	}
	return ss
}

func (ss states) best() int {
	top := 0
	for _, s := range ss {
		if s.a[3] > top {
			top = s.a[3]
		}
	}
	return top
}

func RunPart1(bs []blueprint) {
	ch := make(chan int, len(bs))
	for i := 0; i < len(bs); i++ {
		go func(i int) {
			ch <- (i + 1) * start(24).evolve(&bs[i]).best()
		}(i)
	}
	sum := 0
	for i := 0; i < len(bs); i++ {
		sum += <-ch
	}
	fmt.Println(sum)
}

func RunPart2(bs []blueprint) {
	ch := make(chan int, 3)
	for i := 0; i < len(bs) && i < 3; i++ {
		go func(i int) {
			ch <- start(32).evolve(&bs[i]).best()
		}(i)
	}
	mul := 1
	for i := 0; i < len(bs) && i < 3; i++ {
		mul *= <-ch
	}
	fmt.Println(mul)
}

func Run(s *bufio.Scanner) {
	re := regexp.MustCompile(`-?[0-9]+`)
	bs := make([]blueprint, 0, 50)
	for s.Scan() {
		b := parse(re.FindAllString(s.Text(), -1))
		bs = append(bs, b)
	}
	RunPart1(bs)
	RunPart2(bs)
}
