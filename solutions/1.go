package main

import (
	"bufio"
	"fmt"
	"github.com/advent_of_code_2022/pkg/utils"
	"os"
	"sort"
	"strconv"
)

func partOne1() int {
	mostCalories := 0
	currentCalories := 0

	filePath := pkg.GetInputPath("1.txt", false)
	f, err := os.Open(filePath)
	pkg.Check(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		calories, err := strconv.Atoi(scanner.Text())
		// New elf
		if err != nil {
			if currentCalories > mostCalories {
				mostCalories = currentCalories
			}
			currentCalories = 0
		}
		currentCalories += calories
	}

	err = f.Close()
	if err != nil {
		return 0
	}
	return mostCalories
}

func partTwo1() int {
	var mostCalories [4]int
	currentCalories := 0

	filePath := pkg.GetInputPath("1.txt", false)
	f, err := os.Open(filePath)
	pkg.Check(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		calories, err := strconv.Atoi(scanner.Text())
		// New elf
		if err != nil {
			mostCalories[0] = currentCalories
			sort.Ints(mostCalories[:])
			currentCalories = 0
		}
		currentCalories += calories
	}

	// Last elf
	mostCalories[0] = currentCalories
	sort.Ints(mostCalories[:])

	err = f.Close()
	if err != nil {
		return 0
	}

	sum := 0
	mostCalories[0] = 0
	for _, v := range mostCalories {
		sum += v
	}
	return sum
}

// https://adventofcode.com/2022/day/1
func main() {
	// 70116
	mostCalories := partOne1()
	println(mostCalories)

	// 206582
	sumCalories := partTwo1()
	println(sumCalories)
}
