package main

import (
    "bufio"
    "fmt"
    "os"
)

func readInput(filename string) [][]string {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines [][]string
	for scanner.Scan() {
		line := scanner.Text()
		var chars []string
		for _, char := range line {
			chars = append(chars, string(char))
		}
		lines = append(lines, chars)
	}
	return lines
}

func change(startGrid [][]string) ([][]string, int) {
	rows := len(startGrid)
	seatsPerRow := len(startGrid[0])
	endGrid := make([][]string, rows)
	for i := 0; i < rows; i++ {
		endGrid[i] = make([]string, seatsPerRow)
	}
	var changedSeats int
	// All decisions are based on the number of occupied seats
	// adjacent to a given seat (one of the eight positions immediately
	// up, down, left, right, or diagonal from the seat). The following
	// rules are applied to every seat simultaneously:
	// If a seat is empty (L) and there are no occupied seats adjacent
	// to it, the seat becomes occupied.
	// If a seat is occupied (#) and four or more seats adjacent to it
	// are also occupied, the seat becomes empty.
	// Otherwise, the seat's state does not change.
	for x, row := range startGrid {
		for y, seat := range row {
			var adjacent int
			// fmt.Println("checking upper left")
			if x > 0 && y > 0 {
				if startGrid[x-1][y-1] == "#" {
					adjacent++
				}
			}
			// fmt.Println("checking upper center")
			if x > 0 {
				if startGrid[x-1][y] == "#" {
					adjacent++
				}
			}
			// fmt.Println("checking upper right")
			if x > 0 && y < seatsPerRow - 1 {
				if startGrid[x-1][y+1] == "#" {
					adjacent++
				}
			}
			// fmt.Println("checking center left")
			if y > 0 {
				if startGrid[x][y-1] == "#" {
					adjacent++
				}
			}
			// fmt.Println("checking center right")
			if y < seatsPerRow - 1 {
				if startGrid[x][y+1] == "#" {
					adjacent++
				}
			}
			// fmt.Println("checking lower left")
			if x < rows - 1 && y != 0 {
				if startGrid[x+1][y-1] == "#" {
					adjacent++
				}
			}
			// fmt.Println("checking lower center")
			if x < rows - 1 {
				if startGrid[x+1][y] == "#" {
					adjacent++
				}
			}
			// fmt.Println("checking lower right")
			if x < rows - 1 && y < seatsPerRow - 1 {
				if startGrid [x+1][y+1] == "#" {
					adjacent++
				}
			}
			if adjacent == 0 && seat == "L" {
				endGrid[x][y] = "#"
				changedSeats++
				// fmt.Println("seat in position",x,y,"was unoccupied and is now occupied")
			} else if adjacent >= 4 && seat == "#" {
				endGrid[x][y] = "L"
				changedSeats++
				// fmt.Println("seat in position",x,y,"was occupied and is now unoccupied")
			} else {
				endGrid[x][y] = seat
				// fmt.Println("seat in position",x,y,"remains unchanged")
			}
		}
	}

	return endGrid, changedSeats
}

func countOccupied(countGrid [][]string) int {
	var occupied int
	for _, x := range countGrid {
		for _, y := range x {
			if y == "#" {
				occupied++
			}
		}
	}
	return occupied
}

func q1(grid [][]string) int {
	newGrid, numSeatsChanged := change(grid)
	changed := true
	numChanges := 1
	for changed {
		if numSeatsChanged > 0 {
			numChanges++
			grid = newGrid
			newGrid, numSeatsChanged = change(grid)
		} else if numSeatsChanged == 0 {
			changed = false
			// fmt.Println(newGrid)
		} else {
			fmt.Println("numSeatsChanged is negative")
			break
		}
	}

	return countOccupied(newGrid)
}

func q2(grid [][]string) int {
	newGrid, numSeatsChanged := changeNewRules(grid)
	changed := true
	numChanges := 1
	for changed {
		if numSeatsChanged > 0 {
			numChanges++
			grid = newGrid
			newGrid, numSeatsChanged = changeNewRules(grid)
		} else if numSeatsChanged == 0 {
			changed = false
			// fmt.Println(newGrid)
		} else {
			fmt.Println("numSeatsChanged is negative")
			break
		}
	}

	return countOccupied(newGrid)
}

