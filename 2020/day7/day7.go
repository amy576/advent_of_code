package main

import (
    "fmt"
    "os"
	"bufio"
	"strings"
	"strconv"
)

// how many colors can, eventually, contain
// at least one shiny gold bag?
func q1(filename string) int {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	
	var records []string
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}

	// var canContain []string
	canContain := make(map[string]int)

	// make a map of arrays: bright_white bag (key) contains
	// [shiny_gold] bag (element)
	bagRules := make(map[string][]string)
	for _, rule := range records {
		bagName := rule[:strings.Index(rule, " contain")-5]
		bagsInBag := strings.Split(rule[strings.Index(rule, "contain")+8:],", ")
		// fmt.Println(bagName)
		if bagsInBag[0] != "no other bags." {
			for i, bag := range bagsInBag {
				bagsInBag[i] = bag[strings.Index(bag, " ")+1:strings.Index(bag, " bag")]
				// fmt.Println(bagsInBag[i])
			}
			bagRules[bagName] = bagsInBag
			// fmt.Println(bagName, ": ", bagRules[bagName])
		}
	}
	
	// iterate over map and make a list of keys where the element
	// array contains shiny_gold (this is the list of bags that
	// can directly contain the shiny gold bag)
	for key, element := range bagRules {
		for _, bag := range element {
			if bag == "shiny gold" {
				// canContain = append(canContain, key)
				canContain[key] = 1
			}
		}
	}
	// fmt.Println(canContain)

	// then iterate over the list of bags that can contain and see
	// which bags contain THOSE bags
	// also keep a counter of how many bags got appended
	// if that counter is 0, stop looping
	appended := 1
	for appended > 0 {
		appended = 0
		for bag, _ := range canContain {
			for key, element := range bagRules {
				for _, searchBag := range element {
					if searchBag == bag && canContain[key] != 1 {
						canContain[key] = 1
						appended++
					}
				}
			}
		}
	}

	numContain := len(canContain)
	return numContain

}

// How many individual bags are required inside your single shiny gold bag?
func q2(filename string) int {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	
	var records []string
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records,line)
	}

	var mustContain int

	// make a map of arrays: bright_white bag (key) contains
	// [shiny_gold] bag (element)
	bagRules := make(map[string]map[string]int)
	for _, rule := range records {
		bagName := rule[:strings.Index(rule, " contain")-5]
		bagsInBag := strings.Split(rule[strings.Index(rule, "contain")+8:],", ")
		// fmt.Println(bagName)
		if bagsInBag[0] != "no other bags." {
			bagSpecList := make(map[string]int)
			for _, bag := range bagsInBag {
				bagElementName := bag[strings.Index(bag, " ")+1:strings.Index(bag, " bag")]
				bagNum, _ := strconv.Atoi(bag[:strings.Index(bag, " ")])
				bagSpecList[bagElementName] = bagNum
			}
			bagRules[bagName] = bagSpecList
			// fmt.Println(bagName, ": ", bagRules[bagName])
		}
	}

	lastBags := make(map[string]int)
	lastBags["shiny gold"] = 1
	fmt.Println("shiny gold")
	var nextBags []map[string]int
	nextBags = append(nextBags, bagRules["shiny gold"])
	for _, element := range nextBags {
		for key, _ := range element {
			lastBags[key] = lastBags["shiny gold"] // nextBags is []map[string]int
		}
	}
	fmt.Println(lastBags)
	for len(nextBags) > 0 {
		fmt.Println(nextBags)
		copyNextBags := nextBags
		nextBags = nil
		for _, set_of_bags := range copyNextBags {
			for key, element := range set_of_bags {
				fmt.Println("looking at ", key, element)
				fmt.Println("should be multiplied by ", lastBags[key])
				mustContain += (element * lastBags[key])
				fmt.Println("current count is ", mustContain)
				for nextBagKey, nextBagElement := range bagRules[key] {
					nextBagSet := make(map[string]int)
					nextBagSet[nextBagKey] = nextBagElement
					nextBagSet[nextBagKey] = nextBagElement * lastBags[key]
					// fmt.Println("setting up next bag set ",nextBagSet)
					nextBags = append(nextBags, nextBagSet)
					lastBags[nextBagKey] = element
				}
			}
		}
		// fmt.Println("last bags set", lastBags)
	}

	return mustContain

}

func main() {
    // fmt.Println("part 1: ", q1("day7_input.csv"))
    fmt.Println("part 2: ", q2("day7_test.csv"))
}