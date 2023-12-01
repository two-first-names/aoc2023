package solutions_test

import (
	"testing"

	"github.com/two-first-names/aoc2023/pkg/solutions"
)

func TestDay1CalculateLine(t *testing.T) {
	if got := solutions.Day1CalculateLine("1abc2"); got != 12 {
		t.Errorf("Day1CalculateLine(\"1abc2\") expected 12, got %d", got)
	}

	if got := solutions.Day1CalculateLine("pqr3stu8vwx"); got != 38 {
		t.Errorf("Day1CalculateLine(\"pqr3stu8vwx\") expected 38, got %d", got)
	}

	if got := solutions.Day1CalculateLine("a1b2c3d4e5f"); got != 15 {
		t.Errorf("Day1CalculateLine(\"a1b2c3d4e5f\") expected 15, got %d", got)
	}

	if got := solutions.Day1CalculateLine("treb7uchet"); got != 77 {
		t.Errorf("Day1CalculateLine(\"treb7uchet\") expected 77, got %d", got)
	}
}

func TestDay1CalculateLines(t *testing.T) {
	lines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	if got := solutions.Day1CalculateLines(lines); got != 142 {
		t.Errorf("Day1CalculateLines(...) expected 142, got %d", got)
	}
}

func TestDay1Part2CalculateLine(t *testing.T) {
	if got := solutions.Day1Part2CalculateLine("two1nine"); got != 29 {
		t.Errorf("solutions.Day1CalculateLine(solutions.Day1ReplaceNumbers(\"two1nine\")) expected 29, got %d", got)
	}

	if got := solutions.Day1Part2CalculateLine("eightwothree"); got != 83 {
		t.Errorf("solutions.Day1CalculateLine(solutions.Day1ReplaceNumbers(\"eightwothree\")) expected 83, got %d", got)
	}

	if got := solutions.Day1Part2CalculateLine("abcone2threexyz"); got != 13 {
		t.Errorf("solutions.Day1CalculateLine(solutions.Day1ReplaceNumbers(\"abcone2threexyz\")) expected 13, got %d", got)
	}

	if got := solutions.Day1Part2CalculateLine("xtwone3four"); got != 24 {
		t.Errorf("solutions.Day1CalculateLine(solutions.Day1ReplaceNumbers(\"xtwone3four\")) expected 24, got %d", got)
	}

	if got := solutions.Day1Part2CalculateLine("4nineeightseven2"); got != 42 {
		t.Errorf("solutions.Day1CalculateLine(solutions.Day1ReplaceNumbers(\"4nineeightseven2\")) expected 42, got %d", got)
	}

	if got := solutions.Day1Part2CalculateLine("zoneight234"); got != 14 {
		t.Errorf("solutions.Day1CalculateLine(solutions.Day1ReplaceNumbers(\"zoneight234\")) expected 14, got %d", got)
	}

	if got := solutions.Day1Part2CalculateLine("7pqrstsixteen"); got != 76 {
		t.Errorf("solutions.Day1CalculateLine(solutions.Day1ReplaceNumbers(\"7pqrstsixteen\")) expected 76, got %d", got)
	}
}

func TestDay1Part2CalculateLines(t *testing.T) {
	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	if got := solutions.Day1Part2CalculateLines(lines); got != 281 {
		t.Errorf("Day1CalculateLines(...) expected 281, got %d", got)
	}
}
