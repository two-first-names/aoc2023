package day1_test

import (
	"testing"

	"github.com/two-first-names/aoc2023/pkg/solutions/day1"
)

func TestCalculateLine(t *testing.T) {
	if got := day1.CalculateLine("1abc2"); got != 12 {
		t.Errorf("CalculateLine(\"1abc2\") expected 12, got %d", got)
	}

	if got := day1.CalculateLine("pqr3stu8vwx"); got != 38 {
		t.Errorf("CalculateLine(\"pqr3stu8vwx\") expected 38, got %d", got)
	}

	if got := day1.CalculateLine("a1b2c3d4e5f"); got != 15 {
		t.Errorf("CalculateLine(\"a1b2c3d4e5f\") expected 15, got %d", got)
	}

	if got := day1.CalculateLine("treb7uchet"); got != 77 {
		t.Errorf("CalculateLine(\"treb7uchet\") expected 77, got %d", got)
	}
}

func TestCalculateLines(t *testing.T) {
	lines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	if got := day1.CalculateLines(lines); got != 142 {
		t.Errorf("CalculateLines(...) expected 142, got %d", got)
	}
}

func TestPart2CalculateLine(t *testing.T) {
	if got := day1.Part2CalculateLine("two1nine"); got != 29 {
		t.Errorf("CalculateLine(day1.ReplaceNumbers(\"two1nine\")) expected 29, got %d", got)
	}

	if got := day1.Part2CalculateLine("eightwothree"); got != 83 {
		t.Errorf("CalculateLine(day1.ReplaceNumbers(\"eightwothree\")) expected 83, got %d", got)
	}

	if got := day1.Part2CalculateLine("abcone2threexyz"); got != 13 {
		t.Errorf("CalculateLine(day1.ReplaceNumbers(\"abcone2threexyz\")) expected 13, got %d", got)
	}

	if got := day1.Part2CalculateLine("xtwone3four"); got != 24 {
		t.Errorf("CalculateLine(day1.ReplaceNumbers(\"xtwone3four\")) expected 24, got %d", got)
	}

	if got := day1.Part2CalculateLine("4nineeightseven2"); got != 42 {
		t.Errorf("CalculateLine(day1.ReplaceNumbers(\"4nineeightseven2\")) expected 42, got %d", got)
	}

	if got := day1.Part2CalculateLine("zoneight234"); got != 14 {
		t.Errorf("CalculateLine(day1.ReplaceNumbers(\"zoneight234\")) expected 14, got %d", got)
	}

	if got := day1.Part2CalculateLine("7pqrstsixteen"); got != 76 {
		t.Errorf("CalculateLine(day1.ReplaceNumbers(\"7pqrstsixteen\")) expected 76, got %d", got)
	}
}

func TestPart2CalculateLines(t *testing.T) {
	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	if got := day1.Part2CalculateLines(lines); got != 281 {
		t.Errorf("CalculateLines(...) expected 281, got %d", got)
	}
}
