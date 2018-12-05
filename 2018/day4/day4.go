// read it
// find Aa, aA, bB, Bb, cC, or Cc and remove
// keep removing until the difference between the length of the last round and the new round stops changing

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func q1(filename string) int {

	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)

	var records []string

	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}
	f.Close()

	sort.Strings(records)
	// for _, record := range records {
	// 	fmt.Println(record)
	// }

	// by guard, how many minutes does the guard spend asleep?
	// map where key = guard ID, value = total number of minutes asleep
	total_minutes_asleep := make(map[int]int)
	var guard int
	var start_sleep int
	var wake_up int

	for _, record := range records {
		if strings.Contains(record,"Guard") {
			hash_pos := strings.Index(record, "#")
			begins_pos := strings.Index(record, "b")
			guard_id := record[hash_pos+1:begins_pos-1]
			// fmt.Println(guard_id)
			guard, _ = strconv.Atoi(guard_id)
		}
		if strings.Contains(record,"falls asleep") {
			// fmt.Println(record)
			// fmt.Println("minute :",record[15:17])
			start_sleep, _ = strconv.Atoi(record[15:17])
		}
		if strings.Contains(record,"wakes up") {
			// fmt.Println(record)
			// fmt.Println("minute :",record[15:17])
			wake_up, _ = strconv.Atoi(record[15:17])
			time := wake_up - start_sleep
			// fmt.Println("started sleep at minute ",start_sleep," and woke up at minute ",wake_up)
			// fmt.Println("guard ",guard," slept for ",time," minutes")
			total_minutes_asleep[guard] = total_minutes_asleep[guard] + time
			// fmt.Println("guard ",guard," has slept for ",total_minutes_asleep[guard]," minutes in total")
		}
	}

	var longest_time_asleep int
	var sleepiest_guard int
	for sleeping_guard, time_asleep := range total_minutes_asleep {
		if time_asleep > longest_time_asleep {
			sleepiest_guard = sleeping_guard
			longest_time_asleep = time_asleep
		}
	}
	// fmt.Println(sleepiest_guard)

	// once we've found the guard, go back and grab all the times he was asleep
	asleep_by_minute := [60]int{}
	guard = 0

	for _, record := range records {
		if strings.Contains(record,"Guard") {
			hash_pos := strings.Index(record, "#")
			begins_pos := strings.Index(record, "b")
			guard_id := record[hash_pos+1:begins_pos-1]
			// fmt.Println(guard_id)
			guard, _ = strconv.Atoi(guard_id)
		}
		if guard == sleepiest_guard {
			if strings.Contains(record,"falls asleep") {
				// fmt.Println(record)
				// fmt.Println("minute :",record[15:17])
				start_sleep, _ = strconv.Atoi(record[15:17])
			}
			if strings.Contains(record,"wakes up") {
				// fmt.Println(record)
				// fmt.Println("minute :",record[15:17])
				wake_up, _ = strconv.Atoi(record[15:17])
				for i := start_sleep; i < wake_up; i++ {
					asleep_by_minute[i]++
				}
			}
		}
	}

	var sleepiest_minute int
	var most_times_asleep int
	for minute, times_asleep := range asleep_by_minute {
		if times_asleep > most_times_asleep {
			sleepiest_minute = minute
			most_times_asleep = times_asleep
		}
	}
	// fmt.Println(sleepiest_minute)

	return sleepiest_guard * sleepiest_minute
}

func q2(filename string) int {

	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var records []string

	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}

	sort.Strings(records)
	// for _, record := range records {
	// 	fmt.Println(record)
	// }

	// by guard, how many minutes does the guard spend asleep?
	// map where key = guard ID, value = another map where key = minute
	sleepy_guards := make(map[int]map[int]int)
	var guard int
	var start_sleep int
	var wake_up int

	for _, record := range records {
		if strings.Contains(record,"Guard") {
			hash_pos := strings.Index(record, "#")
			begins_pos := strings.Index(record, "b")
			guard_id := record[hash_pos+1:begins_pos-1]
			// fmt.Println(guard_id)
			guard, _ = strconv.Atoi(guard_id)
			if len(sleepy_guards[guard]) == 0 {
				sleepy_guards[guard] = make(map[int]int)
			}
		}
		if strings.Contains(record,"falls asleep") {
			// fmt.Println(record)
			// fmt.Println("minute :",record[15:17])
			start_sleep, _ = strconv.Atoi(record[15:17])
		}
		if strings.Contains(record,"wakes up") {
			// fmt.Println(record)
			// fmt.Println("minute :",record[15:17])
			wake_up, _ = strconv.Atoi(record[15:17])
			for i := start_sleep; i < wake_up; i++ {
				sleepy_guards[guard][i]++
			}
		}
	}

	var sleepiest_guard int
	var sleepiest_minute int
	var most_times_asleep int
	for sleepy_guard, arr := range sleepy_guards {
		for minute, times_asleep := range arr {
			if times_asleep > most_times_asleep {
				sleepiest_minute = minute
				sleepiest_guard = sleepy_guard
				most_times_asleep = times_asleep
			}
		}
	}
	// fmt.Println(sleepiest_guard)
	// fmt.Println(sleepiest_minute)
	// fmt.Println(most_times_asleep)

	return sleepiest_guard * sleepiest_minute}

func main() {
	fmt.Println("part1: ", q1("day4_input.txt"))
	fmt.Println("part2: ", q2("day4_input.txt"))
}
