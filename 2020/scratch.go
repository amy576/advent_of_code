package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Instruction struct {
	line string
	arg string
	increment int
	executed bool
}

func readInput(filename string) []Instruction {
	// filename := "day8/day8_test.csv"
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	// fmt.Println(lines)

	var records []Instruction
	for _, line := range lines {
		inc, _ := strconv.Atoi(line[4:])
		i := Instruction{line, line[:3], inc, false}
		// fmt.Println(line)
		// fmt.Println(line[:3])
		// fmt.Println(inc)
		// fmt.Println(i)
		records = append(records, i)
	}

	return records
}

func executeInstructions(records []Instruction, start int) int {
	// this will take in a line to start at and return the accumulation
	// it will break when we've seen this instruction already
	step := start
	accumulation := 0
	// fmt.Println(records)
	// fmt.Println(step)
	for records[step].executed == false && step > -1 && step < len(records) {
		instruc := records[step]
		// fmt.Println(instruc)
		records[step].executed = true
		if instruc.arg == "acc" {
			accumulation += instruc.increment
			step ++
		} else if instruc.arg == "nop" {
			step ++
		} else if instruc.arg == "jmp" {
			step += instruc.increment
		} else {
			fmt.Println(instruc)
		}
	}
	return accumulation
}

func main() {
	filename := "day8/day8_test.csv"
	records := readInput(filename)
	q1 := executeInstructions(records, 0)
	fmt.Println("part 1: ", q1)
}