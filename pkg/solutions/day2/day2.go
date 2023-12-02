package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	Red   int
	Green int
	Blue  int
}

func (hand Hand) IsPossible() bool {
	return hand.Red <= 12 && hand.Green <= 13 && hand.Blue <= 14
}

func (hand Hand) Power() int {
	return hand.Red * hand.Green * hand.Blue
}

type Game struct {
	Id    int
	Hands []Hand
}

func (game Game) IsPossible() bool {
	for _, hand := range game.Hands {
		if !hand.IsPossible() {
			return false
		}
	}

	return true
}

func (game Game) FewestColoursRequired() Hand {
	hand := Hand{}

	for _, h := range game.Hands {
		if hand.Red < h.Red {
			hand.Red = h.Red
		}

		if hand.Green < h.Green {
			hand.Green = h.Green
		}

		if hand.Blue < h.Blue {
			hand.Blue = h.Blue
		}
	}

	return hand
}

func DecodeGame(line string) (Game, error) {
	game := Game{}

	colonSplit := strings.Split(line, ": ")

	gameId, err := strconv.Atoi(strings.Split(colonSplit[0], " ")[1])
	if err != nil {
		return Game{}, err
	}

	game.Id = gameId

	hands := []Hand{}

	for _, x := range strings.Split(colonSplit[1], "; ") {
		hand := Hand{}

		for _, y := range strings.Split(x, ", ") {
			parts := strings.Split(y, " ")
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				return Game{}, err
			}
			colour := parts[1]

			if colour == "red" {
				hand.Red = count
			} else if colour == "green" {
				hand.Green = count
			} else if colour == "blue" {
				hand.Blue = count
			}
		}

		hands = append(hands, hand)
	}

	game.Hands = hands

	return game, nil
}

func SumOfPossibleGames(lines []string) (int, error) {
	sum := 0

	for _, line := range lines {
		game, err := DecodeGame(line)
		if err != nil {
			return 0, err
		}

		if game.IsPossible() {
			sum += game.Id
		}
	}

	return sum, nil
}

func SumOfPowers(lines []string) (int, error) {
	sum := 0

	for _, line := range lines {
		game, err := DecodeGame(line)
		if err != nil {
			return 0, err
		}

		fewestHand := game.FewestColoursRequired()

		sum += fewestHand.Power()
	}

	return sum, nil
}

func Solve() (int, error) {
	file, err := os.Open("day2.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return SumOfPossibleGames(lines)
}

func SolvePart2() (int, error) {
	file, err := os.Open("day2.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return SumOfPowers(lines)
}
