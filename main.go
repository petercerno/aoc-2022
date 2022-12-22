package main

import (
	solution "adventofcode/day14"
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data/day_14_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	solution.Run(s, true)

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
