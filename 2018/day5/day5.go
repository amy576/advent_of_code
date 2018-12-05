// read it
// find Aa, aA, bB, Bb, cC, or Cc and remove
// keep removing until the difference between the length of the last round and the new round stops changing

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func q1(filename string) int {

	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)

	var line string

	for scanner.Scan() {
		line = scanner.Text()
	}
	f.Close()

	v := len(line)
	alpha := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	still_replacing := true

	for still_replacing {
		for _, beta := range alpha {
			line = strings.Replace(line, beta+strings.ToUpper(beta), "", -1)
			line = strings.Replace(line, strings.ToUpper(beta)+beta, "", -1)
		}
		if v != len(line) {
			still_replacing = true
			v = len(line)
		} else {
			still_replacing = false
		}
	}
	return v
}

func q2(filename string) int {

	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)

	var original string

	for scanner.Scan() {
		original = scanner.Text()
	}
	f.Close()
	line := original

	alpha := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	shortest := 50000

	for _, beta := range alpha {
		// reset
		v := len(original)
		line = original
		still_replacing := true

		// remove one "unit"
		line = strings.Replace(line, beta, "", -1)
		line = strings.Replace(line, strings.ToUpper(beta), "", -1)
		for still_replacing {
			for _, beta := range alpha {
				line = strings.Replace(line, beta+strings.ToUpper(beta), "", -1)
				line = strings.Replace(line, strings.ToUpper(beta)+beta, "", -1)
			}
			if v != len(line) {
				still_replacing = true
				v = len(line)
			} else {
				still_replacing = false
			}
		}
		if v < shortest {
			shortest = v
		}
	}
	return shortest
}

func main() {
	fmt.Println("part1: ", q1("day5_input.txt"))
	fmt.Println("part2: ", q2("day5_input.txt"))
}
