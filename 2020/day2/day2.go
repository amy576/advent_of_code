package main

import (
    "fmt"
    "os"
    "strconv"
    "bufio"
)

// find the two entries that sum to 2020
// and then multiply those two numbers together
func q1(filename string) int {
    f, _ := os.Open(filename)
    defer f.Close()
	scanner := bufio.NewScanner(f)

	var records []string

	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
    }
}

func main() {
    fmt.Println("part 1: ", q1("day2_input.csv"))
    // fmt.Println("part 2: ", q2("day2_input.csv"))
}