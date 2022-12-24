package main

import (
	solution "adventofcode/day15"
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data/day_15_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	solution.Run(s, 2, 4000000)

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
