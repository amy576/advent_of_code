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

// start at rule 0 and start replacing
func cleanRules(rules map[string]string) map[string]string {
	cleanedRules := make(map[string]string)
	for num, endRule := range rules {
		fmt.Println("cleaning", num)
		containsNumber := dig.FindAllString(endRule, -1)
		// fmt.Println("start with", endRule)
		for len(containsNumber) > 0 {
			for i, elem := range strings.Split(endRule, " ") {
				element := strings.Trim(elem, "()")
				var replace string
				newReplace, present := cleanedRules[element]
				if present {
					replace = newReplace
				} else {
					replace = rules[element]
				}
				if len(dig.FindAllString(elem, -1)) > 0 {
					// fmt.Println(element, ":", rules[element])
					if strings.Contains(replace, "|") {
						if i == 0 {
							endRule = strings.Replace(endRule, element + " ", "(" + replace + ") ", -1)
						} else if i == len(strings.Split(endRule, " ")) - 1 {
							endRule = strings.Replace(endRule, " " + element, " (" + replace + ")", -1)
						} else {
							endRule = strings.Replace(endRule, " " + element + " ", " (" + replace + ") ", -1)
							endRule = strings.Replace(endRule, "(" + element + " ", "((" + replace + ") ", -1)
							endRule = strings.Replace(endRule, " " + element + ")", " (" + replace + "))", -1)
						}
					} else {
						if i == 0 {
							endRule = strings.Replace(endRule, element + " ", replace + " ", -1)
						} else if i == len(strings.Split(endRule, " ")) - 1 {
							endRule = strings.Replace(endRule, " " + element, " " + replace, -1)
						} else {
							endRule = strings.Replace(endRule, " " + element + " ", " " + replace + " ", -1)
							endRule = strings.Replace(endRule, "(" + element + " ", "(" + replace + " ", -1)
							endRule = strings.Replace(endRule, " " + element + ")", " " + replace + ")", -1)
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
		fmt.Println(cleanedRule)
		cleanedRules[num] = cleanedRule
	}

	return cleanedRules
}

func q1(rules map[string]string, messages []string, ruleNum int) int {
	// newRules := cleanRules(rules)
	// re := regexp.MustCompile(newRules[strconv.Itoa(ruleNum)])
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

	clean := cleanRules(rules)
	fmt.Println(clean)
	// fmt.Println("rules: ", rules)
	// fmt.Println("messages: ", messages)

	// Your goal is to determine the number of messages that completely match rule 0.
	q1 := q1(rules, messages, solveFor)
	fmt.Println("part 1: ", q1)

	// Completely replace rules 8: 42 and 11: 42 31 with the following:
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	// rules["8"] = "42 | 42 8"
	// rules["11"] = "42 31 | 42 11 31"
	// now, the rules do contain loops, and the list of messages they could
	// hypothetically match is infinite. You'll need to determine how these
	// changes affect which messages are valid. After updating rules 8 and 11,
	// how many messages completely match rule 0?
	// fmt.Println("part 2: ", q2)
}