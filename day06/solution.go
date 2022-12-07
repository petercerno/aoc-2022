package day06p1

import (
	"bufio"
	"fmt"
)

func Run(s *bufio.Scanner, p int) {
	s.Scan()
	inp := []byte(s.Text())
	cnt := make([]int, 30)
	num := 0
	for i := 0; i < len(inp); i++ {
		u := inp[i] - 'a'
		cnt[u]++
		switch cnt[u] {
		case 1:
			num++
		case 2:
			num--
		}
		if i >= p {
			v := inp[i-p] - 'a'
			cnt[v]--
			switch cnt[v] {
			case 1:
				num++
			case 0:
				num--
			}
		}
		if num == p {
			fmt.Println(i + 1)
			break
		}
	}
}
