package day03p1

import (
	"bufio"
	"fmt"
)

func idx(c byte) int {
	if c >= 'a' && c <= 'z' {
		return 0 + int(c-'a')
	} else if c >= 'A' && c <= 'Z' {
		return 26 + int(c-'A')
	}
	panic(fmt.Sprintf("Invalid char: %c", c))
}

func Run(s *bufio.Scanner) {
	sum := 0
	for s.Scan() {
		b := []byte(s.Text())
		n := len(b)
		p := [52]bool{}
		for _, c := range b[:n/2] {
			i := idx(c)
			p[i] = true
		}
		for _, c := range b[n/2:] {
			i := idx(c)
			if p[i] {
				sum += i + 1
				break
			}
		}
	}
	fmt.Println(sum)
}
