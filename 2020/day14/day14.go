package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]string) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func q1(lines []string) int {
	memValues := make(map[int]int)
	var mask [36]string
	var masks int

	for _, line := range lines {
		// The bitmask is always given as a string of 36 bits, written
		// with the most significant bit (representing 2^35) on the left
		// and the least significant bit (2^0, that is, the 1s bit) on the
		// right. 
		if line[:4] == "mask" {
			for i, char := range line[7:] {
				mask[i] = string(char)
			}
			masks++
			// fmt.Println(mask)

		} else if line[:3] == "mem" {
			address, _ := strconv.Atoi(line[strings.Index(line, "[")+1:strings.Index(line, "]")])
			val, _ := strconv.Atoi(line[strings.Index(line, "=")+2:])
			valBin := strconv.FormatInt(int64(val), 2)
			// fmt.Println(valBin)

			numLeadingZeroes := 36 - len(valBin)
			var valBinArr [36]string
			for j := 0; j < numLeadingZeroes; j++ {
				valBinArr[j] = "0"
			}
			for k, char := range valBin {
				valBinArr[numLeadingZeroes+k] = string(char)
			}

			// The current bitmask is applied to values immediately
			// before they are written to memory: a 0 or 1 overwrites the
			// corresponding bit in the value, while an X leaves the bit in
			// the value unchanged.
			var finalValBinArr [36]string
			for i := 0; i < 36; i++ {
				if mask[i] == "X" {
					finalValBinArr[i] = valBinArr[i]
				} else {
					finalValBinArr[i] = mask[i]
				}
			}
			finalValBin := strings.Join(finalValBinArr[:], "")
			finalVal, _ := strconv.ParseInt(finalValBin, 2, 64)
			memValues[address] = int(finalVal)
			// fmt.Println(int(finalVal), "saved to address", address)
		}
	}

	var total int
	for _, element := range memValues {
		total += element
	}
	return total
}

func q2(lines []string) int {
	memValues := make(map[int]int)
	var mask [36]string
	var masks [][]string
	var floating []int
	numPermutations := 1

	for _, line := range lines {
		// The bitmask is always given as a string of 36 bits, written
		// with the most significant bit (representing 2^35) on the left
		// and the least significant bit (2^0, that is, the 1s bit) on the
		// right. 
		if line[:4] == "mask" {
			for i, char := range line[7:] {
				mask[i] = string(char)
			}
			// make mask permutations
			// fmt.Println("starting permutations")
			fmt.Println(numPermutations)
			masks, floating = permutateMasks(strings.Join(mask[:],""))
			numPermutations++

		} else if line[:3] == "mem" {
			address, _ := strconv.Atoi(line[strings.Index(line, "[")+1:strings.Index(line, "]")])
			val, _ := strconv.Atoi(line[strings.Index(line, "=")+2:])
			addressBin := strconv.FormatInt(int64(address), 2)

			numLeadingZeroes := 36 - len(addressBin)
			var addressBinArr [36]string
			for j := 0; j < numLeadingZeroes; j++ {
				addressBinArr[j] = "0"
			}
			for k, char := range addressBin {
				addressBinArr[numLeadingZeroes+k] = string(char)
			}

			// If the bitmask bit is 0, the corresponding memory address bit is unchanged.
			// If the bitmask bit is 1, the corresponding memory address bit is overwritten with 1.
			// If the bitmask bit is X, the corresponding memory address bit is floating.
			for y := 0; y < 36; y++ {
				if mask[y] == "1" {
					addressBinArr[y] = "1"
				}
			}
			var finalAddrBinArr []int
			for _, maskX := range masks {
				eachAddrBinArr := addressBinArr
				for i := 0; i < 36; i++ {
					for _, j := range floating {
						if i == j {
							eachAddrBinArr[i] = maskX[i]
						}
					}
				}
				finalAddrBin := strings.Join(eachAddrBinArr[:], "")
				finalAddr, _ := strconv.ParseInt(finalAddrBin, 2, 64)
				finalAddrBinArr = append(finalAddrBinArr, int(finalAddr))
			}

			for _, addr := range finalAddrBinArr {
				memValues[addr] = int(val)
			}
			// fmt.Println(int(finalVal), "saved to address", address)
		}
	}

	var total int
	for _, element := range memValues {
		total += element
	}
	return total
}

func permutateMasks(inputMask string) ([][]string, []int) {
	var permutatedMasks [][]string
	xNum := strings.Count(inputMask,"X")
	var xPos []int
	for i, char := range inputMask {
		if string(char) == "X" {
			xPos = append(xPos, i)
		}
	}
	var masksIn []string
	masksIn = append(masksIn, inputMask)
	// var masksOut []string
	fmt.Println("starting to generate binaries")
	binaries := generateBinary(masksIn, xNum)
	// fmt.Println("completed binary generation")
	// fmt.Println(binaries)
	for _, bin := range binaries {
		var tmpMask []string
		for _, char := range bin {
			tmpMask = append(tmpMask, string(char))
		}
		permutatedMasks = append(permutatedMasks, tmpMask)
	}

	return permutatedMasks, xPos
}

func generateBinary(arr []string, xNum int) []string {
	if len(arr) == 1<<xNum {
		// fmt.Println("returning arr", arr)
		return arr
	} else {
		Recurse:
			for i, bin := range arr {
				// fmt.Println(i, arr)
				// fmt.Println(bin)
				xPos := strings.Index(bin, "X")
				if xPos == -1 {
					// fmt.Println("no more x's:", arr)
					break Recurse
				}
				arr[i] = bin[:xPos] + "0" + bin[xPos+1:]
				arr = append(arr, bin[:xPos] + "1" + bin[xPos+1:])
				// fmt.Println(arr[i])
				// fmt.Println(arr)
			}
			arr = generateBinary(arr, xNum)
			return arr
		}
}

func main() {
	filename := "day14_input.csv"
	bits := readInput(filename)

	// This program starts by specifying a bitmask (mask = ....).
	// The program then attempts to write values to certain memory addresses.
	// What is the sum of all values left in memory after it completes?
	q1 := q1(bits)
	fmt.Println("part 1: ", q1)

	q2 := q2(bits)
	fmt.Println("part 2: ", q2)
}