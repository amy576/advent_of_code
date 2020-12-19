package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("a((aa|bb)(ab|ba)|(ab|ba)(aa|bb))b")
	s1 := "aaaabbb"
	containsNumber := re.FindAllString(s1, -1)
	fmt.Println(containsNumber)
}
