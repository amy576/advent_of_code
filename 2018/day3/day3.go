package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func q1(filename string) int {

	// make a 1000x1000 grid of empty strings
	cloth := [1000][1000]string{}

	// read each record
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)

Loop:
	for scanner.Scan() {
		line := scanner.Text()
		// get all the numbers
		// example: #1 @ 1,3: 4x4 where 1 is the x position and 3 is the y position
		atPos := strings.Index(line, "@")
		commaPos := strings.Index(line, ",")
		colonPos := strings.Index(line, ":")
		xPos := strings.Index(line, "x")
		xStart, _ := strconv.Atoi(line[atPos+2 : commaPos])
		yStart, _ := strconv.Atoi(line[commaPos+1 : colonPos])
		xLength, _ := strconv.Atoi(line[colonPos+2 : xPos])
		yLength, _ := strconv.Atoi(line[xPos+1:])

		// for each appropriate gridspace, fill in empty string with a "."
		// if it has already been filled with a "." or an "X", fill it with an "X"
		for y := 0; y < yLength; y++ {
			for x := 0; x < xLength; x++ {
				if yStart+y >= 1000 {
					fmt.Println("y is out of range: ", y)
					fmt.Println(line)
					fmt.Println(yStart)
					fmt.Println(yLength)
					break Loop
				}
				if xStart+x >= 1000 {
					fmt.Println("x is out of range: ", x)
					fmt.Println(line)
					fmt.Println(xStart)
					fmt.Println(xLength)
					break Loop
				}
				if cloth[yStart+y][xStart+x] == "." {
					cloth[yStart+y][xStart+x] = "X"
				}
				if cloth[yStart+y][xStart+x] == "" {
					cloth[yStart+y][xStart+x] = "."
				}
			}
		}
	}

	// count how many Y's there are
	overlap := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if cloth[y][x] == "X" {
				overlap++
			}
		}
	}
	return overlap
}

func q2(filename string) string {

	// make a 1000x1000 grid of empty strings
	cloth := make([][]string, 1000)
	for y := range cloth {
		cloth[y] = make([]string, 1000)
	}
	var collide bool
	var answer string

	// read each record
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		// get all the numbers
		// example: #1 @ 1,3: 4x4 where 1 is the x position and 3 is the y position
		atPos := strings.Index(line, "@")
		commaPos := strings.Index(line, ",")
		colonPos := strings.Index(line, ":")
		xPos := strings.Index(line, "x")
		xStart, _ := strconv.Atoi(line[atPos+2 : commaPos])
		yStart, _ := strconv.Atoi(line[commaPos+1 : colonPos])
		xLength, _ := strconv.Atoi(line[colonPos+2 : xPos])
		yLength, _ := strconv.Atoi(line[xPos+1:])

		// for each appropriate gridspace, fill in empty string with a "."
		// if it has already been filled with a "." or an "X", fill it with an "X"
		for y := 0; y < yLength; y++ {
			for x := 0; x < xLength; x++ {
				if cloth[yStart+y][xStart+x] == "" {
					cloth[yStart+y][xStart+x] = "."
				} else {
					cloth[yStart+y][xStart+x] = "X"
				}
			}
		}
	}

	_, err := f.Seek(0, 0)
	if err != nil {
		fmt.Print(err)
	}
	scanner = bufio.NewScanner(f)
Loop:
	for scanner.Scan() {
		line := scanner.Text()
		// get all the numbers
		// example: #1 @ 1,3: 4x4 where 1 is the x position and 3 is the y position
		atPos := strings.Index(line, "@")
		commaPos := strings.Index(line, ",")
		colonPos := strings.Index(line, ":")
		xPos := strings.Index(line, "x")
		xStart, _ := strconv.Atoi(line[atPos+2 : commaPos])
		yStart, _ := strconv.Atoi(line[commaPos+1 : colonPos])
		xLength, _ := strconv.Atoi(line[colonPos+2 : xPos])
		yLength, _ := strconv.Atoi(line[xPos+1:])

		id := line[1 : atPos-1]
		collide = false
		answer = id

		// check if the gridspace was filled twice
	Fill:
		for y := 0; y < yLength; y++ {
			for x := 0; x < xLength; x++ {
				if cloth[yStart+y][xStart+x] == "X" {
					collide = true
					answer = ""
					break Fill
				}
			}
		}

		if collide == false {
			answer = id
			break Loop
		}
	}
	return answer

}

func main() {
	fmt.Println("part1: ", q1("day3_input.txt"))
	fmt.Println("part2: ", q2("day3_input.txt"))
}
