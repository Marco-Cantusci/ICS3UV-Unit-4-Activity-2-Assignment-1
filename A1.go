/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-04
 * @fileoverview Prompts the sum of integers that are in range and out of range for a given range.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// declare variables
	var startString string
	var startNumber float64
	var endString string
	var endNumber float64
	var numberString string
	var numberNumber float64 = 1 // makes sure it runs once
	var inRange int
	var outRange int

	reader := bufio.NewReader(os.Stdin)

	// asking for range
	fmt.Print("Enter an integer for the start of the range: ")
	startString, _ = reader.ReadString('\n')
	startString = strings.TrimSpace(startString)
	startNumber, _ = strconv.ParseFloat(startString, 64)

	fmt.Print("Enter an integer for the end of the range: ")
	endString, _ = reader.ReadString('\n')
	endString = strings.TrimSpace(endString)
	endNumber, _ = strconv.ParseFloat(endString, 64)

	// do while loop
	for int(numberNumber) != 0 {

		fmt.Print("Enter an integer or zero(0) to end: ") // Asks for integers until 0 is entered.
		numberString, _ = reader.ReadString('\n')
		numberString = strings.TrimSpace(numberString)
		numberNumber, _ = strconv.ParseFloat(numberString, 64)

		if int(numberNumber) >= int(startNumber) && int(numberNumber) <= int(endNumber) { // adds to the total sum
			inRange += int(numberNumber)
		} else {
			outRange += int(numberNumber)
		}
	}

	fmt.Printf("The sum of the integers inside the range is %d \n", inRange)
	fmt.Printf("The sum of the integers outside the range is %d \n", outRange)

	fmt.Println("\nDone.")
}
