package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func readInput(filename string) ([]string, []int, int) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Bus schedules are defined based on a timestamp that measures the
	// number of minutes since some fixed reference point in the past.
	// At timestamp 0, every bus simultaneously departed from the sea port.
	// Your notes (your puzzle input) consist of two lines. The first line
	// is your estimate of the earliest timestamp you could depart on a bus.
	// The second line lists the bus IDs that are in service according to the
	// shuttle company; entries that show x must be out of service, so you
	// decide to ignore them.
	timeToLeave, _ := strconv.Atoi(lines[0])
	buses := strings.Split(lines[1],",")
	var schedule []int
	for _, bus := range buses {
		if bus != "x" {
			time, _ := strconv.Atoi(bus)
			schedule = append(schedule, time)
		}
	}
	return buses, schedule, timeToLeave
}

func q1(schedule []int, startTime int) int {
	var min int
	var answer int
	for i, e := range schedule {
		earliestTime := int(math.Ceil(float64(startTime) / float64(e))) * e
		fmt.Println(min, i, e)
		if i == 0 || earliestTime - startTime < min {
			min = earliestTime - startTime
			answer = min * e
		}
	}
	return answer
}

func q2(buses []string) int {
	// However, with so many bus IDs in your list, surely the actual earliest
	// timestamp will be larger than 100000000000000!
	magicTime := 100000000000000

	// An x in the schedule means there are no constraints on what bus
	// IDs must depart at that time.

	return magicTime
}

func main() {
	filename := "day13_input.csv"
	buses, schedule, startTime := readInput(filename)

	// To save time once you arrive, your goal is to figure out
	// the earliest bus you can take to the airport.
	q1 := q1(schedule, startTime)

	// The shuttle company is running a contest: one gold coin for
	// anyone that can find the earliest timestamp such that the first
	// bus ID departs at that time and each subsequent listed bus ID
	// departs at that subsequent minute. (The first line in your input
	// is no longer relevant.)
	q2 := q2(buses)

	fmt.Println("part 1: ", q1)
	fmt.Println("part 2: ", q2)
}