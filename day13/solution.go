package day13

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

type packet struct {
	p []packet
	v int
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func split(s string) []string {
	n := len(s)
	if n == 0 {
		return []string{}
	}

	t := make([]string, 0, (n+1)/2)
	b := make([]byte, 0, n)
	d := 0
	for i := 0; i < n; i++ {
		switch s[i] {
		case '[':
			d++
		case ']':
			d--
		case ',':
			if d == 0 {
				t = append(t, string(b))
				b = b[:0]
				continue
			}
		}
		b = append(b, s[i])
	}
	t = append(t, string(b))
	b = nil
	return t
}

func parse(s string) packet {
	p := packet{}
	n := len(s)
	if s[0] == '[' {
		t := split(s[1 : n-1])
		m := len(t)
		p.p = make([]packet, m)
		for i := 0; i < m; i++ {
			p.p[i] = parse(t[i])
		}
	} else {
		p.v, _ = strconv.Atoi(s)
	}
	return p
}

func (p packet) nest() packet {
	if p.p != nil {
		return p
	}

	return packet{
		p: []packet{
			p,
		},
	}
}

func cmpV(v1, v2 int) int {
	switch {
	case v1 == v2:
		return 0
	case v1 < v2:
		return -1
	default:
		return 1
	}
}

func cmpP(p1, p2 packet) int {
	if p1.p == nil && p2.p == nil {
		return cmpV(p1.v, p2.v)
	}

	p1 = p1.nest()
	p2 = p2.nest()
	n1 := len(p1.p)
	n2 := len(p2.p)
	n0 := min(n1, n2)
	for i := 0; i < n0; i++ {
		c := cmpP(p1.p[i], p2.p[i])
		if c != 0 {
			return c
		}
	}
	return cmpV(n1, n2)
}

type indexPacket struct {
	i int
	p packet
}
type byIndexPackets []indexPacket

func (a byIndexPackets) Len() int           { return len(a) }
func (a byIndexPackets) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIndexPackets) Less(i, j int) bool { return cmpP(a[i].p, a[j].p) == -1 }

func RunPart1(s *bufio.Scanner) {
	sum := 0
	row := 0
	for s.Scan() {
		l1 := s.Text()
		s.Scan()
		l2 := s.Text()
		s.Scan()
		p1, p2 := parse(l1), parse(l2)
		row++
		if cmpP(p1, p2) == -1 {
			sum += row
		}
	}
	fmt.Println(sum)
}

func RunPart2(s *bufio.Scanner) {
	p := []indexPacket{
		{0, parse("[[2]]")},
		{1, parse("[[6]]")},
	}
	n := 2
	for s.Scan() {
		l := s.Text()
		if l == "" {
			continue
		}

		p = append(p, indexPacket{n, parse(l)})
		n++
	}
	sort.Sort(byIndexPackets(p))
	pos := make([]int, 0, 2)
	for i := 0; i < n; i++ {
		if p[i].i < 2 {
			pos = append(pos, i+1)
		}
	}
	fmt.Println(pos[0] * pos[1])
}

func Run(s *bufio.Scanner, part int) {
	switch part {
	case 1:
		RunPart1(s)
	case 2:
		RunPart2(s)
	}
}
