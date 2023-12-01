package main

import (
	"fmt"
	"os"

	"github.com/two-first-names/aoc2023/pkg/solutions"
)

func main() {
	day1, err := solutions.Day1()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Day 1:", day1)
}
