package day10p2

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type cpu struct {
	x int
	n int
	d [240]bool
}

func newCpu() *cpu {
	return &cpu{x: 1, n: 1}
}

func (c *cpu) noop() {
	c.callback(c.x, c.n)
	c.n++
}

func (c *cpu) addx(x int) {
	c.callback(c.x, c.n)
	c.callback(c.x, c.n+1)
	c.x += x
	c.n += 2
}

func (c *cpu) callback(x, n int) {
	if (n-1)%40 >= x-1 && (n-1)%40 <= x+1 {
		c.d[n-1] = true
	}
}

func (c *cpu) print() {
	var b [60]byte
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			if c.d[40*i+j] {
				b[j] = '#'
			} else {
				b[j] = '.'
			}
		}
		fmt.Println(string(b[:]))
	}
}

func Run(s *bufio.Scanner) {
	c := newCpu()
	for s.Scan() {
		l := strings.Split(s.Text(), " ")
		switch l[0] {
		case "noop":
			c.noop()
		case "addx":
			x, _ := strconv.Atoi(l[1])
			c.addx(x)
		}
	}
	c.print()
}
