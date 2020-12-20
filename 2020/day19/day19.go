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
			// rules[splits[0]] = strings.Trim(splits[1], string('"'))
			rules[splits[0]] = splits[1]
		}
	}
	return rules, messages
}

func regex(rules map[string]string, rule string) (re string) {
	// fmt.Println(rule)
	if rules[rule][0] == '"' {
		return rules[rule][1 : len(rules[rule])-1]
	}
	for _, s := range strings.Split(rules[rule], " | ") {
		re += "|"
		for _, s := range strings.Fields(s) {
			re += regex(rules, s)
		}
	}
	return "(?:" + re[1:] + ")"
}

func count(rules map[string]string, messages []string, ruleNum int) int {
	// newRules := cleanRules(rules)
	// re := regexp.MustCompile(newRules[strconv.Itoa(ruleNum)])
	re := regexp.MustCompile("(?m)^"+regex(rules, strconv.Itoa(ruleNum))+"$")
	// re := regexp.MustCompile(regex(rules, strconv.Itoa(ruleNum)))
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

	// clean := cleanRules(rules)
	// fmt.Println(clean)
	// fmt.Println("rules: ", rules)
	// fmt.Println("messages: ", messages)

	// Your goal is to determine the number of messages that completely match rule 0.
	q1ans := count(rules, messages, solveFor)
	fmt.Println("part 1: ", q1ans)

	// Completely replace rules 8: 42 and 11: 42 31 with the following:
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	// now, the rules do contain loops, and the list of messages they could
	// hypothetically match is infinite. You'll need to determine how these
	// changes affect which messages are valid. After updating rules 8 and 11,
	// how many messages completely match rule 0?

	rules["8"] = string('"') + regex(rules, "42") + "+" + string('"')

	var build string
	for i := 1; i <= 10; i++ {
		build = build + "|" + regex(rules, "42") + "{" + strconv.Itoa(i) + "}" + regex(rules, "31") + "{" + strconv.Itoa(i) + "}"
	}
	rules["11"] = string('"') + "(?:" + build[1:] + ")" + string('"')

	q2 := count(rules, messages, solveFor)
	fmt.Println("part 2: ", q2)
}