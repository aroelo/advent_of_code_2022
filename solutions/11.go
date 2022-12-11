package main

func partOne11() int {
	// Test input
	//monkeys := 3
	//items := [][]int{
	//	{79, 98},
	//	{54, 65, 75, 74},
	//	{79, 60, 97},
	//	{74},
	//}
	//operations := []func(int) int{
	//	func(worry int) int { return worry * 19 },
	//	func(worry int) int { return worry + 6 },
	//	func(worry int) int { return worry * worry },
	//	func(worry int) int { return worry + 3 },
	//}
	//testCondition := []int{23, 19, 13, 17}
	//testTrue := []int{2, 2, 1, 0}
	//testFalse := []int{3, 0, 3, 1}

	// Real input
	monkeys := 7
	items := [][]int{
		{89, 73, 66, 57, 64, 80},
		{83, 78, 81, 55, 81, 59, 69},
		{76, 91, 58, 85},
		{71, 72, 74, 76, 68},
		{98, 85, 84},
		{78},
		{86, 70, 60, 88, 88, 78, 74, 83},
		{81, 58},
	}
	operations := []func(int) int{
		func(worry int) int { return worry * 3 },
		func(worry int) int { return worry + 1 },
		func(worry int) int { return worry * 13 },
		func(worry int) int { return worry * worry },
		func(worry int) int { return worry + 7 },
		func(worry int) int { return worry + 8 },
		func(worry int) int { return worry + 4 },
		func(worry int) int { return worry + 5 },
	}
	testCondition := []int{13, 3, 7, 2, 19, 5, 11, 17}
	testTrue := []int{6, 7, 1, 6, 5, 3, 1, 3}
	testFalse := []int{2, 4, 4, 0, 7, 0, 2, 5}

	inspections := make([]int, monkeys+1)
	round := 0
	for round < 20 {
		monkey := 0
		for monkey <= monkeys {
			// Loop through items that monkey has
			for _, oldWorry := range items[monkey] {
				// count inspection
				inspections[monkey] += 1

				// inspect item
				worry := operations[monkey](oldWorry)
				// divide by 3
				worry = worry / 3
				// decide which monkey receives item
				var receivingMonkey int
				if worry%testCondition[monkey] == 0 {
					receivingMonkey = testTrue[monkey]
				} else {
					receivingMonkey = testFalse[monkey]
				}
				// give item to new monkey
				items[receivingMonkey] = append(items[receivingMonkey], worry)
				// remove item from current monkey
				items[monkey] = items[monkey][1:]
			}
			monkey += 1
		}
		round += 1
	}
	maxInspections := 0
	secondMaxInspections := 0
	for _, i := range inspections {
		if i > maxInspections {
			secondMaxInspections = maxInspections
			maxInspections = i
		} else if i > secondMaxInspections {
			secondMaxInspections = i
		}
	}
	return maxInspections * secondMaxInspections
}

func partTwo11() int {
	// Test input
	//monkeys := 3
	//items := [][]int{
	//	{79, 98},
	//	{54, 65, 75, 74},
	//	{79, 60, 97},
	//	{74},
	//}
	//operations := []func(int) int{
	//	func(worry int) int { return worry * 19 },
	//	func(worry int) int { return worry + 6 },
	//	func(worry int) int { return worry * worry },
	//	func(worry int) int { return worry + 3 },
	//}
	//testCondition := []int{23, 19, 13, 17}
	//testTrue := []int{2, 2, 1, 0}
	//testFalse := []int{3, 0, 3, 1}

	// Real input
	monkeys := 7
	items := [][]int{
		{89, 73, 66, 57, 64, 80},
		{83, 78, 81, 55, 81, 59, 69},
		{76, 91, 58, 85},
		{71, 72, 74, 76, 68},
		{98, 85, 84},
		{78},
		{86, 70, 60, 88, 88, 78, 74, 83},
		{81, 58},
	}
	operations := []func(int) int{
		func(worry int) int { return worry * 3 },
		func(worry int) int { return worry + 1 },
		func(worry int) int { return worry * 13 },
		func(worry int) int { return worry * worry },
		func(worry int) int { return worry + 7 },
		func(worry int) int { return worry + 8 },
		func(worry int) int { return worry + 4 },
		func(worry int) int { return worry + 5 },
	}
	testCondition := []int{13, 3, 7, 2, 19, 5, 11, 17}
	testTrue := []int{6, 7, 1, 6, 5, 3, 1, 3}
	testFalse := []int{2, 4, 4, 0, 7, 0, 2, 5}

	// Get LCM of test conditions
	lcm := LCM(1, 1, testCondition...)

	inspections := make([]int, monkeys+1)
	round := 0
	for round < 10000 {
		monkey := 0
		for monkey <= monkeys {
			// Loop through items that monkey has
			for _, oldWorry := range items[monkey] {
				// count inspection
				inspections[monkey] += 1

				// inspect item
				worry := operations[monkey](oldWorry)
				// decide which monkey receives item
				var receivingMonkey int
				if worry%testCondition[monkey] == 0 {
					receivingMonkey = testTrue[monkey]
				} else {
					receivingMonkey = testFalse[monkey]
				}
				worry = worry % lcm
				// give item to new monkey
				items[receivingMonkey] = append(items[receivingMonkey], worry)
				// remove item from current monkey
				items[monkey] = items[monkey][1:]
			}
			monkey += 1
		}
		round += 1
	}
	maxInspections := 0
	secondMaxInspections := 0
	for _, i := range inspections {
		if i > maxInspections {
			secondMaxInspections = maxInspections
			maxInspections = i
		} else if i > secondMaxInspections {
			secondMaxInspections = i
		}
	}
	return maxInspections * secondMaxInspections
}

// GCD greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// https://adventofcode.com/2022/day/11
func main() {
	// 119715
	answerOne := partOne11()
	println(answerOne)

	// 18085004878
	answerTwo := partTwo11()
	println(answerTwo)
}
