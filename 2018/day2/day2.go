package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func q1(filename string) int {
	f, _ := os.Open(filename)
	defer f.Close()

	r := csv.NewReader(f)
	boxesWithTwo := 0
	boxesWithThree := 0

	records, _ := r.ReadAll()
	for _, record := range records {
		for _, field := range record {
			seen := make(map[rune]int)
			for _, c := range field {
				seen[c] = strings.Count(field, string(c))
			}
		TwoLoop:
			for r := range seen {
				if seen[r] == 2 {
					boxesWithTwo++
					break TwoLoop
				}
			}
		ThreeLoop:
			for r := range seen {
				if seen[r] == 3 {
					boxesWithThree++
					break ThreeLoop
				}
			}
		}
	}

	checksum := boxesWithTwo * boxesWithThree
	return checksum
}

func q2(filename string) string {
	f, _ := os.Open(filename)
	defer f.Close()

	r := csv.NewReader(f)

	var id1 string
	// var id2 string
	diffPos := 0

	records, _ := r.ReadAll()
Loop:
	for i, record := range records {
		for j := i + 1; j < len(records); j++ {
			diff := 0
			for k := 0; k < len(record[0]); k++ {
				if record[0][k] != records[j][0][k] {
					diff++
					diffPos = k
				}
			}
			if diff == 1 {
				id1 = record[0]
				// id2 = records[j][0]
				break Loop
			}
		}
	}

	sameStr := id1[:diffPos] + id1[diffPos+1:]

	// fmt.Println(id1)
	// fmt.Println(id2)
	// fmt.Println(diff_pos)
	return sameStr

}

func main() {
	fmt.Println("part1: ", q1("day2_input.csv"))
	fmt.Println("part2: ", q2("day2_input.csv"))
}
