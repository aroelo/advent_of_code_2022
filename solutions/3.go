package main

import (
	"bufio"
	"github.com/advent_of_code_2022/pkg/utils"
	"github.com/juliangruber/go-intersect/v2"
	"os"
	"strings"
)

func partOne3() int {
	filePath := pkg.GetInputPath("3.txt", false)
	f, err := os.Open(filePath)
	pkg.Check(err)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	totalPriority := 0
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "")
		compartmentOne := items[:len(items)/2]
		compartmentTwo := items[len(items)/2:]

		result := intersect.HashGeneric(compartmentOne, compartmentTwo)
		if len(result) != 0 {
			priority := int([]rune(result[0])[0])
			priority -= 96
			if priority < 1 {
				priority += 58
			}
			totalPriority += priority
		}
	}
	err = f.Close()
	if err != nil {
		return 0
	}
	return totalPriority
}

func partTwo3() int {
	filePath := pkg.GetInputPath("3.txt", false)
	f, err := os.Open(filePath)
	pkg.Check(err)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	totalPriority := 0
	var items [][]string
	for scanner.Scan() {
		items = append(items, strings.Split(scanner.Text(), ""))

		if len(items) == 3 {
			firstIntersect := intersect.HashGeneric(items[0], items[1])
			result := intersect.HashGeneric(firstIntersect, items[2])

			if len(result) != 0 {
				priority := int([]rune(result[0])[0])
				priority -= 96
				if priority < 1 {
					priority += 58
				}
				totalPriority += priority
			}
			items = make([][]string, 0)
		}
	}
	err = f.Close()
	if err != nil {
		return 0
	}
	return totalPriority
}

// https://adventofcode.com/2022/day/3
func main() {
	// 7568
	answerOne := partOne3()
	println(answerOne)

	// 2780
	answerTwo := partTwo3()
	println(answerTwo)
}
