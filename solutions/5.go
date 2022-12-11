package main

import (
	"bufio"
	"github.com/advent_of_code_2022/pkg/utils"
	"os"
	"regexp"
	"strconv"
)

func partOne5() string {
	filePath := pkg.GetInputPath("5.txt", false)
	f, err := os.Open(filePath)
	pkg.Check(err)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var stacks [10][]string
	var moves [][]int
	for scanner.Scan() {
		line := regexp.MustCompile("[ \\[\\]]").Split(scanner.Text(), -1)
		if line[0] == "move" {
			var move []int
			for _, i := range line {
				j, err := strconv.Atoi(i)
				if err != nil {
					continue
				}
				move = append(move, j)
			}
			moves = append(moves, move)
		} else {
			for width := range line {
				if line[width] != "" && line[width] != "x" {
					stacks[((width + 2) / 3)] = append(stacks[((width+2)/3)], line[width])
				}
			}
		}
	}
	for i, move := range moves {
		println(i, move[0], move[1], move[2])
		n := 0
		for n < move[0] {
			// Move crate from one stack to another
			stacks[move[2]] = append([]string{stacks[move[1]][0]}, stacks[move[2]]...)
			// Delete moved crate from old stack
			stacks[move[1]] = stacks[move[1]][1:]
			n += 1
		}
	}

	finalCrates := ""
	for _, stack := range stacks {
		if stack != nil {
			finalCrates += stack[0]
		}
	}

	return finalCrates
}

func partTwo5() string {
	filePath := pkg.GetInputPath("5.txt", false)
	f, err := os.Open(filePath)
	pkg.Check(err)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var stacks [10][]string
	var moves [][]int
	for scanner.Scan() {
		line := regexp.MustCompile("[ \\[\\]]").Split(scanner.Text(), -1)
		if line[0] == "move" {
			var move []int
			for _, i := range line {
				j, err := strconv.Atoi(i)
				if err != nil {
					continue
				}
				move = append(move, j)
			}
			moves = append(moves, move)
		} else {
			for width := range line {
				if line[width] != "" && line[width] != "x" {
					stacks[((width + 2) / 3)] = append(stacks[((width+2)/3)], line[width])
				}
			}
		}
	}
	for i, move := range moves {
		println(i, move[0], move[1], move[2])
		n := move[0] - 1
		for n >= 0 {
			// Move crate from one stack to another
			stacks[move[2]] = append([]string{stacks[move[1]][n]}, stacks[move[2]]...)
			n -= 1
		}
		// Delete moved crates from old stack
		stacks[move[1]] = stacks[move[1]][move[0]:]
	}

	finalCrates := ""
	for _, stack := range stacks {
		if stack != nil {
			finalCrates += stack[0]
		}
	}

	return finalCrates
}

// https://adventofcode.com/2022/day/5
func main() {
	// ZSQVCCJLL
	answerOne := partOne5()
	println(answerOne)

	// QZFJRWHGS
	answerTwo := partTwo5()
	println(answerTwo)
}
