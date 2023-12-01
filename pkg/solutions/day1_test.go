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
