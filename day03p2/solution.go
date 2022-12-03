package day03p2

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
	done := false
	for !done {
		p := [52]int{}
		for k := 0; k < 3; k++ {
			if !s.Scan() {
				done = true
				break
			}
			b := []byte(s.Text())
			for _, c := range b {
				i := idx(c)
				if p[i] == k {
					p[i] = k + 1
				}
				if p[i] == 3 {
					sum += i + 1
					break
				}
			}
		}
	}
	fmt.Println(sum)
}
