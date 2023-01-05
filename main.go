package main

import (
	solution "adventofcode/day24"
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data/day_24_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	solution.Run(s)

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
