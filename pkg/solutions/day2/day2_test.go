package day2_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/two-first-names/aoc2023/pkg/solutions/day2"
)

func testDecodeGame(t *testing.T, line string, game day2.Game) {
	out, err := day2.DecodeGame(line)
	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(out, game) {
		t.Errorf("DecodeGame(...) got %+v, expected %+v.", out, game)
	}
}

func TestDecodeGame(t *testing.T) {
	game := day2.Game{
		Id: 1,
		Hands: []day2.Hand{
			{
				Red:  4,
				Blue: 3,
			},
			{
				Red:   1,
				Green: 2,
				Blue:  6,
			},
			{
				Green: 2,
			},
		},
	}
	line := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	testDecodeGame(t, line, game)

	game = day2.Game{
		Id: 2,
		Hands: []day2.Hand{
			{
				Blue:  1,
				Green: 2,
			},
			{
				Red:   1,
				Green: 3,
				Blue:  4,
			},
			{
				Green: 1,
				Blue:  1,
			},
		},
	}
	line = "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
	testDecodeGame(t, line, game)

	game = day2.Game{
		Id: 3,
		Hands: []day2.Hand{
			{
				Red:   20,
				Blue:  6,
				Green: 8,
			},
			{
				Red:   4,
				Green: 13,
				Blue:  5,
			},
			{
				Red:   1,
				Green: 5,
			},
		},
	}
	line = "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	testDecodeGame(t, line, game)

	game = day2.Game{
		Id: 4,
		Hands: []day2.Hand{
			{
				Red:   3,
				Blue:  6,
				Green: 1,
			},
			{
				Red:   6,
				Green: 3,
			},
			{
				Red:   14,
				Green: 3,
				Blue:  15,
			},
		},
	}
	line = "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"
	testDecodeGame(t, line, game)

	game = day2.Game{
		Id: 5,
		Hands: []day2.Hand{
			{
				Red:   6,
				Blue:  1,
				Green: 3,
			},
			{
				Red:   1,
				Green: 2,
				Blue:  2,
			},
		},
	}
	line = "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	testDecodeGame(t, line, game)
}

func testIsPossible(t *testing.T, line string, expected bool) {
	game, err := day2.DecodeGame(line)
	if err != nil {
		t.Error(err)
	}

	if actual := game.IsPossible(); actual != expected {
		t.Errorf("IsPossible(Id: %d) should return %t, got %t.", game.Id, expected, actual)
	}
}

func TestIsPossible(t *testing.T) {
	testIsPossible(t, "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true)
	testIsPossible(t, "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true)
	testIsPossible(t, "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false)
	testIsPossible(t, "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false)
	testIsPossible(t, "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true)
}

func TestSumOfPossibleGames(t *testing.T) {
	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	actual, err := day2.SumOfPossibleGames(lines)
	if err != nil {
		t.Error(err)
	}

	if actual != 8 {
		t.Errorf("SumOfPossibleGames(...) returned %d, expected 8.", actual)
	}
}

func testFewestColoursRequired(t *testing.T, line string, expected day2.Hand) {
	game, err := day2.DecodeGame(line)
	if err != nil {
		t.Error(err)
	}

	fewestHand := game.FewestColoursRequired()

	if !cmp.Equal(fewestHand, expected) {
		t.Errorf("Id %d: FewestColoursRequired expected %+v, got %+v.", game.Id, expected, fewestHand)
	}
}

func TestFewestColoursRequired(t *testing.T) {
	testFewestColoursRequired(t, "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", day2.Hand{
		Red:   4,
		Green: 2,
		Blue:  6,
	})
	testFewestColoursRequired(t, "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", day2.Hand{
		Red:   1,
		Green: 3,
		Blue:  4,
	})
	testFewestColoursRequired(t, "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", day2.Hand{
		Red:   20,
		Green: 13,
		Blue:  6,
	})
	testFewestColoursRequired(t, "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", day2.Hand{
		Red:   14,
		Green: 3,
		Blue:  15,
	})
	testFewestColoursRequired(t, "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", day2.Hand{
		Red:   6,
		Green: 3,
		Blue:  2,
	})
}

func testPower(t *testing.T, line string, expected int) {
	game, err := day2.DecodeGame(line)
	if err != nil {
		t.Error(err)
	}

	fewestHand := game.FewestColoursRequired()
	power := fewestHand.Power()

	if power != expected {
		t.Errorf("Id %d: Power expected %d, got %d.", game.Id, expected, power)
	}
}

func TestPower(t *testing.T) {
	testPower(t, "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48)
	testPower(t, "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12)
	testPower(t, "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560)
	testPower(t, "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630)
	testPower(t, "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36)
}

func TestSumOfPowers(t *testing.T) {
	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	actual, err := day2.SumOfPowers(lines)
	if err != nil {
		t.Error(err)
	}

	if actual != 2286 {
		t.Errorf("SumOfPowers(...) returned %d, expected 2286.", actual)
	}
}
