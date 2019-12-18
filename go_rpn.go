package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Reverse Polish Notation Calculator")
		fmt.Print("Enter in calculation: ")
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		// Need the space at the end for the 'readString' delimeter
		input = input + string(' ')

		result := rpnCalc(input)

		fmt.Printf("Entered '%s', Result '%f'\n", input, result)
	}
}

func rpnCalc(s string) float64 {
	var items []float64

	reader := bufio.NewReader(strings.NewReader(s))

	for i := 0; ; i++ {

		// Reach each part of the string
		readString, err := reader.ReadString(' ')
		readString = strings.Trim(readString, " ")

		if err == nil {
			fmt.Println("\nPrinting ", i, " parameter ", readString)

			// Check for number
			if strings.ContainsAny(readString, "0123456789") {
				fmt.Println("Checking number ", readString)
				// Convert string to int
				num, err := strconv.ParseFloat(readString, 64)

				if err == nil {
					// Add number to stack
					items = append(items, num)
					fmt.Println("items[", len(items)-1, "] = ", num)

				} else {
					fmt.Println("Error parsing number")
					log.Fatal(err)
				}
			} else if strings.ContainsAny(readString, "+-*x/") { // Check for operator
				// Pop
				i := len(items) - 1
				num1 := items[i]
				items = items[:i]

				// Pop
				i = len(items) - 1
				num2 := items[i]
				items = items[:i]

				if readString == "+" {
					items = append(items, num1+num2)
					fmt.Println("Added ", num1, " and ", num2, " result = ", items[len(items)-1])

				} else if readString == "-" {
					items = append(items, num1-num2)
					fmt.Println("Subtracted ", num1, " and ", num2, " result = ", items[len(items)-1])

				} else if readString == "*" || readString == "x" {
					items = append(items, num1*num2)
					fmt.Println("Multiplied ", num1, " and ", num2, " result = ", items[len(items)-1])

				} else if readString == "/" {
					items = append(items, num1/num2)
					fmt.Println("Divided ", num1, " and ", num2, " result = ", items[len(items)-1])

				} else {
					fmt.Println("Somehow we got here, check ", readString)
					os.Exit(1)
				}
			} else { // Handle off case
				fmt.Println("Invalid characters ", readString)
				return 0
			}
		} else if err == io.EOF { // Calculate answer
			return items[len(items)-1]
		} else {
			fmt.Println("Error")
			log.Fatal(err)
		}
	}
}
