package day02p2

import (
	"bufio"
	"fmt"
)

func Run(s *bufio.Scanner) {
	sum := 0
	for s.Scan() {
		b := []byte(s.Text())
		i := int(b[0] - 'A')
		j := int(b[2] - 'X')
		k := (i + j + 2) % 3
		sum += 3*j + k + 1
	}
	fmt.Println(sum)
}
