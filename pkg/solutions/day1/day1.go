package day1

import (
	"bufio"
	"os"
	"strings"
)

func CalculateLine(line string) int {
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

func CalculateLines(lines []string) int {
	var sum = 0
	for _, line := range lines {
		sum += CalculateLine(line)
	}

	return sum
}

func ReplaceNumbers(line string) string {
	var s = line

	s = strings.Replace(s, "one", "o1e", -1)
	s = strings.Replace(s, "two", "t2w", -1)
	s = strings.Replace(s, "three", "t3e", -1)
	s = strings.Replace(s, "four", "f4r", -1)
	s = strings.Replace(s, "five", "f5e", -1)
	s = strings.Replace(s, "six", "s6x", -1)
	s = strings.Replace(s, "seven", "s7n", -1)
	s = strings.Replace(s, "eight", "e8t", -1)
	s = strings.Replace(s, "nine", "n9e", -1)

	return s
}

func Part2CalculateLine(line string) int {
	return CalculateLine(ReplaceNumbers(line))
}

func Part2CalculateLines(lines []string) int {
	var sum = 0
	for _, line := range lines {
		sum += Part2CalculateLine(line)
	}

	return sum
}

func Solve() (int, error) {
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

	return CalculateLines(lines), nil
}

func SolvePart2() (int, error) {
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

	return Part2CalculateLines(lines), nil
}
