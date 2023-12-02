package main

import (
	"fmt"
	"os"

	"github.com/two-first-names/aoc2023/pkg/solutions/day1"
	"github.com/two-first-names/aoc2023/pkg/solutions/day2"
)

func main() {
	d1, err := day1.Solve()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Day 1:", d1)

	d1p1, err := day1.SolvePart2()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Day 1, Part 2:", d1p1)

	d2, err := day2.Solve()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Day 2:", d2)

	d2p2, err := day2.SolvePart2()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Day 2:", d2p2)
}
