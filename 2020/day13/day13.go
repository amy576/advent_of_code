package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func readInput(filename string) ([][]int, int) {
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
	var schedule [][]int
	for i, bus := range buses {
		if bus != "x" {
			time, _ := strconv.Atoi(bus)
			var sched []int
			sched = append(sched, i)
			sched = append(sched, time)
			schedule = append(schedule, sched)
		}
	}
	return schedule, timeToLeave
}

func q1(schedule [][]int, startTime int) int {
	var min int
	var answer int
	for i, e := range schedule {
		bus := e[1]
		earliestTime := int(math.Ceil(float64(startTime) / float64(bus))) * bus
		// fmt.Println(min, bus)
		if i == 0 || earliestTime - startTime < min {
			min = earliestTime - startTime
			answer = min * bus
		}
	}
	return answer
}

func q2(buses [][]int) int {
	// In the example, the first time Bus 7 (t offset = 0) and Bus 13 align (t offset = 1),
	// is t = 77. With the current increment being 7, the new increment is 7 * 13 = 91,
	// meaning the current t of 77 + 91 is the next time the pattern will repeat. You
	// keep incrementing with 91, until Bus 59 can be found (at t offset = 4, since
	// we're skipping minutes 2 and 3). Rinse and repeat until you reach the end of
	// your line.
	magicTime := 1
	stillLooking := true
	for stillLooking {
		timeInc := 1
		valid := true
		for i := 0; i < len(buses); i++ {
			// fmt.Println(i, timeInc, magicTime)
			// keep checking the same first i bus(ses), incrementing time, until we find
			// a number that works for it/them
			if (magicTime + buses[i][0]) % buses[i][1] != 0 {
				valid = false
				break
			}
			// only do this if we found a number that works for the first i bus(ses)
			// multiplied instead of LCD because all our inputs are prime; this is how we
			// know to skip to the next one that definitely works for the first i bus(ses)
			// and can then check for the next bus
			timeInc *= buses[i][1]
		}

		// we got through the whole loop (e.g. all busses) without valid ever flipping
		// to false
		if valid {
			stillLooking = false
		} else {
			magicTime += timeInc
		}
	}

	return magicTime
}

func main() {
	filename := "day13_input.csv"
	schedule, startTime := readInput(filename)

	// To save time once you arrive, your goal is to figure out
	// the earliest bus you can take to the airport.
	q1 := q1(schedule, startTime)
	fmt.Println("part 1: ", q1)

	// The shuttle company is running a contest: one gold coin for
	// anyone that can find the earliest timestamp such that the first
	// bus ID departs at that time and each subsequent listed bus ID
	// departs at that subsequent minute. (The first line in your input
	// is no longer relevant.)
	q2 := q2(schedule)
	fmt.Println("part 2: ", q2)
}