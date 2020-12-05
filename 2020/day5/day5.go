package main

import (
    "fmt"
    "os"
	"bufio"
	// "strings"
	// "regexp"
	// "strconv"
)

// The first 7 characters will either be F or B;
// these specify exactly one of the 128 rows on the plane
// (numbered 0 through 127). Each letter tells you which half
// of a region the given seat is in. Start with the whole list
// of rows; the first letter indicates whether the seat is
// in the front (0 through 63) or the back (64 through 127).
// The next letter indicates which half of that region the
// seat is in, and so on until you're left with exactly one row.
// The last three characters will be either L or R; these
// specify exactly one of the 8 columns of seats on the plane
// (numbered 0 through 7). The same process as above proceeds
// again, this time with only three steps. L means to keep the
// lower half, while R means to keep the upper half.
// Every seat also has a unique seat ID: multiply the row by 8,
// then add the column. In this example, the seat has ID 44 * 8 + 5 = 357.
// What is the highest seat ID on a boarding pass?
func q1(filename string) int {
    f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var seatIds []int
	totalRows := 128
	totalCols := 8
	var cols []int
	var rows []int
	for x := 0; x < totalCols; x++ {
		cols = append(cols,x)
	}
	for y := 0; y < totalRows; y++ {
		rows = append(rows,y)
	}

	var records []string
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}

	for _, record := range records {
		var col int
		var row int

		validCol := cols
		colId := record[7:] // just the RLR bit
		for _, colChar := range colId {
			if string(colChar) == "R" {
				validCol = validCol[len(validCol)/2:]
			} else if string(colChar) == "L" {
				validCol = validCol[:len(validCol)/2]
			} else {
				fmt.Println("not R or L: ",string(colChar), " in ",record)
			}
		}

		if len(validCol) == 1 {
			col = validCol[0]
			// fmt.Println(col)
		} else {
			fmt.Println("didn't get down to just one number, left with ",validCol)
		}

		validRow := rows
		rowId := record[:7] // just the FBF bit
		for _, rowChar := range rowId {
			if string(rowChar) == "B" {
				validRow = validRow[len(validRow)/2:]
			} else if string(rowChar) == "F" {
				validRow = validRow[:len(validRow)/2]
			} else {
				fmt.Println("not B or F: ",string(rowChar), " in ",record)
			}
		}

		if len(validRow) == 1 {
			row = validRow[0]
			// fmt.Println(col)
		} else {
			row = 0
			fmt.Println("didn't get down to just one number, left with ",validRow)
		}

		id := (row * 8) + col
		// fmt.Println(id)
		seatIds = append(seatIds, id)
	}

	var maxId int

	for i, seatId := range seatIds {
		if i==0 || seatId > maxId {
			maxId = seatId
		}
	}

	return maxId
}

// It's a completely full flight, so your seat should be the
// only missing boarding pass in your list. However, there's
// a catch: some of the seats at the very front and back of the
// plane don't exist on this aircraft, so they'll be missing from
// your list as well. Your seat wasn't at the very front or back,
// though; the seats with IDs +1 and -1 from yours will be in your list.
// What is the ID of your seat?
func q2(filename string) int {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var records []string
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}

	totalRows := 128
	totalCols := 8
	// make a map of the airplane
	plane := make([][]int, totalRows)
	for y := 0; y < totalRows; y++ {
		plane[y] = make([]int, totalCols)
	}
	var seatIds []int

	// fill in the plane
	var cols []int
	var rows []int
	for x := 0; x < totalCols; x++ {
		cols = append(cols,x)
	}
	for y := 0; y < totalRows; y++ {
		rows = append(rows,y)
	}

	for _, record := range records {
		var col int
		var row int

		validCol := cols
		colId := record[7:] // just the RLR bit
		for _, colChar := range colId {
			if string(colChar) == "R" {
				validCol = validCol[len(validCol)/2:]
			} else if string(colChar) == "L" {
				validCol = validCol[:len(validCol)/2]
			} else {
				fmt.Println("not R or L: ",string(colChar), " in ",record)
			}
		}

		if len(validCol) == 1 {
			col = validCol[0]
			// fmt.Println(col)
		} else {
			fmt.Println("didn't get down to just one number, left with ",validCol)
		}

		validRow := rows
		rowId := record[:7] // just the FBF bit
		for _, rowChar := range rowId {
			if string(rowChar) == "B" {
				validRow = validRow[len(validRow)/2:]
			} else if string(rowChar) == "F" {
				validRow = validRow[:len(validRow)/2]
			} else {
				fmt.Println("not B or F: ",string(rowChar), " in ",record)
			}
		}

		if len(validRow) == 1 {
			row = validRow[0]
			// fmt.Println(col)
		} else {
			row = 0
			fmt.Println("didn't get down to just one number, left with ",validRow)
		}

		id := (row * 8) + col
		seatIds = append(seatIds, id)

		plane[row][col] = id
	}

	var myId int

	for y := 1; y < totalRows - 1; y++ { // not the first or last row
		for x := 0; x < totalCols; x++ {
			if plane[y][x] == 0 { // if the seat is empty
				maybeMine := (y * 8) + x
				if intInSlice(maybeMine-1, seatIds) && intInSlice(maybeMine+1, seatIds) {
					myId = maybeMine
				}
			}
		}
	}

	return myId
}

func intInSlice(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


func main() {
    fmt.Println("part 1: ", q1("day5_input.csv"))
    fmt.Println("part 2: ", q2("day5_input.csv"))
}