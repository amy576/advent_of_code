package main

import (
    "fmt"
    "os"
	"bufio"
	"strconv"
)

// acc increases or decreases a single global value
// called the accumulator by the value given in the argument.
// For example, acc +7 would increase the accumulator by 7.
// The accumulator starts at 0. After an acc instruction,
// the instruction immediately below it is executed next.
// jmp jumps to a new instruction relative to itself.
// The next instruction to execute is found using the
// argument as an offset from the jmp instruction; for
// example, jmp +2 would skip the next instruction,
// jmp +1 would continue to the instruction immediately
// below it, and jmp -20 would cause the instruction 20
// lines above to be executed next.
// nop stands for No OPeration - it does nothing.
// The instruction immediately below it is executed next.

// This is an infinite loop: with this sequence of jumps,
// the program will run forever. The moment the program tries
// to run any instruction a second time, you know it will never
// terminate.

// make a struct where all the rows (a.k.a. instructions) are set
// to false
type Instruction struct {
	line string
	arg string
	increment int
	executed bool
}

type InstructionOrder struct {
	line string
	arg string
	increment int
	originalIndex int
}

func readInput(filename string) []Instruction {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var records []Instruction
	for _, line := range lines {
		inc, _ := strconv.Atoi(line[4:])
		i := Instruction{line, line[:3], inc, false}
		records = append(records, i)
	}

	return records
}

func executeInstructionsQ1(recordList []Instruction, start int) (int, []InstructionOrder, int) {
	// this will take in a line to start at and return the accumulation
	// it will break when we've seen this instruction already
	step := start
	accumulation := 0

	// to save ourselves time in q2, let's only keep track of
	// EXECUTED jmp and nop orders; these should be the only
	// possible ones to change
	var executedInstructions []InstructionOrder
	var looped int

	newRecords := make([]Instruction, len(recordList))
	copy(newRecords, recordList)

	for step > -1 && step < len(newRecords) {
		if newRecords[step].executed == false {
			instruc := newRecords[step]
			// fmt.Println(instruc)
			newRecords[step].executed = true
			if instruc.arg == "acc" {
				accumulation += instruc.increment
				step ++
			} else if instruc.arg == "nop" {
				executed := InstructionOrder{instruc.line, instruc.arg, instruc.increment, step}
				executedInstructions = append(executedInstructions, executed)	
				step ++
			} else if instruc.arg == "jmp" {
				executed := InstructionOrder{instruc.line, instruc.arg, instruc.increment, step}
				executedInstructions = append(executedInstructions, executed)	
				step += instruc.increment
			} else {
				fmt.Println(instruc)
			}
		} else {
			looped++
			break
		}
	}
	return accumulation, executedInstructions, looped
}

func swapInstructionsQ2(executedRecords []InstructionOrder, recordSet []Instruction) int {
	var accumulation int

	// fmt.Println(recordSet)
	// i just want to know how many times it tried an option
	var tested int

	// go backwards since the problems seem to stem closer to the end
	for executed := len(executedRecords) - 1; executed >= 0; executed += -1 {
		tested++

		changedRecords := make([]Instruction, len(recordSet))
		copy(changedRecords, recordSet)

		ogIndex := executedRecords[executed].originalIndex
		if executedRecords[executed].arg == "nop" {
			changedRecords[ogIndex].arg = "jmp"
		} else if executedRecords[executed].arg == "jmp" {
			changedRecords[ogIndex].arg = "nop"
		}

		attempt, _, looped := executeInstructionsQ1(changedRecords,0)
		if looped == 0 {
			accumulation = attempt
			fmt.Println(executedRecords[executed])
			fmt.Println(tested, " out of ", len(executedRecords))
			break
		}
	}
	return accumulation
}

func main() {
	filename := "day8_input.csv"
	inputs := readInput(filename)

	// Immediately before any instruction is
	// executed a second time, what value is in the accumulator?
	q1, recordsPostLoop, _ := executeInstructionsQ1(inputs, 0)

	// Fix the program so that it terminates normally by changing
	// exactly one jmp (to nop) or nop (to jmp). What is the value
	// of the accumulator after the program terminates?
	q2 := swapInstructionsQ2(recordsPostLoop, inputs)
	fmt.Println("part 1: ", q1)
	fmt.Println("part 2: ", q2)
}