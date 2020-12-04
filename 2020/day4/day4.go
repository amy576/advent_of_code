package main

import (
    "fmt"
    "os"
	"bufio"
	"strings"
)

// Count the number of valid passports - those that have all required fields.
// Treat cid as optional. In your batch file, how many passports are valid?
// Required fields:
// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)
func q1(filename string) int {
    f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	
	valid := 0
	var passports []string
	var passport string
	newLines := 0

	var records []string
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}

	// we only care about number of valid and do not care how many are invalid
	for i, record := range records {
		if record == "" {
			passports = append(passports, passport)
			// fmt.Println(passport)
			newLines++
			passport = ""
		} else {
			passport += " "
			passport += record
		}

		
		// without this, we keep missing the last record
		if i == len(records) - 1 {
			passports = append(passports, passport)
		}
	}

	for _, check := range passports {
		if strings.Contains(check, " byr:") && strings.Contains(check, " iyr:") && strings.Contains(check, " eyr:") && strings.Contains(check, " hgt:") && strings.Contains(check, " hcl:") && strings.Contains(check, " ecl:") && strings.Contains(check, " pid:") {
			valid++
		}
	}

	fmt.Println(newLines)
	fmt.Println(len(passports))

	return valid
}


func main() {
    fmt.Println("part 1: ", q1("day4_input.csv"))
    // fmt.Println("part 2: ", q2("day4_input.csv"))
}