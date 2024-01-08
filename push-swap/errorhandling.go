package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// This function checks that all of the items in os.Args[1] are numbers then returns, if not it prints an error messsage.
func argChecker() string {
	var args string

	if len(os.Args) == 2 {
		args = os.Args[1]
	} else if len(os.Args) == 1 {
		os.Exit(0)
	} else {
		args = strings.Join(os.Args[1:], " ")
	}

	// Make the string into an array of runes
	r := []rune(args)

	for i := 0; i < len(r)-1; i++ {
		// Loop throught the rune array and search for items that are not nubers and not a space.
		if r[i] < 48 && r[i] != 32 || r[i] > 57 {
			// If the statement is true then print the error message below and exit.
			fmt.Println("ERROR: invalid data format, args are not all integers.")
			os.Exit(1)
		}
	}
	return args
}

// Sorts the array of ints, for error handling
func sortSliceofInts(argsSlice []int) []int {

	sort.Ints(argsSlice)
	return argsSlice
}

// Checks if input is already sorted
func IntArrayEquals(unsorted []int, sorted []int) bool {
	if len(unsorted) != len(sorted) {
		//fmt.Println("They are not the same 1.")
		return false
	}
	for i, v := range unsorted {
		if v != sorted[i] {
			//fmt.Println("They are not the same 2.")
			return false
		}
	}
	// If true then exit with no error message as requested in task.

	os.Exit(0)
	return true
}

// This function checks if the slice of string contains any duplicate numbers, if it does then print the error message.
func duplicateChecker(data string) {

	splitData := strings.Split(data, " ")
	// I loop through our loop looking for repeat numbers.
	for i := 0; i < len(splitData); i++ {
		counter := 0
		for j := 0; j < len(splitData); j++ {
			// As the same number exists in each loop once then we will always find one occurence.
			if splitData[i] == splitData[j] {
				counter++
				// We add the occurences of duplicates to the counter and if the counter reaches two then we have a duplicate.
				// We then print the error message below and exit.
				if counter == 2 {
					fmt.Println("ERROR: invalid data format, there are duplicate integers.")
					os.Exit(2)
				}
			}
		}
	}
}
