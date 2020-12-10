package main

import (
    "fmt"
    "os"
	"bufio"
	"strconv"
	"sort"
	// "math"
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

func q1(startValue int, inputs []int) int {
	// sort the values in order from smallest to biggest
	sortedValues := make([]int, len(inputs))
	copy(sortedValues, inputs)
	sort.Ints(sortedValues)
	fmt.Println(sortedValues)

	diffs := make(map[int]int)
	
	// there is always at least 1 3-jolt difference
	// because your device's built-in adapter is always 3
	// higher than the highest adapter
	diffs[3] = 1

	// start making the chain
	for i, adapter := range sortedValues {
		diff := adapter - startValue
		if diff >= 1 && diff <= 3 {
			diffs[diff]++
		} else {
			fmt.Println(i, adapter)
			if i != 0 {
				fmt.Println(i-1, startValue)
			}
			fmt.Println(diff)
		}
		startValue = adapter
	}

	fmt.Println(diffs)

	return diffs[1] * diffs[3]
}

func q2(startValue int, inputs []int) int {
	// inputs with 0
	var inputsPlus []int
	inputsPlus = append(inputsPlus, startValue)
	for _, val := range inputs {
		inputsPlus = append(inputsPlus, val)
	}
	device := findMax(inputs)+3
	inputsPlus = append(inputsPlus, device)
	
	// we'll count branches starting from the end, which is the
	// device adapter, because that cannot change
	// sort the values in order from biggest to smallest
	sortedValues := make([]int, len(inputsPlus))
	copy(sortedValues, inputsPlus)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedValues)))
	// fmt.Println(sortedValues)

	// make a map of paths from each node, again going backwards
	paths := make(map[int]int)
	paths[device] = 1 // only one option from here: the max adapter
	for i := 1; i < len(sortedValues); i++ {
		// for the previous nodes that connected to this node (e.g. 
		// within 3), and add their number of paths to this nodes'
		// number of paths
		adapter := sortedValues[i]
		for j := 1; j <= i && sortedValues[i-j] - adapter <= 3; j++ {
			paths[adapter] += paths[sortedValues[i-j]]
		}
		// fmt.Println(paths[adapter])
	}
	// fmt.Println(paths)
	return paths[startValue]
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

func factorial(n int)(result int) {
	x := 1
	if (n > 0) {
		for i := 1; i <= n; i++ {
			// fmt.Println("x ",x,"n ",n)
			x *= i
		}
	}
	return x
}

// func makeChain(startValue int, sortedValues []int) bool {
// 	// check if a chain can be made
// 	var keptGoing int
// 	for _, adapter := range sortedValues {
// 		diff := adapter - startValue
// 		if diff > 3 {
// 			break
// 		} else {
// 			startValue = adapter
// 			keptGoing++
// 		}
// 	}
// 	if keptGoing == len(sortedValues) {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func main() {
	filename := "day10_input.csv"
	inputs := readInput(filename)
	start := 0

	// Find a chain that uses all of your adapters to connect
	// the charging outlet to your device's built-in adapter
	// and count the joltage differences between the charging outlet,
	// the adapters, and your device. What is the number of 1-jolt
	// differences multiplied by the number of 3-jolt differences?
	q1 := q1(start, inputs)

	// What is the total number of distinct ways you can arrange
	// the adapters to connect the charging outlet to your device?
	q2 := q2(start, inputs)

	fmt.Println("part 1: ", q1)
	fmt.Println("part 2: ", q2)
}