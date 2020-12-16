package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
)

type NumberTurns struct {
	turnLastSpoken int
	turnLastSpokenPlusOne int
}

func readInput(filename string) []int {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []int
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line,",")
		for _, number := range numbers {
			num, _ := strconv.Atoi(number)
			lines = append(lines, num)
		}
	}
	return lines
}

func q1(startNums []int, stop int) int {
	var turns []int
	turns = append(turns, 0)
	numbersSpoken := make(map[int]NumberTurns)
	for inc := 1; inc <= stop; inc++ {
		// In this game, the players take turns saying numbers. They begin by
		// taking turns reading from a list of starting numbers (your puzzle input).
		if inc <= len(startNums) {
			turns = append(turns, startNums[inc-1])
			numbersSpoken[startNums[inc-1]] = NumberTurns{inc, 0}
		} else {
			// Then, each turn consists of considering the most recently spoken number:
			lastNumberSpoken := turns[inc-1]
			var nowSpeak int
			lastSpokenTurns, alreadySpoken := numbersSpoken[lastNumberSpoken]

			// If that was the first time the number has been spoken, the current player says 0.
			// Otherwise, the number had been spoken before; the current player announces how
			// many turns apart the number is from when it was previously spoken.
		
			if alreadySpoken && lastSpokenTurns.turnLastSpokenPlusOne > 0 {
				nowSpeak = lastSpokenTurns.turnLastSpoken - lastSpokenTurns.turnLastSpokenPlusOne
			} else {
				nowSpeak = 0
			}

			nowSpeakSpoken, present := numbersSpoken[nowSpeak]
			if present {
				numbersSpoken[nowSpeak] = NumberTurns{inc, nowSpeakSpoken.turnLastSpoken}
			} else {
				numbersSpoken[nowSpeak] = NumberTurns{inc, 0}
			}
			turns = append(turns, nowSpeak)
		}
	}

	return turns[stop]
}

func main() {
	filename := "day15_input.csv"
	inputs := readInput(filename)
	goUntilQ1 := 2020
	goUntilQ2 := 30000000

	// Given your starting numbers, what will be the 2020th number spoken?
	q1ans := q1(inputs, goUntilQ1)
	fmt.Println("part 1: ", q1ans)

	// Given your starting numbers, what will be the 30000000th number spoken? 
	q2 := q1(inputs, goUntilQ2)
	fmt.Println("part 2: ", q2)
}