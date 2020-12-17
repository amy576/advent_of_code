package main

import (
    "bufio"
    "fmt"
	"os"
)

type Cell struct {
	state string
	x int
	y int
	z int
}

func readInput(filename string) []Cell {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	var cells []Cell
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	for i, row := range lines {
		for j, char := range row {
			var state string
			if string(char) == "." {
				state = "inactive"
			} else if string(char) == "#" {
				state = "active"
			}
			cell := Cell{state, i, j, 0}
			cells = append(cells, cell)
		}
	}
	return cells
}

func main() {
	filename := "day17_test.csv"
	startingCells := readInput(filename)
	fmt.Println(startingCells)

	// Adding together all of the invalid values produces your ticket scanning
	// error rate. Consider the validity of the nearby tickets you scanned.
	// What is your ticket scanning error rate
	// q1, validTickets := q1(fieldValues, nearbyTickets)
	// fmt.Println("part 1: ", q1)

	// Given your starting numbers, what will be the 30000000th number spoken? 
	// q2 := q2(fieldValues, validTickets, myTicket)
	// fmt.Println("part 2: ", q2)
}