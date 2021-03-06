package main

import (
    "bufio"
    "fmt"
	"os"
)

type Cell struct {
	x int
	y int
	z int
	w int
}

var defaultState = "inactive"

func readInput(filename string) map[Cell]string {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	cells := make(map[Cell]string)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	for i, row := range lines {
		for j, char := range row {
			// initialize one row before and one row after
			// initialize one column before and one after
			for w := -1; w <= 1; w++ {
				for z := -1; z <= 1; z++ {
					if i == 0 {
						cells[Cell{i-1, j, z, w}] = defaultState
					} else if i == len(lines) - 1 {
						cells[Cell{i+1, j, z, w}] = defaultState
					}
		
					if j == 0 {
						if i == 0 {
							cells[Cell{i-1, j-1, z, w}] = defaultState
						} else if i == len(lines) - 1 {
							cells[Cell{i+1, j-1, z, w}] = defaultState
						}
						cells[Cell{i, j-1, z, w}] = defaultState
					} else if j == len(row) - 1 {
						if i == 0 {
							cells[Cell{i-1, j+1, z, w}] = defaultState
						} else if i == len(lines) - 1 {
							cells[Cell{i+1, j+1, z, w}] = defaultState
						}
						cells[Cell{i, j+1, z, w}] = defaultState
					}

					if z == 0 && w == 0 {
					} else {
						cells[Cell{i, j, z, w}] = defaultState
					}
				}
			}

			var state string
			if string(char) == "." {
				state = "inactive"
			} else if string(char) == "#" {
				state = "active"
			}
			cell := Cell{i, j, 0, 0}
			cells[cell] = state
		}
	}
	return cells
}

func change(startGrid map[Cell]string) (map[Cell]string) {
	endGrid := make(map[Cell]string)

	for coordinates, state := range startGrid {
		// Each cube only ever considers its neighbors: any of
		// the 26 other cubes where any of their coordinates differ
		// by at most 1. 
		xCoord, yCoord, zCoord := coordinates.x, coordinates.y, coordinates.z
		var adjacent int
		// fmt.Println("neighbors of ", xCoord, yCoord, zCoord)
		for i := xCoord - 1; i <= xCoord + 1; i++ {
			for j := yCoord - 1; j <= yCoord + 1; j++ {
				for k := zCoord - 1; k <= zCoord + 1; k++ {
					// skip this actual coordinate
					if i == xCoord && j == yCoord && k == zCoord {
					} else {
						// fmt.Println(i, j, k)
						adjacentCell := Cell{i, j, k, 0}
						adjacentStatus, exists := startGrid[adjacentCell]
						_, alreadyInitialized := endGrid[adjacentCell]
						if exists {
							if adjacentStatus == "active" {
								adjacent++
							}
						} else if alreadyInitialized {
						// } else if state == "active" {
						} else {
							// the space is technically infinite
							// if we come across something where there is no
							// coordinate there in the 26 adjacent cells, we
							// need to initialize one as an inactive cell
							// we should only need to do this for cells adjacent
							// to at least one active cell, since those are the only
							// cells that have any chance of flipping to active
							endGrid[adjacentCell] = defaultState
						}
					}
				}
			}
		}

		// During a cycle, all cubes simultaneously change their state
		// according to the following rules:
		// If a cube is active and exactly 2 or 3 of its neighbors are
		// also active, the cube remains active. Otherwise, the cube
		// becomes inactive.
		// If a cube is inactive but exactly 3 of its neighbors are
		// active, the cube becomes active. Otherwise, the cube remains
		// inactive.
		if state == "active" {
			if adjacent == 2 || adjacent == 3 {
				endGrid[coordinates] = "active"
			} else {
				endGrid[coordinates] = "inactive"
			}
		} else if state == "inactive" {
			if adjacent == 3 {
				endGrid[coordinates] = "active"
			} else {
				endGrid[coordinates] = "inactive"
			}
		}

	}

	return endGrid
}

