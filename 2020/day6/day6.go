package main

import (
    "fmt"
    "os"
	"bufio"
	"strings"
)

// Each group's answers are separated by a blank line,
// and within each group, each person's answers are on a single line.
// Duplicate answers to the same question don't count extra;
// each question counts at most once.
// For each group, count the number of questions to which
// anyone answered "yes". What is the sum of those counts?
func q1(filename string) int {
    f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	
	yes := 0
	var groups []string
	var groupAnswer string
	newLines := 0

	var records []string
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}

	for i, record := range records {
		if record == "" {
			groups = append(groups, groupAnswer)
			newLines++
			groupAnswer = ""
		} else {
			// assume the first person's answer contains unique characters
			for _, answer := range record {
				if strings.ContainsAny(groupAnswer, string(answer)) == false {
					groupAnswer += string(answer)
				}
			}
		}

		// without this, we keep missing the last record
		if i == len(records) - 1 {
			groups = append(groups, groupAnswer)
		}
	}

	for _, group := range groups {
		// fmt.Println(len(group))
		yes += len(group)
	}

	fmt.Println(newLines)
	// there should always be 1 more group than newlines
	fmt.Println(len(groups))

	return yes
}

// For each group, count the number of questions to which
// everyone answered "yes". What is the sum of those counts?
func q2(filename string) int {
    f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	
	yes := 0
	var groups [][]string
	var groupAnswer []string
	newLines := 0

	var records []string
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}

	for i, record := range records {
		// make an array of arrays: array of groups, where each
		// group is an array of strings, each string is one person's answer
		if record == "" {
			groups = append(groups, groupAnswer)
			newLines++
			groupAnswer = []string{} // reset to empty array
		} else {
			groupAnswer = append(groupAnswer, record)
		}

		// without this, we keep missing the last record
		if i == len(records) - 1 {
			groups = append(groups, groupAnswer)
		}
	}
	
	// now iterate over the array of groups
	// check each character in the first person's answer and see if
	// subsequent answers contain it; if so, yes ++
	// if there is only 1 answer in the group, then yes + that len

	for _, group := range groups {
		firstAnswer := group[0]
		if len(group) == 1 {
			yes += len(firstAnswer)
		} else {
			var allAnswers string
			for _, indivAnswer := range group {
				allAnswers += indivAnswer
			}

			answers := make(map[string]int)
			for _, char := range firstAnswer {
				answers[string(char)] = strings.Count(allAnswers, string(char))
			}
			for _, element := range answers {
				if element == len(group) {
					yes++
				}
			}
		}
	}

	fmt.Println(newLines)
	// there should always be 1 more group than newlines
	fmt.Println(len(groups))

	return yes
}

func main() {
    fmt.Println("part 1: ", q1("day6_input.csv"))
    fmt.Println("part 2: ", q2("day6_input.csv"))
}