func changeNewRules(startGrid [][]string) ([][]string, int) {
	rows := len(startGrid)
	seatsPerRow := len(startGrid[0])
	endGrid := make([][]string, rows)
	for i := 0; i < rows; i++ {
		endGrid[i] = make([]string, seatsPerRow)
	}
	var changedSeats int
	// Now, instead of considering just the eight immediately adjacent seats,
	// consider the first seat in each of those eight directions.
	// If a seat is empty (L) and there are no occupied seats adjacent
	// to it, the seat becomes occupied.
	// If a seat is occupied (#) and five or more seats adjacent to it
	// are also occupied, the seat becomes empty.
	// Otherwise, the seat's state does not change.
	for x, row := range startGrid {
		for y, seat := range row {
			var adjacent int
			// fmt.Println("checking upper left")
			LoopUL:
				for i := 1; x - i >= 0 && y - i >= 0; i++ {
					switch startGrid[x-i][y-i] {
					case "#":
						adjacent++
						break LoopUL
					case "L":
						break LoopUL
					}
				}
			// fmt.Println("checking upper center")
			LoopUC:
				for i := 1; x - i >= 0; i++ {
					switch startGrid[x-i][y] {
					case "#":
						adjacent++
						break LoopUC
					case "L":
						break LoopUC
					}
				}
			// fmt.Println("checking upper right")
			LoopUR:
				for i := 1; x - i >= 0 && y + i <= seatsPerRow - 1; i++ {
					switch startGrid[x-i][y+i] {
					case "#":
						adjacent++
						break LoopUR
					case "L":
						break LoopUR
					}
				}
			// fmt.Println("checking center left")
			LoopCL:
				for i := 1; y - i >= 0; i++ {
					switch startGrid[x][y-i] {
					case "#":
						adjacent++
						break LoopCL
					case "L":
						break LoopCL
					}
				}
			// fmt.Println("checking center right")
			LoopCR:
				for i := 1; y + i <= seatsPerRow - 1; i++ {
					switch startGrid[x][y+i] {
					case "#":
						adjacent++
						break LoopCR
					case "L":
						break LoopCR
					}
				}
			// fmt.Println("checking lower left")
			LoopLL:
				for i := 1; x + i <= rows - 1 && y - i >= 0; i++ {
					switch startGrid[x+i][y-i] {
					case "#":
						adjacent++
						break LoopLL
					case "L":
						break LoopLL
					}
				}
			// fmt.Println("checking lower center")
			LoopLC:
				for i := 1; x + i <= rows - 1; i++ {
					switch startGrid[x+i][y] {
					case "#":
						adjacent++
						break LoopLC
					case "L":
						break LoopLC
					}
				}
			// fmt.Println("checking lower right")
			LoopLR:
				for i := 1; x + i <= rows - 1 && y + i <= seatsPerRow - 1; i++ {
					switch startGrid [x+i][y+i] {
					case "#":
						adjacent++
						break LoopLR
					case "L":
						break LoopLR			
					}
				}
			if adjacent == 0 && seat == "L" {
				endGrid[x][y] = "#"
				changedSeats++
				// fmt.Println("seat in position",x,y,"was unoccupied and is now occupied")
			} else if adjacent >= 5 && seat == "#" {
				endGrid[x][y] = "L"
				changedSeats++
				// fmt.Println("seat in position",x,y,"was occupied and is now unoccupied")
			} else {
				endGrid[x][y] = seat
				// fmt.Println("seat in position",x,y,"remains unchanged")
			}
		}
	}
	// fmt.Println(endGrid)

	return endGrid, changedSeats
}



func main() {
	filename := "day11_input.csv"
	inputs := readInput(filename)

	// Simulate your seating area by applying the seating rules
	// repeatedly until no seats change state. How many seats end
	// up occupied?
	q1 := q1(inputs)
	q2 := q2(inputs)

	fmt.Println("part 1: ", q1)
	fmt.Println("part 2: ", q2)
}