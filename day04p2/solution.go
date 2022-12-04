package day04p2

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parse(s string) (int, int) {
	i := strings.Split(s, "-")
	a, _ := strconv.Atoi(i[0])
	b, _ := strconv.Atoi(i[1])
	return a, b
}

func Run(s *bufio.Scanner) {
	cnt := 0
	for s.Scan() {
		l := s.Text()
		i := strings.Split(l, ",")
		a, b := parse(i[0])
		c, d := parse(i[1])
		if b >= c && d >= a {
			cnt++
		}
	}
	fmt.Println(cnt)
}
