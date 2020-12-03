package main

import (
    "fmt"
    "os"
	"bufio"
)

// Starting at the top-left corner of your map and following
// a slope of right 3 and down 1, how many trees would you encounter?
func q1(filename string) int {
    f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	
	trees := 0
	x := 0

	var records []string
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}

	numLines := len(records)
	fmt.Println("number of lines ", numLines)
	numCols := len(records[0])
	fmt.Println("number of base characters in a line ", numCols)

	// the actual record pattern repeats many times
	// we need x to be able to be 3x as large as y
	// for i, record := range records {
	// 	for z := 0; z < ((numLines * 3 / numCols)+1); z ++ {
	// 		records[i] += record
	// 	}
	// }

	// fmt.Println("successfully repeated patterns")

	for y := 0; y < numLines; y++ {
		currentChar := records[y][x:x+1]
		if currentChar == "#" {
			trees ++
		}
		x+=3
		if x >= numCols {
			x = x - numCols
		}
	}

	return trees
}

// Determine the number of trees you would encounter if,
// for each of the following slopes, you start at the top-left corner
// and traverse the map all the way to the bottom:
// Right 1, down 1.
// Right 3, down 1. (This is the slope you already checked.)
// Right 5, down 1.
// Right 7, down 1.
// Right 1, down 2.
// What do you get if you multiply together the number of trees encountered on each of the listed slopes?
func q2(filename string) int {
    f, _ := os.Open(filename)
    defer f.Close()
	scanner := bufio.NewScanner(f)

	var trees []int
	// var slopes [5][2]int
	slopes := make(map[int][2]int)
	slopes[1] = [2]int{1, 1}
	slopes[2] = [2]int{3, 1}
	slopes[3] = [2]int{5, 1}
	slopes[4] = [2]int{7, 1}
	slopes[5] = [2]int{1, 2}

	var records []string
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}

	numLines := len(records)
	fmt.Println("number of lines ", numLines)
	numCols := len(records[0])
	fmt.Println("number of base characters in a line ", numCols)

	for _, slope := range slopes {
		x := 0
		xIncr := slope[0]
		yIncr := slope[1]
		fmt.Println("current slope is ", xIncr, ", ", yIncr)
		runTrees := 0

		for y := 0; y < numLines; y+=yIncr {
			currentChar := records[y][x:x+1]
			if currentChar == "#" {
				runTrees ++
			}
			x+=xIncr
			if x >= numCols {
				x = x - numCols
			}
		}

		trees = append(trees,runTrees)
		fmt.Println(runTrees)
	}

	totalTrees := 1

	for _, number := range trees {
		totalTrees = totalTrees * number
	}

	return totalTrees
}

func main() {
    fmt.Println("part 1: ", q1("day3_input.csv"))
    fmt.Println("part 2: ", q2("day3_input.csv"))
}