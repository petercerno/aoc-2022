package day20

import (
	"bufio"
	"fmt"
	"strconv"
)

type value struct {
	i int
	x int64
}

type file struct {
	l []value
	m map[value]int
}

func newFile() *file {
	return &file{
		l: make([]value, 0, 10000),
		m: make(map[value]int),
	}
}

func (f *file) add(x int64) {
	n := len(f.l)
	v := value{n, x}
	f.l = append(f.l, v)
	f.m[v] = n
}

func (f *file) copy() []value {
	l := make([]value, len(f.l))
	copy(l, f.l)
	return l
}

func (f *file) swap(i, j int) {
	f.m[f.l[i]], f.m[f.l[j]] = j, i
	f.l[i], f.l[j] = f.l[j], f.l[i]
}

func (f *file) move(i int, k int64) {
	n := len(f.l)
	d := 1
	if k < 0 {
		d = n - 1
		k = -k
	}
	k = k % int64(n-1)
	for j := int64(0); j < k; j++ {
		f.swap(i, (i+d)%n)
		i = (i + d) % n
	}
}

func (f *file) mix(l []value) {
	for _, v := range l {
		f.move(f.m[v], v.x)
	}
}

func (f *file) zero() int {
	for _, v := range f.l {
		if v.x == 0 {
			return f.m[v]
		}
	}
	return 0
}

func (f *file) get(i int) int64 {
	return f.l[i%len(f.l)].x
}

func RunPart1(inp []int64) {
	f := newFile()
	for _, x := range inp {
		f.add(x)
	}
	f.mix(f.copy())
	z := f.zero()
	fmt.Println(f.get(z+1000) + f.get(z+2000) + f.get(z+3000))
}

func RunPart2(inp []int64) {
	f := newFile()
	for _, x := range inp {
		f.add(x * 811589153)
	}
	l := f.copy()
	for i := 0; i < 10; i++ {
		f.mix(l)
	}
	z := f.zero()
	fmt.Println(f.get(z+1000) + f.get(z+2000) + f.get(z+3000))
}

func Run(s *bufio.Scanner) {
	inp := make([]int64, 0, 10000)
	for s.Scan() {
		x, _ := strconv.Atoi(s.Text())
		inp = append(inp, int64(x))
	}
	RunPart1(inp)
	RunPart2(inp)
}
