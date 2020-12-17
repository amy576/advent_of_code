package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) (map[string][][]int, [][]int, []int) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	validValues := make(map[string][][]int)
	var yourTicket []int
	var nearbyTickets [][]int
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	var validValuesListStop int
	var nearbyTicketsStart int
	for i, x := range lines {
		if x == "your ticket:" {
			validValuesListStop = i - 1
			yourTicketTemp := strings.Split(lines[i+1], ",")
			for _, t := range yourTicketTemp {
				ticketVal, _ := strconv.Atoi(t)
				yourTicket = append(yourTicket, ticketVal)
			}
		} else if x == "nearby tickets:" {
			nearbyTicketsStart = i + 1
		}
	}

	// The rules for ticket fields specify a list of fields that exist somewhere
	// on the ticket and the valid ranges of values for each field.
	for j := 0; j < validValuesListStop; j++ {
		valid := strings.Split(lines[j], ": ")[1]
		field := strings.Split(lines[j], ": ")[0]
		valSets := strings.Split(valid, " or ")
		var fieldVals [][]int
		for _, vals := range valSets {
			valList := strings.Split(vals,"-")
			var val []int
			for _, z := range valList {
				v, _ := strconv.Atoi(z)
				val = append(val, v)
			}
			fieldVals = append(fieldVals, val)
		}
		validValues[field] = fieldVals
	}
	for k := nearbyTicketsStart; k < len(lines); k++ {
		var ticket []int
		ticketVals := strings.Split(lines[k], ",")
		for _, y := range ticketVals {
			ticketVal, _ := strconv.Atoi(y)
			ticket = append(ticket, ticketVal)
		}
		nearbyTickets = append(nearbyTickets, ticket)
	}
	return validValues, nearbyTickets, yourTicket
}

func q1(validValues map[string][][]int, checkTickets [][]int) (int, [][]int) {
	var scanningErrRate int
	var validTickets [][]int

	// It doesn't matter which position corresponds to which field;
	// you can identify invalid nearby tickets by considering only
	// whether tickets contain values that are not valid for any field.

	// check each ticket
	for _, ticket := range checkTickets {
		ticketValid := true
		// check each value in the ticket
		for _, ticketVal := range ticket {
			valid := false
			CheckAgainstValidValues:
				for _, element := range validValues {
					for _, valueSet := range element {
						if ticketVal >= valueSet[0] && ticketVal <= valueSet[1] {
							valid = true
							break CheckAgainstValidValues
						}
					}
				}
			if valid == false {
				scanningErrRate += ticketVal
				ticketValid = false
			}
		}
		if ticketValid {
			validTickets = append(validTickets, ticket)
		}
	}

	return scanningErrRate, validTickets
}

func q2(validValues map[string][][]int, validTickets [][]int, yourTicket []int) int {
	// Use the remaining valid tickets to determine which field is which.
	// Using the valid ranges for each field, determine what order the fields
	// appear on the tickets. The order is consistent between all tickets:
	// if seat is the third field, it is the third field on every ticket,
	// including your ticket.

	fieldPos := make(map[string]int)
	fieldPosPossible := make(map[string][]int)
	validTicketCount := len(validTickets)

	// check the first position value of each ticket, and iterate from there
	for i := 0; i < len(validValues); i++ {
		var possibleFields []string
		for _, ticket := range validTickets {
			checkVal := ticket[i]
			for key, element := range validValues {
				for _, valueSet := range element {
					if checkVal >= valueSet[0] && checkVal <= valueSet[1] {
						possibleFields = append(possibleFields, key)
					}
				}
			}
		}
		// fmt.Println(i, possibleFields)
		allPossibleFields := strings.Join(possibleFields, ",")
		for field, _ := range validValues {
			if strings.Count(allPossibleFields, field) == validTicketCount {
				fieldPosPossible[field] = append(fieldPosPossible[field], i)
			}
		}
	}

	var usedUpPos []int

	for j := 0; j < len(validValues); j++ {
		for key, element := range fieldPosPossible {
			if len(element) == j + 1 {
				if len(usedUpPos) == 0 && j == 0 {
					fieldPos[key] = element[0]
					usedUpPos = append(usedUpPos, element[0])
				} else {
					for _, maybePos := range element {
						alreadyUsed := false
						for _, usedUp := range usedUpPos {
							if usedUp == maybePos {
								alreadyUsed = true
							}
						}
						if alreadyUsed == false {
							fieldPos[key] = maybePos
							usedUpPos = append(usedUpPos, maybePos)
						}
					}
				}
			}
		}
	}

	fmt.Println(fieldPos)
	departure := 1
	for field, pos := range fieldPos {
		// fmt.Println(field, pos)
		if strings.Contains(field, "departure") {
			departure *= yourTicket[pos]
		}
	}

	return departure
}

func main() {
	filename := "day16_input.csv"
	fieldValues, nearbyTickets, myTicket := readInput(filename)

	// Adding together all of the invalid values produces your ticket scanning
	// error rate. Consider the validity of the nearby tickets you scanned.
	// What is your ticket scanning error rate
	q1, validTickets := q1(fieldValues, nearbyTickets)
	fmt.Println("part 1: ", q1)

	// Once you work out which field is which, look for the six fields on
	// your ticket that start with the word departure. What do you get if
	// you multiply those six values together? 
	q2 := q2(fieldValues, validTickets, myTicket)
	fmt.Println("part 2: ", q2)
}