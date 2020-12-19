package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

func readInput(filename string) (map[string][][]string, []string) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var messages []string
	rules := make(map[string][][]string)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	for _, line := range lines {
		if line[:1] == "a" || line[:1] == "b" {
			messages = append(messages, line)
		} else if strings.Contains(line, ":") {
			splits := strings.Split(line, ": ")
			var multi [][]string
			for _, x := range strings.Split(splits[1], " | ") {
				var single []string
				for _, y := range strings.Split(x, " ") {
					single = append(single, strings.Trim(y, string('"')))
				}
				multi = append(multi, single)
			}
			rules[splits[0]] = multi
		}
	}
	return rules, messages
}

func turnIntoString(rule [][]string) string {
	var ruleString string
	for i, ruleSet := range rule {
		ruleString += "("
		ruleString += strings.Join(ruleSet, " ")
		ruleString += ")"
		if i < len(rule) - 1 {
			ruleString += " | "
		}
	}

	return ruleString
}

// start at rule 0 and start replacing
func solveForRules(rules map[string][][]string, end int) []string {
	endRule := rules[strconv.Itoa(end)] // this is a [][]string
	re := regexp.MustCompile("[0-9]+")

	// turn each number within endRule into a non-number
	containsNumber := re.FindAllString(turnIntoString(endRule), -1)
	for len(containsNumber) > 0 {
		for i, ruleSet := range endRule {		 // ruleSet is a []string
			for j, number := range ruleSet {
				initialRuleSet := make([]string, len(ruleSet))
				copy(initialRuleSet, ruleSet)
				if number == "a" || number == "b" {
				} else {
					replace := rules[number] // replace is a [][]string

					// replace the current ruleSet using replace[0]
					overwriteRule := make([]string, len(ruleSet[:j]))
					copy(overwriteRule, ruleSet[:j])
					// endRule[i] = ruleSet[:j]
					for _, x := range replace[0] {
						overwriteRule = append(overwriteRule, x)
					}
					for _, y := range initialRuleSet[j+1:] {
						overwriteRule = append(overwriteRule, y)
					}
					endRule[i] = overwriteRule

					// if len(replace) > 1 then we need to append a new rule to endRule
					if len(replace) > 1 {
						newRule := make([]string, len(ruleSet[:j]))
						copy(newRule, ruleSet[:j])
						for _, x := range replace[1] {
							newRule = append(newRule, x)
						}
						for _, y := range initialRuleSet[j+1:] {
							newRule = append(newRule, y)
						}
						endRule = append(endRule, newRule)
					}

					break
				}
			}
			containsNumber = re.FindAllString(turnIntoString(endRule), -1)
			// fmt.Println("current endRule array", endRule)
		}
	}

	var finalStrings []string
	for _, rules := range endRule { // rules will be a []string
		finalStrings = append(finalStrings, strings.Join(rules, ""))
	}
	return finalStrings
}

func q1(rules map[string][][]string, messages []string, ruleNum int) int {
	validStrings := solveForRules(rules, ruleNum)
	// iterate over messages, iterate over rules[0] and increment match number if they match
	var valid int
	for _, message := range messages {
		for _, validString := range validStrings {
			if message == validString {
				valid++
			}
		}
	}
	return valid
}

func main() {
	filename := "day19_input.csv"
	rules, messages := readInput(filename)
	solveFor := 0
	// fmt.Println("rules: ", rules)
	// fmt.Println("messages: ", messages)

	// Your goal is to determine the number of messages that completely match rule 0.
	q1 := q1(rules, messages, solveFor)
	fmt.Println("part 1: ", q1)

	// q2 := q2(bits)
	// fmt.Println("part 2: ", q2)
}