package main

import (
    "fmt"
    "os"
    "strconv"
	"bufio"
	"strings"
)

// Each line gives the password policy and then the password.
// The password policy indicates the lowest and highest number
// of times a given letter must appear for the password to be valid.
// For example, 1-3 a means that the password must contain a at least
// 1 time and at most 3 times.
func q1(filename string) int {
    f, _ := os.Open(filename)
    defer f.Close()
	
	valid := 0
	var dashPos int
	var colonPos int
	var minChar int
	var maxChar int
	var char string
	var seekChar string

	scanner := bufio.NewScanner(f)

	var records []string

	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}
	
	for _, record := range records {
		dashPos = strings.Index(record, "-")
		colonPos = strings.Index(record, ":")
		minChar, _ = strconv.Atoi(record[:dashPos])
		maxChar, _ = strconv.Atoi(record[dashPos+1:colonPos-2])
		char = record[colonPos-1:colonPos]
		seekChar = record[colonPos+2:]

		numFound := strings.Count(seekChar, char)

		if numFound >= minChar && numFound <= maxChar {
			valid++
		}
	}

	return valid
}

// Each policy actually describes two positions in the password,
// where 1 means the first character, 2 means the second character,
// and so on. (Be careful; Toboggan Corporate Policies have no concept
// of "index zero"!) Exactly one of these positions must
// contain the given letter
func q2(filename string) int {
    f, _ := os.Open(filename)
    defer f.Close()
	
	valid := 0
	var dashPos int
	var colonPos int
	var charPosOne int
	var charPosTwo int
	var char string
	var seekChar string

	scanner := bufio.NewScanner(f)

	var records []string

	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}
	
	for _, record := range records {
		dashPos = strings.Index(record, "-")
		colonPos = strings.Index(record, ":")
		charPosOne, _ = strconv.Atoi(record[:dashPos])
		charPosTwo, _ = strconv.Atoi(record[dashPos+1:colonPos-2])
		char = record[colonPos-1:colonPos]
		seekChar = record[colonPos+2:]

		if seekChar[charPosOne-1:charPosOne] == char && seekChar[charPosTwo-1:charPosTwo] != char {
			valid++
		}

		if seekChar[charPosOne-1:charPosOne] != char && seekChar[charPosTwo-1:charPosTwo] == char {
			valid++
		}
	}

	return valid
}

func main() {
    fmt.Println("part 1: ", q1("day2_input.csv"))
    fmt.Println("part 2: ", q2("day2_input.csv"))
}