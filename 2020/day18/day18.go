package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) []string {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func q1(equations []string) int {
	var answers int
	for _, equation := range equations {
		val := calculate(calcParens(equation, 1))
		answers += val 
	}
	return answers
}

func q2(equations []string) int {
	var answers int
	// f, _ := os.Create("day18_output.txt")
    // w := bufio.NewWriter(f)
	for _, equation := range equations {
		val := calculateQ2(calcParens(equation, 2))
		answers += val
		// w.WriteString(equation + "=" + strconv.Itoa(val) + "\n")
	}
    // w.Flush()
	return answers
}

// some function to handle parentheses and return the value within them
func calcParens(equation string, question int) string {
	for strings.Count(equation, "(") > 0 {
		// fmt.Println("calcParens:", equation)
		openBracket := strings.Index(equation, "(")
		closeBracket := strings.Index(equation, ")")
		nextOpenBracket := strings.Index(equation[openBracket+1:], "(") + openBracket + 1
		if nextOpenBracket > 0 && nextOpenBracket < closeBracket {
			// there's something nested in here
			openBracket = nextOpenBracket
		}
		var parenVal int
		if question == 1 {
			parenVal = calculate(equation[openBracket + 1 : closeBracket])
		} else if question == 2 {
			parenVal = calculateQ2(equation[openBracket + 1 : closeBracket])
		}
		equation = strings.Replace(equation, equation[openBracket : closeBracket+1], strconv.Itoa(parenVal), -1)
	}
	// fmt.Println("calcParens output:", equation)
	return equation
}

// function to go from left to right of the given sequence
func calculate(input string) int {
	// Now, addition and multiplication have different precedence levels,
	// but they're not the ones you're familiar with. Instead, addition is
	// evaluated before multiplication.
	var total int
	splits := strings.Split(input, " ")
	operator := "add"
	for _, char := range splits {
		if char == "+" {
			operator = "add"
		} else if char == "*" {
			operator = "multiply"
		} else {
			val, _ := strconv.Atoi(char)
			if operator == "add" {
				total += val
			} else if operator == "multiply" {
				total *= val
			}
		}
	}
	return total
}

func calculateQ2(input string) int {
	for strings.Count(input, "+") > 0 {
		// fmt.Println("calculateQ2", input)
		plusPos := strings.Index(input, "+")
		var startFrom int
		prevMultPos := strings.LastIndex(input[:plusPos], "*")
		if prevMultPos > 0 {
			startFrom = prevMultPos + 2
		} else {
			startFrom = 0
		}
		nextPlusPos := strings.Index(input[plusPos+1:], "+") + plusPos + 1
		nextMultPos := strings.Index(input[plusPos+1:], "*") + plusPos + 1
		goUntil := len(input)
		if nextPlusPos == plusPos || nextPlusPos > nextMultPos {
			if nextMultPos > plusPos {
				goUntil = nextMultPos - 1	
			}
		} else if nextPlusPos > plusPos {
			goUntil = nextPlusPos - 1
		}
		val := calculate(input[startFrom : goUntil])
		// fmt.Println(input[startFrom : goUntil], " = ", val)
		input = strings.Replace(input, input[startFrom : goUntil], strconv.Itoa(val), 1)
	}
	// fmt.Println("calculateQ2 output", input)
	total := calculate(input)
	return total
}

func main() {
	filename := "day18_input.csv"
	equations := readInput(filename)
	// Key assumptions:
	// - Parentheses are never nested more than 2 deep

	// Evaluate the expression on each line of the homework;
	// what is the sum of the resulting values?
	q1 := q1(equations)
	fmt.Println("part 1: ", q1)
	q2 := q2(equations)
	fmt.Println("part 2: ", q2)
}