package day10p1

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type cpu struct {
	x   int
	n   int
	sum int
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
	if (n-20)%40 == 0 {
		c.sum += n * x
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
	fmt.Println(c.sum)
}
