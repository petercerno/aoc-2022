package day21

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type expr struct {
	a, b *expr
	n    string
	o    string
	v    float64
	f    bool
}

func (e *expr) eval() float64 {
	if e.f || e.a == nil || e.b == nil {
		return e.v
	}

	e.a.eval()
	e.b.eval()
	switch e.o {
	case "+":
		e.v = e.a.v + e.b.v
	case "-":
		e.v = e.a.v - e.b.v
	case "*":
		e.v = e.a.v * e.b.v
	case "/":
		e.v = e.a.v / e.b.v
	}
	e.f = true
	return e.v
}

func (e *expr) reset() *expr {
	if e.o == "" {
		return e
	}
	if e.a != nil {
		e.a.reset()
	}
	if e.b != nil {
		e.b.reset()
	}
	e.f = false
	return e
}

type exprTree map[string]*expr

func newExprTree() exprTree {
	return make(map[string]*expr)
}

func (t exprTree) parse(s string) {
	get := func(w string) *expr {
		if _, ok := t[w]; !ok {
			t[w] = &expr{}
		}
		return t[w]
	}
	l := strings.Split(s, ": ")
	r := strings.Split(l[1], " ")
	if len(r) == 1 {
		v, _ := strconv.Atoi(r[0])
		lhs := get(l[0])
		*lhs = expr{
			n: l[0],
			v: float64(v),
			f: true,
		}
	} else {
		lhs := get(l[0])
		*lhs = expr{
			n: l[0],
			a: get(r[0]),
			b: get(r[2]),
			o: r[1],
			f: false,
		}
	}
}

func Run(s *bufio.Scanner) {
	t := newExprTree()
	for s.Scan() {
		t.parse(s.Text())
	}
	fmt.Println(int64(t["root"].eval()))

	eval := func(v int64) float64 {
		t["humn"].v = float64(v)
		return t["root"].a.reset().eval() - t["root"].b.reset().eval()
	}
	lb := int64(0)
	ub := int64(10000000000000)
	for lb < ub {
		v := (lb + ub) / 2
		y := eval(v)
		if y > 0 {
			lb = v + 1
		} else if y < 0 {
			ub = v - 1
		} else {
			lb, ub = v, v
		}
	}
	if lb == ub {
		fmt.Println(lb)
	}
}
