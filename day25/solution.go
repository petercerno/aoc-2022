package day25

import (
	"bufio"
	"fmt"
)

func num(c byte) int64 {
	switch c {
	case '=':
		return -2
	case '-':
		return -1
	default:
		return int64(c - '0')
	}
}

func snafu(d int64) byte {
	switch d {
	case -2:
		return '='
	case -1:
		return '-'
	default:
		return byte('0' + d)
	}
}

func snafu2num(s []byte) int64 {
	x := int64(0)
	for _, c := range s {
		x = 5*x + num(c)
	}
	return x
}

func num2snafu(x int64) string {
	b := make([]byte, 0, 10)
	for x > 0 {
		d := x % 5
		c := int64(0)
		if d > 2 {
			c = 5
		}
		b = append(b, snafu(d-c))
		x += c
		x /= 5
	}
	n := len(b)
	for i := 0; i < n/2; i++ {
		b[i], b[n-1-i] = b[n-1-i], b[i]
	}
	return string(b)
}

func Run(s *bufio.Scanner) {
	sum := int64(0)
	for s.Scan() {
		sum += snafu2num([]byte(s.Text()))
	}
	fmt.Println(sum, num2snafu(sum))
}
