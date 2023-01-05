package day24

import (
	"bufio"
	"fmt"

	"github.com/gammazero/deque"
)

type coord [2]int

type blizzard struct {
	p coord
	d byte
}

type valley struct {
	b    []*blizzard
	c    [][]int
	m, n int
}

func newValley() *valley {
	return &valley{
		b: make([]*blizzard, 0, 10000),
		c: make([][]int, 0, 100),
	}
}

func (v *valley) append(s string) {
	v.n = len(s)
	c := make([]int, v.n)
	for j := 0; j < v.n; j++ {
		switch s[j] {
		case '#':
			c[j] = -1
		case '.':
			c[j] = 0
		default:
			c[j] = 1
			v.b = append(v.b, &blizzard{p: coord{v.m, j}, d: s[j]})
		}
	}
	v.c = append(v.c, c)
	v.m++
}

func (v *valley) copy() [][]int {
	c := make([][]int, v.m)
	for i := 0; i < v.m; i++ {
		c[i] = make([]int, v.n)
		copy(c[i], v.c[i])
	}
	return c
}

func (v *valley) move(b *blizzard) {
	v.c[b.p[0]][b.p[1]]--
	switch b.d {
	case '^':
		b.p[0]--
	case 'v':
		b.p[0]++
	case '<':
		b.p[1]--
	case '>':
		b.p[1]++
	}
	for i, k := range [2]int{v.m, v.n} {
		switch b.p[i] {
		case 0:
			b.p[i] = k - 2
		case k - 1:
			b.p[i] = 1
		}
	}
	v.c[b.p[0]][b.p[1]]++
}

func (v *valley) evolve() {
	for _, b := range v.b {
		v.move(b)
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func path(vs [][][]int, m, n int, p [3]int, f [2]int) int {
	o := [5][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {0, 0}}
	k := len(vs)
	v := make(map[[3]int]int)
	d := deque.New[[3]int]()
	v[p] = 0
	d.PushBack(p)
	for d.Len() > 0 {
		p = d.PopFront()
		for i := 0; i < len(o); i++ {
			q := [3]int{p[0] + o[i][0], p[1] + o[i][1], p[2] + 1}
			if q[0] < 0 || q[0] >= m || q[1] < 0 || q[1] >= n ||
				vs[q[2]%k][q[0]][q[1]] != 0 {
				continue
			}

			if _, ok := v[q]; ok {
				continue
			}

			v[q] = v[p] + 1
			d.PushBack(q)
			if q[0] == f[0] && q[1] == f[1] {
				return v[q]
			}
		}
	}
	return -1
}

func Run(s *bufio.Scanner) {
	v := newValley()
	for s.Scan() {
		v.append(s.Text())
	}
	k := lcm(v.m-2, v.n-2)
	vs := make([][][]int, k)
	for i := 0; i < k; i++ {
		vs[i] = v.copy()
		v.evolve()
	}

	t := path(vs, v.m, v.n, [3]int{0, 1, 0}, [2]int{v.m - 1, v.n - 2})
	fmt.Println(t)

	t = 0
	t += path(vs, v.m, v.n, [3]int{0, 1, t}, [2]int{v.m - 1, v.n - 2})
	t += path(vs, v.m, v.n, [3]int{v.m - 1, v.n - 2, t}, [2]int{0, 1})
	t += path(vs, v.m, v.n, [3]int{0, 1, t}, [2]int{v.m - 1, v.n - 2})
	fmt.Println(t)
}
