package day11p2

import (
	"fmt"
	"sort"
)

type monkey struct {
	w []int
	o func(x int) int
	t func(x int) int
	c int
}

type monkeys []monkey

func (m monkeys) process(i int) {
	for k := 0; k < len(m[i].w); k++ {
		w := m[i].w[k]
		w = m[i].o(w)
		j := m[i].t(w)
		m[j].w = append(m[j].w, w)
		m[i].c++
	}
	m[i].w = []int{}
}

func (m monkeys) round() {
	for i := 0; i < len(m); i++ {
		m.process(i)
	}
}

func choose(b bool, x, y int) int {
	if b {
		return x
	}

	return y
}

func initMonkeys(test int) monkeys {
	switch test {
	case 0:
		const N = 23 * 19 * 13 * 17
		return monkeys{
			monkey{
				w: []int{79, 98},
				o: func(x int) int { return (x * 19) % N },
				t: func(x int) int { return choose(x%23 == 0, 2, 3) },
			},
			monkey{
				w: []int{54, 65, 75, 74},
				o: func(x int) int { return (x + 6) % N },
				t: func(x int) int { return choose(x%19 == 0, 2, 0) },
			},
			monkey{
				w: []int{79, 60, 97},
				o: func(x int) int { return (x * x) % N },
				t: func(x int) int { return choose(x%13 == 0, 1, 3) },
			},
			monkey{
				w: []int{74},
				o: func(x int) int { return (x + 3) % N },
				t: func(x int) int { return choose(x%17 == 0, 0, 1) },
			},
		}
	default:
		const N = 5 * 2 * 13 * 7 * 19 * 11 * 3 * 17
		return monkeys{
			monkey{
				w: []int{61},
				o: func(x int) int { return (x * 11) % N },
				t: func(x int) int { return choose(x%5 == 0, 7, 4) },
			},
			monkey{
				w: []int{76, 92, 53, 93, 79, 86, 81},
				o: func(x int) int { return (x + 4) % N },
				t: func(x int) int { return choose(x%2 == 0, 2, 6) },
			},
			monkey{
				w: []int{91, 99},
				o: func(x int) int { return (x * 19) % N },
				t: func(x int) int { return choose(x%13 == 0, 5, 0) },
			},
			monkey{
				w: []int{58, 67, 66},
				o: func(x int) int { return (x * x) % N },
				t: func(x int) int { return choose(x%7 == 0, 6, 1) },
			},
			monkey{
				w: []int{94, 54, 62, 73},
				o: func(x int) int { return (x + 1) % N },
				t: func(x int) int { return choose(x%19 == 0, 3, 7) },
			},
			monkey{
				w: []int{59, 95, 51, 58, 58},
				o: func(x int) int { return (x + 3) % N },
				t: func(x int) int { return choose(x%11 == 0, 0, 4) },
			},
			monkey{
				w: []int{87, 69, 92, 56, 91, 93, 88, 73},
				o: func(x int) int { return (x + 8) % N },
				t: func(x int) int { return choose(x%3 == 0, 5, 2) },
			},
			monkey{
				w: []int{71, 57, 86, 67, 96, 95},
				o: func(x int) int { return (x + 7) % N },
				t: func(x int) int { return choose(x%17 == 0, 3, 1) },
			},
		}
	}
}

func Run(test int) {
	m := initMonkeys(test)
	for r := 0; r < 10000; r++ {
		m.round()
	}
	n := len(m)
	c := make([]int, n)
	for i := 0; i < len(m); i++ {
		c[i] = m[i].c
	}
	sort.IntSlice.Sort(c)
	fmt.Println(c[n-1] * c[n-2])
}
