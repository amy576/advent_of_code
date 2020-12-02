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

    var first_number int
    var second_number int
    
    for x, number_one := range records {
        for y, number_two := range records {
            first_number, _ = strconv.Atoi(number_one)
            second_number, _ = strconv.Atoi(number_two)
            if x != y && first_number + second_number == 2020 {
                return first_number * second_number
            }
        }
    }

    return first_number * second_number
}

// find the THREE entries that sum to 2020
// and then multiply those three numbers together
func q2(filename string) int {
    f, _ := os.Open(filename)
    defer f.Close()
	scanner := bufio.NewScanner(f)

	var records []string

	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
    }

    var first_number int
    var second_number int
    var third_number int
    
    for x, number_one := range records {
        for y, number_two := range records {
            for z, number_three := range records {
                first_number, _ = strconv.Atoi(number_one)
                second_number, _ = strconv.Atoi(number_two)
                third_number, _ = strconv.Atoi(number_three)
                if x != y && x != z && y != z && first_number + second_number + third_number == 2020 {
                    return first_number * second_number * third_number
                }
            }
        }
    }

    return first_number * second_number * third_number
}

func main() {
    fmt.Println("part 1: ", q1("day1_input.csv"))
    fmt.Println("part 2: ", q2("day1_input.csv"))
}