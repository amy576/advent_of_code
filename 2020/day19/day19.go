package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

func readInput(filename string) (map[string][][]string, map[string]string, []string) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var messages []string
	rules := make(map[string][][]string)
	starterRules := make(map[string]string)

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
			// if strings.Contains(line, string('"')) {
			// 	starterRules[splits[0]] = strings.Trim(splits[1], string('"'))
			// } else {
			var multi [][]string
			for _, x := range strings.Split(splits[1], " | ") {
				multi = append(multi, strings.Split(x, " "))
			}
			rules[splits[0]] = multi
			// }
		}
	}
	return rules, starterRules, messages
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

// combination of day 18 equational replace + day 14 permutation generation

// start at rule 0 and start replacing
func solveForRules(rules map[string][][]string, end int) []string {
	endRule := rules[strconv.Itoa(end)] // this is a [][]string
	re := regexp.MustCompile("[0-9]+")

	// turn each number within endRule into a non-number
	containsNumber := re.FindAllString(turnIntoString(endRule), -1)
	for len(containsNumber) > 0 {
		for i, ruleSet := range endRule {		 // ruleSet is a []string
			for j, number := range ruleSet {
				fmt.Println("number", number)
				replace := rules[number]
				var new string
				for x, replaceRule := range replace {
					new += strings.Join(replaceRule, " ")
					if x < len(replace) -1 {
						new += " | "
					}
				}
				fmt.Println("new", new)
				endRule[i][j] = new
			}
			fmt.Println("ruleSet", endRule[i])
			containsNumber = re.FindAllString(turnIntoString(endRule), -1)
			fmt.Println(containsNumber)
			fmt.Println(endRule)
		}
	}

	var finalStrings []string
	for _, rules := range endRule { // rules will be a []string
		finalStrings = append(finalStrings, strings.Join(rules, ""))
	}
	return finalStrings
}

// function to create permutations appropriately
// assumes no more than one | in any given rule (at the start)

// now go through all the rules and replace, cleaning through until there are no
// more numbers, only a's and b's

// once that's done, consolidate rules[0] into []string (strings.join all the [][]string)
// iterate over messages, iterate over rules[0] and increment match number if they match

func main() {
	filename := "day19_test.csv"
	rules, starterRules, messages := readInput(filename)
	fmt.Println("rules: ", rules)
	fmt.Println("starting rules: ", starterRules)
	fmt.Println("messages: ", messages)

	fmt.Println(solveForRules(rules, 0))

	// This program starts by specifying a bitmask (mask = ....).
	// The program then attempts to write values to certain memory addresses.
	// What is the sum of all values left in memory after it completes?
	// q1 := q1(bits)
	// fmt.Println("part 1: ", q1)

	// q2 := q2(bits)
	// fmt.Println("part 2: ", q2)
}