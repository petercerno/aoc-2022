package main

import (
	solution "adventofcode/day17"
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data/day_17_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	solution.Run(s, 2)

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
