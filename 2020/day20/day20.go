package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	// "strings"
)

type Border struct {
	border string
	matched bool
}

type Tile struct {
	// id int
	top Border
	bottom Border
	left Border
	right Border
	setBorders int
}

func readInput(filename string) (map[int][]string, map[int]Tile) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	bordersOnly := make(map[int]Tile)
	fullTiles := make(map[int][]string)

	var id int
	var tileTmp []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			if line[:1] == "T" {
				if id != 0 {
					fullTiles[id] = tileTmp
				}
				id, _ = strconv.Atoi(line[5:len(line)-1])
				tileTmp = []string{}
			} else {
				tileTmp = append(tileTmp, line)
			}
		}
	}

	for tileId, tileHash := range fullTiles {
		topBorder := tileHash[0]
		bottomBorder := tileHash[len(tileHash)-1]
		var leftBorder string
		var rightBorder string
		for row := 0; row < len(tileHash); row++ {
			leftBorder += string(tileHash[row][0])
			rightBorder += string(tileHash[row][len(tileHash)-1])
		}
		bordersOnly[tileId] = Tile{Border{topBorder, false}, Border{bottomBorder, false}, Border{leftBorder, false}, Border{rightBorder, false}, 0}
	}

	return fullTiles, bordersOnly
}

func checkMatch(border1 Border, border2 Border, tile1 Tile, tile2 Tile) (Border, Border, int, int) {
	if border1.matched == false && border2.matched == false && border1.border == border2.border {
		border1.matched = true
		border2.matched = true
		tile1.setBorders++
		tile2.setBorders++
		fmt.Println("matched a border: ", border1, border2, tile1.setBorders, tile2.setBorders)
	}
	return border1, border2, tile1.setBorders, tile2.setBorders
}

func rotate(tile Tile) Tile {
	// rotate the tile clockwise 90 degrees
	// first rotation: top becomes left, right becomes top, bottom becomes right, and left becomes bottom
	// top and bottom need to go "backwards" then
	var backwardsTop string
	for _, char := range tile.top.border {
		backwardsTop = string(char) + backwardsTop
	}
	newTop := Border{backwardsTop, tile.top.matched}
	var backwardsBottom string
	for _, char := range tile.bottom.border {
		backwardsBottom = string(char) + backwardsBottom
	}
	newBottom := Border{backwardsBottom, tile.bottom.matched}
	newTile := Tile{tile.right, tile.left, newTop, newBottom, tile.setBorders}
	fmt.Println(tile)
	fmt.Println(newTile)
	return newTile
	// second rotation: original top is now bottom, original right is now left, original bottom is now top, and original left is now right
	// third rotation: original top is now right, original right is now bottom, original bottom is now left, original left is now top
}

func q1(borders map[int]Tile) int {
	// By rotating, flipping, and rearranging them, you can find a square
	// arrangement that causes all adjacent borders to line up.
	for tileId, bordersToCheck := range borders {
		// check each border while it is unset
		for checkTile, checkBorder := range borders {
			if tileId != checkTile {
				// bordersToCheck.bottom, checkBorder.top, bordersToCheck.setBorders, checkBorder.setBorders = checkMatch(bordersToCheck.bottom, checkBorder.top, bordersToCheck, checkBorder)
				// bordersToCheck.top, checkBorder.bottom, bordersToCheck.setBorders, checkBorder.setBorders = checkMatch(bordersToCheck.bottom, checkBorder.bottom, bordersToCheck, checkBorder)
				// bordersToCheck.left, checkBorder.right, bordersToCheck.setBorders, checkBorder.setBorders = checkMatch(bordersToCheck.bottom, checkBorder.right, bordersToCheck, checkBorder)
				// bordersToCheck.right, checkBorder.left, bordersToCheck.setBorders, checkBorder.setBorders = checkMatch(bordersToCheck.bottom, checkBorder.left, bordersToCheck, checkBorder)

				// rotatingTile := checkBorder
				for r := 0; r < 4; r++ {
					checkBorder = rotate(checkBorder)
					bordersToCheck.bottom, checkBorder.top, bordersToCheck.setBorders, checkBorder.setBorders = checkMatch(bordersToCheck.bottom, checkBorder.top, bordersToCheck, checkBorder)
					bordersToCheck.top, checkBorder.bottom, bordersToCheck.setBorders, checkBorder.setBorders = checkMatch(bordersToCheck.bottom, checkBorder.bottom, bordersToCheck, checkBorder)
					bordersToCheck.left, checkBorder.right, bordersToCheck.setBorders, checkBorder.setBorders = checkMatch(bordersToCheck.bottom, checkBorder.right, bordersToCheck, checkBorder)
					bordersToCheck.right, checkBorder.left, bordersToCheck.setBorders, checkBorder.setBorders = checkMatch(bordersToCheck.bottom, checkBorder.left, bordersToCheck, checkBorder)
				}
			}
		}
	}

	var corners int

	for tile, borders := range borders {
		if borders.setBorders == 2 {
			fmt.Println(tile)
			corners *= tile
		}
	}

	return corners
}

func main() {
	filename := "day20_test.csv"
	_, tileBorders := readInput(filename)
	// for key, elem := range tileBorders {
	// 	fmt.Println("tile", key)
	// 	fmt.Println(elem.top)
	// 	fmt.Println(elem.bottom)
	// 	fmt.Println(elem.right)
	// 	fmt.Println(elem.left)
	// 	fmt.Println("")
	// }

	// Assemble the tiles into an image. What do you get if you multiply
	// together the IDs of the four corner tiles?
	q1ans := q1(tileBorders)
	fmt.Println("part 1: ", q1ans)

	// q2 := count(rules, messages, solveFor)
	// fmt.Println("part 2: ", q2)
}