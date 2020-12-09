package main

import (
    "fmt"
    "os"
	"bufio"
	"strconv"
)

func readInput(filename string) []int {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []int
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}

	return lines
}

func checkNumber(findSum int, numbers []int) bool {
	// this will check if a number is the sum of
	// 2 of the numbers in the given array
	var foundNumber bool
	for i := 0; i < len(numbers) - 1; i++ {
		for j := 1; j < len(numbers); j++ {
			if findSum == numbers[i] + numbers[j] {
				foundNumber = true
				break
			}
		}
		if foundNumber == true {
			break
		}
	}
	return foundNumber
}

func findContiguous(findSum int, numbers[]int) int {
	// return the array of contiguous numbers within
	// the numbers slice whose values result in findSum
	var contiguous []int
	for i, val := range numbers {
		total := val
		for j := i + 1; j < len(numbers); j++ {
			total += numbers[j]
			if total > findSum {
				break
			} else if total == findSum {
				contiguous = numbers[i:j+1]
			}
		}
	}
	return findMin(contiguous) + findMax(contiguous)
}

func findMin(set []int) int {
	var min int
	for i, e := range set {
		if i==0 || e < min {
			min = e
		}
	}
	return min
}

func findMax(set []int) int {
	var max int
	for i, e := range set {
		if i==0 || e > max {
			max = e
		}
	}
	return max
}

func q1(preamble int, records []int) int {
	var wrongNumber int

	// make an array that will "move"
	numbersToCheck := make([]int, preamble)
	copy(numbersToCheck, records[:preamble])
	y := 0

	for x := preamble; x < len(records); x++ {
		numbersToCheck := make([]int, preamble)
		copy(numbersToCheck, records[y:x])
		// fmt.Println(records[x])
		// fmt.Println(numbersToCheck)
		if checkNumber(records[x], numbersToCheck) {
			y++
		} else {
			wrongNumber = records[x]
			break
		}
	}
	return wrongNumber
}

func main() {
	filename := "day9_input.csv"
	inputs := readInput(filename)
	preamble := 25

	// The first step of attacking the weakness in the XMAS data is to
	// find the first number in the list (after the preamble) which is
	// not the sum of two of the 25 numbers before it. What is the first
	// number that does not have this property?
	q1 := q1(preamble, inputs)
	q2 := findContiguous(q1, inputs)


	fmt.Println("part 1: ", q1)
	fmt.Println("part 2: ", q2)
}