func change4d(startGrid map[Cell]string) (map[Cell]string) {
	endGrid := make(map[Cell]string)

	for coordinates, state := range startGrid {
		// Each cube only ever considers its neighbors: any of
		// the 26 other cubes where any of their coordinates differ
		// by at most 1. 
		xCoord, yCoord, zCoord, wCoord := coordinates.x, coordinates.y, coordinates.z, coordinates.w
		var adjacent int
		// fmt.Println("neighbors of ", xCoord, yCoord, zCoord)
		for i := xCoord - 1; i <= xCoord + 1; i++ {
			for j := yCoord - 1; j <= yCoord + 1; j++ {
				for k := zCoord - 1; k <= zCoord + 1; k++ {
					for l := wCoord - 1; l <= wCoord + 1; l++ {
						// skip this actual coordinate
						if i == xCoord && j == yCoord && k == zCoord && l == wCoord {
						} else {
							// fmt.Println(i, j, k)
							adjacentCell := Cell{i, j, k, l}
							adjacentStatus, exists := startGrid[adjacentCell]
							_, alreadyInitialized := endGrid[adjacentCell]
							if exists {
								if adjacentStatus == "active" {
									adjacent++
								}
							} else if alreadyInitialized {
							// } else if state == "active" {
							} else {
								// the space is technically infinite
								// if we come across something where there is no
								// coordinate there in the 26 adjacent cells, we
								// need to initialize one as an inactive cell
								// we should only need to do this for cells adjacent
								// to at least one active cell, since those are the only
								// cells that have any chance of flipping to active
								endGrid[adjacentCell] = defaultState
							}
						}
					}
				}
			}
		}

		// During a cycle, all cubes simultaneously change their state
		// according to the following rules:
		// If a cube is active and exactly 2 or 3 of its neighbors are
		// also active, the cube remains active. Otherwise, the cube
		// becomes inactive.
		// If a cube is inactive but exactly 3 of its neighbors are
		// active, the cube becomes active. Otherwise, the cube remains
		// inactive.
		if state == "active" {
			if adjacent == 2 || adjacent == 3 {
				endGrid[coordinates] = "active"
			} else {
				endGrid[coordinates] = "inactive"
			}
		} else if state == "inactive" {
			if adjacent == 3 {
				endGrid[coordinates] = "active"
			} else {
				endGrid[coordinates] = "inactive"
			}
		}

	}

	return endGrid
}

func countActive(countGrid map[Cell]string) int {
	var active int
	for _, status := range countGrid {
		if status == "active" {
			active++
		}
	}
	return active
}

func q1(grid map[Cell]string, numIter int) int {
	runGrid := grid
	for cycle := 0; cycle < numIter; cycle++ {
		runGrid = change(runGrid)
		fmt.Println(cycle)
		fmt.Println(countActive(runGrid))
		// fmt.Println(runGrid)
	}

	return countActive(runGrid)
}

func q2(grid map[Cell]string, numIter int) int {
	runGrid := grid
	for cycle := 0; cycle < numIter; cycle++ {
		runGrid = change4d(runGrid)
		fmt.Println(cycle)
		fmt.Println(countActive(runGrid))
		// fmt.Println(runGrid)
	}

	return countActive(runGrid)
}

func main() {
	filename := "day17_input.csv"
	startingCells := readInput(filename)
	iterations := 6
	fmt.Println(startingCells)

	// Starting with your given initial configuration, simulate six cycles.
	// How many cubes are left in the active state after the sixth cycle?
	q1 := q1(startingCells, iterations)
	fmt.Println("part 1: ", q1)

	// Starting with your given initial configuration, simulate six cycles
	// in a 4-dimensional space. How many cubes are left in the active state
	// after the sixth cycle? 
	q2 := q2(startingCells, iterations)
	fmt.Println("part 2: ", q2)
}