package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func q1(filename string) int {
	f, _ := os.Open(filename)
	defer f.Close()

	r := csv.NewReader(f)
	val := 0

	records, _ := r.ReadAll()
	for _, record := range records {
		for _, field := range record {
			i, _ := strconv.Atoi(field)
			val = val + i
		}
	}

	return val
}

func q2(filename string) int {
	val := 0
	frequencies := make(map[int]int)
	frequencies[val] = 1

	for frequencies[val] < 2 {
		f, _ := os.Open(filename)
		defer f.Close()

		r := csv.NewReader(f)

		records, _ := r.ReadAll()
		for _, record := range records {
			for _, field := range record {
				i, _ := strconv.Atoi(field)
				val = val + i
				frequencies[val]++
				if frequencies[val] == 2 {
					return val
				}
			}
		}
	}
	return val
}

func main() {
	fmt.Println("part1: ", q1("day1_input.csv"))
	fmt.Println("part2: ", q2("day1_input.csv"))
}
