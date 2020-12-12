package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"math"
)

type Instruction struct {
	direction string
	movement int
}

func readInput(filename string) []Instruction {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []Instruction
	for scanner.Scan() {
		line := scanner.Text()
		amt, _ := strconv.Atoi(line[1:])
		instr := Instruction{line[:1], amt}
		lines = append(lines, instr)
	}
	return lines
}

func q1(directions []Instruction, initDirection int) int {
	facing := initDirection
	var x int
	var y int

	// The navigation instructions (your puzzle input) consists of a
	// sequence of single-character actions paired with integer input values.
	// Action N means to move north by the given value.
	// Action S means to move south by the given value.
	// Action E means to move east by the given value.
	// Action W means to move west by the given value.
	// Action L means to turn left the given number of degrees.
	// Action R means to turn right the given number of degrees.
	// Action F means to move forward by the given value in the direction the ship
	// is currently facing.
	for _, instruction := range directions {
		heading := instruction.direction
		squares := instruction.movement
		if heading == "L" || heading == "R" {
			facing = turn(facing, heading, squares)
		} else if heading == "N" || (heading == "F" && facing == 90) {
			x += squares
		} else if heading == "S" || (heading == "F" && facing == 270) {
			x -= squares
		} else if heading == "E" || (heading == "F" && facing == 0) {
			y += squares
		} else if heading == "W" || (heading == "F" && facing == 180) {
			y -= squares
		} else {
			fmt.Println(facing, instruction)
		}
		// fmt.Println(instruction, x, y)
	}

	manhattan := int(math.Abs(float64(x)) + math.Abs(float64(y)))

	return manhattan
}

func turn(startFacing int, direction string, degrees int) int {
	var newFacing int
	if direction == "L" {
		newFacing = startFacing + degrees
	} else if direction == "R" {
		newFacing = startFacing - degrees
	} else {
		fmt.Println(startFacing, direction, degrees)
	}

	newFacing = ((newFacing % 360) + 360) % 360
	// fmt.Println("facing new direction:", newFacing)

	return newFacing
}

func q2(directions []Instruction) int {
	var x int
	var y int
	xway := 10
	yway := 1

	// The navigation instructions (your puzzle input) consists of a
	// sequence of single-character actions paired with integer input values.
	// Action N means to move the waypoint north by the given value.
	// Action S means to move the waypoint south by the given value.
	// Action E means to move the waypoint east by the given value.
	// Action W means to move the waypoint west by the given value.
	// Action L means to rotate the waypoint around the ship left (counter-clockwise)
	// the given number of degrees.
	// Action R means to rotate the waypoint around the ship right (clockwise) the given
	// number of degrees.
	// Action F means to move forward to the waypoint a number of times equal to the given value.
	for _, instruction := range directions {
		heading := instruction.direction
		squares := instruction.movement
		if heading == "L" || heading == "R" {
			xway, yway = rotate(xway, yway, x, y, heading, squares)
		} else if heading == "N" {
			yway += squares
		} else if heading == "S" {
			yway -= squares
		} else if heading == "E" {
			xway += squares
		} else if heading == "W" {
			xway -= squares
		} else if heading == "F" {
			newx := x + ((xway-x) * squares)
			newy := y + ((yway-y) * squares)
			xway = newx + (xway-x)
			yway = newy + (yway-y)
			x = newx
			y = newy
		} else {
			fmt.Println(instruction)
		}
		// fmt.Println(instruction, x, y, "with waypoint at", xway, yway)
	}

	manhattan := int(math.Abs(float64(x)) + math.Abs(float64(y)))

	return manhattan
}

func rotate(initxway int, inityway int, x int, y int, direction string, degrees int) (int, int) {
	newxway := initxway
	newyway := inityway
	diffx := initxway - x
	diffy := inityway - y
	if (degrees == 90 && direction == "L") || (degrees == 270 && direction == "R") {
		newxway = x - diffy
		newyway = y + diffx
	} else if (degrees == 180) {
		newxway = x - diffx
		newyway = y - diffy
	} else if (degrees == 270 && direction == "L") || (degrees == 90 && direction == "R") {
		newxway = x + diffy
		newyway = y - diffx
	} else {
		fmt.Println(initxway, inityway, direction, degrees)
	}
	// fmt.Println(newxway, newyway, direction, degrees, "to", newxway, newyway)
	return newxway, newyway
}

func main() {
	filename := "day12_input.csv"
	inputs := readInput(filename)
	startingDirection := 0 // 0 == east

	// Figure out where the navigation instructions lead.
	// What is the Manhattan distance between that location and the
	// ship's starting position?
	q1 := q1(inputs, startingDirection)
	q2 := q2(inputs)

	fmt.Println("part 1: ", q1)
	fmt.Println("part 2: ", q2)
}