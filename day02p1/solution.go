package day02p1

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
		k := (j + 4 - i) % 3
		sum += 3*k + j + 1
	}
	fmt.Println(sum)
}
