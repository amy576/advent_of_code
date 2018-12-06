package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

func q1(filename string) int {
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)

	var coordinates [][2]int

	// the infinite coordinates are the coordinates where one of the coordinates
	// is either the smallest or the largest
	smallest_x := 10000
	smallest_y := 10000
	var largest_x int
	var largest_y int

	for scanner.Scan() {
		line := scanner.Text()
		comma_pos := strings.Index(line,",")
		x, _ := strconv.Atoi(line[:comma_pos])
		y, _ := strconv.Atoi(line[comma_pos+2:])
		if x > largest_x {
			largest_x = x
		}
		if y > largest_y {
			largest_y = y
		}
		if x < smallest_x {
			smallest_x = x
		}
		if y < smallest_y {
			smallest_y = y
		}
		xy := [2]int{x,y}
		coordinates = append(coordinates,xy)
	}
	f.Close()

	// make a map that has max x and max y
	grid := make([][]string, largest_y+1)
	for i := range grid {
	    grid[i] = make([]string, largest_x+1)
	}
	// make a map of record number and number of locations (including the coordinate itself)
	areas := make(map[string]int)
	// for each spot on the map, iterate over all the coordinates to find the length
	// if length < shortest_length so far, write the record number of the coordinate in that spot
	// if length == shortest_length so far, write "." in that spot
	for y, inner := range grid {
		for x, _ := range inner {
			shortest_length := 1000
			var closest string
			for i, coord := range coordinates {
				// find Manhattan distance
				dist := int(math.Abs(float64(coord[0]-x)) + math.Abs(float64(coord[1]-y)))
				if dist == shortest_length {
					grid[y][x] = "."
				}
				if dist < shortest_length {
					grid[y][x] = strconv.Itoa(i)
					closest = strconv.Itoa(i)
					shortest_length = dist
				}
			}
			if x == 0 || y == 0 {
				areas[closest] = -1
			}
		}
	}

	for _, inner := range grid {
		for _, coord_val := range inner {
			if areas[coord_val] != -1 {
				areas[coord_val]++
			}
		}
	}

	// // check that I have as many filled in as are in the grid
	// var sum int
	// for _, val := range areas {
	// 	sum += val
	// }
	// fmt.Println(sum)
	// fmt.Println(len(grid)*len(grid[0]))

	// iterate over the map and find the largest number of locations
	var max_locations int
	for i, coord := range coordinates {
		if coord[0] == smallest_x || coord[0] == largest_x || coord[1] == smallest_y || coord[1] == largest_y {
			areas[strconv.Itoa(i)] = -1
		} else {
			if areas[strconv.Itoa(i)] > max_locations {
				max_locations = areas[strconv.Itoa(i)]
			}
		}
	}
	return max_locations
}

func q2(filename string) int {

	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)

	var coordinates [][2]int

	var largest_x int
	var largest_y int

	for scanner.Scan() {
		line := scanner.Text()
		comma_pos := strings.Index(line,",")
		x, _ := strconv.Atoi(line[:comma_pos])
		y, _ := strconv.Atoi(line[comma_pos+2:])
		if x > largest_x {
			largest_x = x
		}
		if y > largest_y {
			largest_y = y
		}
		xy := [2]int{x,y}
		coordinates = append(coordinates,xy)
	}
	f.Close()

	// make a map that has max x and max y
	grid := make([][]string, largest_y+1)
	for i := range grid {
	    grid[i] = make([]string, largest_x+1)
	}
	// for each spot on the map, iterate over all the coordinates to find the distances and add to a "total distance"
	// if total distance < 10000, increment "region size" by 1
	var region_size int
	for y, inner := range grid {
		for x, _ := range inner {
			var total_distance int
			for _, coord := range coordinates {
				// find Manhattan distance
				dist := int(math.Abs(float64(coord[0]-x)) + math.Abs(float64(coord[1]-y)))
				total_distance += dist
			}
			if total_distance < 10000 {
				region_size++
			}
		}
	}

	return region_size
}

func main() {
	fmt.Println("part1: ", q1("day6_input.txt"))
	fmt.Println("part2: ", q2("day6_input.txt"))
}
