package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func processLine(line string) int {
	var firstNum, lastNum int = 0, 0
	for _, char := range line {
		if unicode.IsDigit(char) {
			numericValue, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println("Error converting string to int: ", err)
				return 0
			}
			if firstNum == 0 {
				firstNum = numericValue
			}
			lastNum = numericValue
		}
	}
	fmt.Println("Line: ", line, "First: ", firstNum, "Last: ", lastNum)
	return firstNum*10 + lastNum
}

func main() {
	// path for file with input data
	input_file := "input.txt"

	// open file
	fp, err := os.Open(input_file)

	// handle error
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer fp.Close() // ensures file is closed when main function is done or if there is an error in processing

	// create buffer to read lines in the file
	lineReader := bufio.NewScanner(fp)

	// create global sum variable
	var totalSum int = 0

	// loop through the lines using the buffer
	for lineReader.Scan() {

		line := lineReader.Text()
		var calibrationValue int = processLine(line)
		totalSum += calibrationValue
	}

	// handle errors
	if err := lineReader.Err(); err != nil {
		fmt.Println("Error while reading file: ", err)
		return
	}

	fmt.Println("The sum of all calibration values is: ", totalSum)
}
