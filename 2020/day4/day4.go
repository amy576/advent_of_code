package main

import (
    "fmt"
    "os"
	"bufio"
	"strings"
	"regexp"
	"strconv"
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
	// there should always be 1 more passport record than newlines
	fmt.Println(len(passports))

	return valid
}

// You can continue to ignore the cid field, but each other
// field has strict rules about what values are valid for automatic validation:
// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
func q2(filename string) int {
    f, _ := os.Open(filename)
    defer f.Close()
	scanner := bufio.NewScanner(f)

	valid := 0
	var passports []string
	var passport string
	newLines := 0

	validEclCol := "amb blu brn gry grn hzl oth"
	validEcl := strings.Split(validEclCol," ")

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
		// make each passport into a map
		byrValid := false
		iyrValid := false
		eyrValid := false
		hgtValid := false
		hclValid := false
		eclValid := false
		pidValid := false
		fields := make(map[string]string)
		check = check[1:]
		s := strings.Split(check, " ")
		for _, field := range s {
			unmapped := strings.Split(field, ":")
			keyValue := unmapped[0]
			elementValue := unmapped[1]
			fields[keyValue] = elementValue
			// fmt.Println("key is ",keyValue," for value ",elementValue)
		}

		byr, foundByr := fields["byr"]
		// fmt.Println(byr, foundByr)
		iyr, foundIyr := fields["iyr"]
		// fmt.Println(iyr, foundIyr)
		eyr, foundEyr := fields["eyr"]
		// fmt.Println(eyr, foundEyr)
		hgt, foundHgt := fields["hgt"]
		// fmt.Println(hgt, foundHgt)
		hcl, foundHcl := fields["hcl"]
		// fmt.Println(hcl, foundHcl)
		ecl, foundEcl := fields["ecl"]
		// fmt.Println(ecl, foundEcl)
		pid, foundPid := fields["pid"]
		// fmt.Println(pid, foundPid)

		// check if all the keys are even there first
		if foundByr && foundIyr && foundEyr && foundHgt && foundHcl && foundEcl && foundPid {
			byrValue, _ := strconv.Atoi(byr)
			// fmt.Println(byrValue)
			if len(byr) == 4 && byrValue >= 1920 && byrValue <= 2002 {
				byrValid = true
				// fmt.Println(byrValid)
			}

			iyrValue, _ := strconv.Atoi(iyr)
			// fmt.Println(iyrValue)
			if len(iyr) == 4 && iyrValue >= 2010 && iyrValue <= 2020 {
				iyrValid = true
				// fmt.Println(iyrValid)
			}

			eyrValue, _ := strconv.Atoi(eyr)
			// fmt.Println(eyrValue)
			if len(eyr) == 4 && eyrValue >= 2020 && eyrValue <= 2030 {
				eyrValid = true
				// fmt.Println(eyrValid)
			}

			hgtUnit := hgt[len(hgt)-2:]
			hgtAmount, _ := strconv.Atoi(hgt[:len(hgt)-2])
			// fmt.Println(hgtAmount, hgtUnit)
			if hgtUnit == "cm" && hgtAmount >= 150 && hgtAmount <= 193 {
				hgtValid = true
			}
			if hgtUnit == "in" && hgtAmount >= 59 && hgtAmount <= 76 {
				hgtValid = true
			}
			// fmt.Println(hgtValid)

			hclMatched, _ := regexp.MatchString(`^#(?:[0-9a-fA-F]{3}){1,2}$`, hcl)
			if hclMatched {
				hclValid = true
			}

			for _, validColor := range validEcl {
				if ecl == validColor {
					eclValid = true
				}
			}

			_, piderr := strconv.Atoi(pid)
			if piderr == nil && len(pid) == 9 {
				pidValid = true
			}


			if byrValid && iyrValid && eyrValid && hgtValid && hclValid && eclValid && pidValid {
				valid++
			}

		}
	}

	fmt.Println(newLines)
	fmt.Println(len(passports))

	return valid
}

func main() {
    fmt.Println("part 1: ", q1("day4_input.csv"))
    fmt.Println("part 2: ", q2("day4_input.csv"))
}