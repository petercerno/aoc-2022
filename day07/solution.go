package day07

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type dir struct {
	p *dir
	d map[string]*dir
	f map[string]int
	s int
}

func newDir(p *dir) *dir {
	return &dir{
		p: p,
		d: make(map[string]*dir),
		f: make(map[string]int),
	}
}

func (d *dir) getSubDir(n string) *dir {
	c := d.d[n]
	if c != nil {
		return c
	}

	c = newDir(d)
	d.d[n] = c
	return c
}

func (d *dir) getSize() int {
	d.s = 0
	for _, c := range d.d {
		d.s += c.getSize()
	}
	for _, s := range d.f {
		d.s += s
	}
	return d.s
}

func (d *dir) getAllDirs() []*dir {
	dirs := []*dir{d}
	for _, c := range d.d {
		dirs = append(dirs, c.getAllDirs()...)
	}
	return dirs
}

type input struct {
	cmd []string
	out [][]string
}

func Run(s *bufio.Scanner, part1 bool) {
	inp := make([]input, 0, 100)
	for s.Scan() {
		l := strings.Split(s.Text(), " ")
		if l[0] == "$" {
			inp = append(inp, input{cmd: l[1:]})
		} else {
			i := len(inp) - 1
			inp[i].out = append(inp[i].out, l)
		}
	}
	r := newDir(nil)
	d := r
	for _, x := range inp {
		switch x.cmd[0] {
		case "cd":
			switch x.cmd[1] {
			case "/":
				d = r
			case "..":
				d = d.p
			default:
				d = d.getSubDir(x.cmd[1])
			}
		case "ls":
			for _, o := range x.out {
				if o[0] == "dir" {
					d.getSubDir(o[1])
				} else {
					s, _ := strconv.Atoi(o[0])
					d.f[o[1]] = s
				}
			}
		}
	}
	r.getSize()
	dirs := r.getAllDirs()
	if part1 {
		sum := 0
		for _, d := range dirs {
			if d.s <= 100000 {
				sum += d.s
			}
		}
		fmt.Println(sum)
	} else {
		best := r.s
		unused := 70000000 - r.s
		for _, d := range dirs {
			if unused+d.s >= 30000000 && d.s < best {
				best = d.s
			}
		}
		fmt.Println(best)
	}
}
