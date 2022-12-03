package main

import (
	"bufio"
	"fmt"
	"github.com/advent_of_code_2022/pkg/utils"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partOne() int {
	mostCalories := 0
	currentCalories := 0

	filePath := pkg.GetInputPath("1.txt", false)
	f, err := os.Open(filePath)
	check(err)

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

func partTwo() int {
	var mostCalories [4]int
	currentCalories := 0

	filePath := pkg.GetInputPath("1.txt", false)
	f, err := os.Open(filePath)
	check(err)

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
	mostCalories := partOne()
	println(mostCalories)

	// 206582
	sumCalories := partTwo()
	println(sumCalories)
}
