package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

var dig = regexp.MustCompile("[0-9]+")

func readInput(filename string) (map[string]string, []string) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var messages []string
	rules := make(map[string]string)

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
			rules[splits[0]] = strings.Trim(splits[1], string('"'))
		}
	}
	return rules, messages
}

// start at rule 0 and start replacing
func rulesToRegex(rules map[string]string, end int) string {
	endRule := rules[strconv.Itoa(end)]
	containsNumber := dig.FindAllString(endRule, -1)
	// fmt.Println("start with", endRule)
	for len(containsNumber) > 0 {
		for i, elem := range strings.Split(endRule, " ") {
			element := strings.Trim(elem, "()")
			if len(dig.FindAllString(elem, -1)) > 0 {
				// fmt.Println(element, ":", rules[element])
				if strings.Contains(rules[element], "|") {
					if i == 0 {
						endRule = strings.Replace(endRule, element + " ", "(" + rules[element] + ") ", 1)
					} else if i == len(strings.Split(endRule, " ")) - 1 {
						endRule = strings.Replace(endRule, " " + element, " (" + rules[element] + ")", 1)
					} else {
						endRule = strings.Replace(endRule, " " + element + " ", " (" + rules[element] + ") ", 1)
						endRule = strings.Replace(endRule, "(" + element + " ", "((" + rules[element] + ") ", 1)
						endRule = strings.Replace(endRule, " " + element + ")", " (" + rules[element] + "))", 1)
					}
				} else {
					if i == 0 {
						endRule = strings.Replace(endRule, element + " ", rules[element] + " ", 1)
					} else if i == len(strings.Split(endRule, " ")) - 1 {
						endRule = strings.Replace(endRule, " " + element, " " + rules[element], 1)
					} else {
						endRule = strings.Replace(endRule, " " + element + " ", " " + rules[element] + " ", 1)
						endRule = strings.Replace(endRule, "(" + element + " ", "(" + rules[element] + " ", 1)
						endRule = strings.Replace(endRule, " " + element + ")", " " + rules[element] + ")", 1)
					}
				}
			}
			// fmt.Println("replace element and now we have", endRule)
		}
		containsNumber = dig.FindAllString(endRule, -1)
		// fmt.Println(len(endRule))
		// fmt.Println(endRule)
		// fmt.Println(regexp.MustCompile(strings.ReplaceAll(endRule, " ", "")))
	}

	// fmt.Println(endRule)

	cleanedRule := strings.ReplaceAll(endRule, " ", "")
	// fmt.Println(cleanedRule)
	return cleanedRule
}

func q1(rules map[string]string, messages []string, ruleNum int) int {
	re := regexp.MustCompile(rulesToRegex(rules, ruleNum))
	// iterate over messages, iterate over rules[0] and increment match number if they match
	var valid int
	for _, message := range messages {
		found := re.FindAllString(message, -1)
		if len(found) == 1 {
			if len(found[0]) == len(message) {
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