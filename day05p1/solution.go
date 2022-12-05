package day05p1

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type crates struct {
	v [][]byte
}

type move struct {
	n, f, t int
}

func (c *crates) apply(m move) {
	nf := len(c.v[m.f])
	if m.n > nf {
		panic(fmt.Sprintf("Invalid move: %v for crates: %v", m, c))
	}
	for i := 0; i < m.n; i++ {
		c.v[m.t] = append(c.v[m.t], c.v[m.f][nf-i-1])
	}
	c.v[m.f] = c.v[m.f][:nf-m.n]
}

func (c *crates) topStr() string {
	m := len(c.v)
	b := make([]byte, m)
	for i := 0; i < m; i++ {
		n := len(c.v[i])
		if n > 0 {
			b = append(b, c.v[i][n-1])
		}
	}
	return string(b)
}

func newCrates(top [][]byte) *crates {
	c := &crates{}
	n := len(top) - 1
	m := (len(top[0]) + 1) / 4
	c.v = make([][]byte, m)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			if top[i][1+4*j] != ' ' {
				c.v[j] = append(c.v[j], top[i][1+4*j])
			}
		}
	}
	return c
}

func parseMove(l string) move {
	s := strings.Split(l, " ")
	n, _ := strconv.Atoi(s[1])
	f, _ := strconv.Atoi(s[3])
	t, _ := strconv.Atoi(s[5])
	return move{n, f - 1, t - 1}
}

func Run(s *bufio.Scanner) {
	var c *crates
	top := [][]byte{}
	for s.Scan() {
		l := s.Text()
		if c == nil {
			if l == "" {
				c = newCrates(top)
			} else {
				top = append(top, []byte(l))
			}
		} else {
			c.apply(parseMove(l))
		}
	}
	fmt.Println(c.topStr())
}
