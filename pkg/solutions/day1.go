package solutions

import (
	"bufio"
	"os"
)

func Day1CalculateLine(line string) int {
	var firstDigit = 0
	var lastDigit = 0
	for _, c := range line {
		if c >= 48 && c <= 57 {
			if firstDigit == 0 {
				firstDigit = int(c) - 48
			}
			lastDigit = int(c) - 48
		}
	}

	return (firstDigit * 10) + lastDigit
}

func Day1CalculateLines(lines []string) int {
	var sum = 0
	for _, line := range lines {
		sum += Day1CalculateLine(line)
	}

	return sum
}

func Day1() (int, error) {
	file, err := os.Open("day1.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return Day1CalculateLines(lines), nil
}
