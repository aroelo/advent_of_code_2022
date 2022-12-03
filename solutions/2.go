package main

import (
	"bufio"
	"github.com/advent_of_code_2022/pkg/utils"
	"os"
	"strings"
)

// Rock: A/X - 1
// Paper: B/Y - 2
// Scissors: C/Z - 3
func partOne2() int {
	shapeScoringMatrix := map[string]int{"X": 1, "Y": 2, "Z": 3}
	winScoringMatrix := map[string]map[string]int{
		"X": {"A": 3, "B": 0, "C": 6},
		"Y": {"A": 6, "B": 3, "C": 0},
		"Z": {"A": 0, "B": 6, "C": 3},
	}

	filePath := pkg.GetInputPath("2.txt", false)
	f, err := os.Open(filePath)
	pkg.Check(err)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	totalScore := 0
	for scanner.Scan() {
		shapes := strings.Split(scanner.Text(), " ")
		// Get shapes
		opponentShape := shapes[0]
		myShape := shapes[1]

		// Calculate score
		score := shapeScoringMatrix[myShape]              // Shape
		score += winScoringMatrix[myShape][opponentShape] // Winner/Loser

		totalScore += score
	}

	err = f.Close()
	if err != nil {
		return 0
	}
	return totalScore
}

// X: 1, Y: 2, Z: 3
// Lose: X (A-Z, B-X, C-Y)
// Draw: Y (A-X, B-Y, C-Z)
// Win: Z  (A-Y, B-Z, C-X)
func partTwo2() int {
	shapeScoringMatrix := map[string]map[string]int{
		"X": {"A": 3, "B": 1, "C": 2},
		"Y": {"A": 1, "B": 2, "C": 3},
		"Z": {"A": 2, "B": 3, "C": 1},
	}
	winScoringMatrix := map[string]int{"X": 0, "Y": 3, "Z": 6}

	filePath := pkg.GetInputPath("2.txt", false)
	f, err := os.Open(filePath)
	pkg.Check(err)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	totalScore := 0
	for scanner.Scan() {
		shapes := strings.Split(scanner.Text(), " ")
		// Get shapes
		opponentShape := shapes[0]
		myShape := shapes[1]

		// Calculate score
		score := shapeScoringMatrix[myShape][opponentShape] // Shape
		score += winScoringMatrix[myShape]                  // Winner/Loser

		totalScore += score
	}

	err = f.Close()
	if err != nil {
		return 0
	}
	return totalScore
}

// https://adventofcode.com/2022/day/2
func main() {
	// 14069
	answerOne := partOne2()
	println(answerOne)

	// 12411
	answerTwo := partTwo2()
	println(answerTwo)
}
