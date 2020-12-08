package main

import (
    "fmt"
    "os"
	"bufio"
	"strings"
	"strconv"
)

func readFiles(filename string) []string {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	
	var records []string
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}

	return records
}

func main() {
    fmt.Println("part 1: ", q1("day8_test.csv"))
    // fmt.Println("part 2: ", q2("day8_test.csv"))
}