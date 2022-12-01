package day01p1

import (
	"bufio"
	"fmt"
	"strconv"
)

type calories struct {
	acc int
	max int
}

func (c *calories) add(cal int) {
	c.acc += cal
}

func (c *calories) close() {
	if c.max < c.acc {
		c.max = c.acc
	}
	c.acc = 0
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
	fmt.Println(c.max)
}
