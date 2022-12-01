package day01p2

import (
	"bufio"
	"fmt"
	"strconv"
)

type calories struct {
	acc int
	max [3]int
}

func (c *calories) add(cal int) {
	c.acc += cal
}

func (c *calories) close() {
	if c.max[0] < c.acc {
		c.max[0] = c.acc
		if c.max[0] > c.max[1] {
			c.max[0], c.max[1] = c.max[1], c.max[0]
		}
		if c.max[1] > c.max[2] {
			c.max[1], c.max[2] = c.max[2], c.max[1]
		}
	}
	c.acc = 0
}

func (c *calories) total() int {
	return c.max[0] + c.max[1] + c.max[2]
}

func Run(s *bufio.Scanner) {
	c := &calories{}
	for s.Scan() {
		l := s.Text()
		if l == "" {
			c.close()
			continue
		}
		i, _ := strconv.Atoi(l)
		c.add(i)
	}
	c.close()
	fmt.Println(c.total())
}
