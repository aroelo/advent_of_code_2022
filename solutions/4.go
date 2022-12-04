package main

import (
	"bufio"
	"github.com/advent_of_code_2022/pkg/utils"
	"os"
	"regexp"
	"strconv"
)

func partOne4() int {
	filePath := pkg.GetInputPath("4.txt", false)
	f, err := os.Open(filePath)
	pkg.Check(err)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	overlappingPairs := 0
	for scanner.Scan() {
		sections := regexp.MustCompile("[-,]").Split(scanner.Text(), 4)
		// Convert to list of integers
		s := make([]int, len(sections))
		for i, section := range sections {
			s[i], err = strconv.Atoi(section)
			if err != nil {
				return 0
			}
		}

		startOne, endOne, startTwo, endTwo := s[0], s[1], s[2], s[3]
		if startOne <= startTwo {
			if endOne >= endTwo {
				overlappingPairs += 1
				continue
			}
		}
		if startTwo <= startOne {
			if endTwo >= endOne {
				overlappingPairs += 1
			}
		}
	}
	err = f.Close()
	if err != nil {
		return 0
	}
	return overlappingPairs
}

func partTwo4() int {
	filePath := pkg.GetInputPath("4.txt", false)
	f, err := os.Open(filePath)
	pkg.Check(err)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	overlappingPairs := 0
	for scanner.Scan() {
		sections := regexp.MustCompile("[-,]").Split(scanner.Text(), 4)
		// Convert to list of integers
		s := make([]int, len(sections))
		for i, section := range sections {
			s[i], err = strconv.Atoi(section)
			if err != nil {
				return 0
			}
		}

		startOne, endOne, startTwo, endTwo := s[0], s[1], s[2], s[3]
		if startTwo > endOne {
			if endTwo > startOne {
				continue
			}
		}
		if startOne > endTwo {
			if endOne > startTwo {
				continue
			}
		}
		overlappingPairs += 1
	}
	err = f.Close()
	if err != nil {
		return 0
	}
	return overlappingPairs
}

// https://adventofcode.com/2022/day/4
func main() {
	// 569
	answerOne := partOne4()
	println(answerOne)

	// 936
	answerTwo := partTwo4()
	println(answerTwo)
}
