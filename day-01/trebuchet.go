package main

import (
	"bufio"
	"fmt"
	"os"
)

func processLine(line string) int {
	// return dummy data to test, will add logic later
	return 77
}

func main() {
	// path for file with input data
	input_file := "input.txt"

	// open file
	fp, err := os.Open(input_file)

	// handle error
	if err != nil {
		fmt.Println("Error oprning file: ", err)
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
