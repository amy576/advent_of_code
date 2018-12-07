package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
)

func indexOf(element string, data []string) (int) {
   for k, v := range data {
       if element == v {
           return k
       }
   }
   return -1    //not found.
}

func q1(filename string) string {
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	var records []string
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}
	f.Close()

	// make a map of all the steps
	// Step Q must be finished before step O can begin.
	steps := make(map[string][]int)
	for _, record := range records {
		step_letter_first := strings.Split(record," ")[1]
		step_letter_second := strings.Split(record," ")[7]
		steps[step_letter_first] = make([]int,0)
		steps[step_letter_second] = make([]int,0)
	}

	// make a map of all steps with their prerequisite steps
	step_reqs := make(map[string][]string)
	for _, record := range records {
		step_letter_first := strings.Split(record," ")[1]
		step_letter_second := strings.Split(record," ")[7]
		step_reqs[step_letter_second] = append(step_reqs[step_letter_second],step_letter_first)
	}

	var answer string

	// first character is where it is in steps, but not in step_reqs
	// append first characters to "answer" in sorted order
	ready_steps := make([]string,0)
	for step_letter, _ := range steps {
		_, ok := step_reqs[step_letter]
		if ok {
			continue
		} else {
			ready_steps = append(ready_steps, step_letter)
			step_reqs[step_letter] = make([]string,0)
		}
	}
	sort.Strings(ready_steps)
	answer += ready_steps[0]

	// while the answer has not used all the letters yet
	for i:=0; i < len(steps) - 1; i++ {
		ready_steps = nil
		// iterate over all of the step IDs in step_reqs
		for step_letter, reqs := range step_reqs {
			// don't iterate back over letters that are already done
			if strings.Contains(answer,step_letter) {
				continue
			} else {
				// for each value in step_reqs[step ID], check if step ID is in "answer" (to see if is done)
				new_reqs := make([]string,0)
				for _, req := range reqs {
					// if it is done, remove that value from step_reqs[step ID]
					if strings.Contains(answer,req) {
						continue
					} else {
						new_reqs = append(new_reqs, req)
					}
				}
				step_reqs[step_letter] = new_reqs
				if len(step_reqs[step_letter]) == 0 {
					ready_steps = append(ready_steps, step_letter)
				}
			}
		}
		// sort the slice of "ready to do" step IDs, and append the first value to "answer"
		sort.Strings(ready_steps)
		answer += ready_steps[0]
	}
	return answer
}

func q2(filename string) int {

	return 0
}

func main() {
	fmt.Println("part1: ", q1("day7_input.txt"))
	fmt.Println("part2: ", q2("day6_input.txt"))
}
