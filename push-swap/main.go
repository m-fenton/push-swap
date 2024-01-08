package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stacks struct {
	stackA, stackB []int
}

func main() {
	var stackA []int
	var stackB []int
	//var sliceToSort []int

	args := argChecker()
	duplicateChecker(args)
	stackA = inputStringToSliceofInts(args)
	sliceToSort := inputStringToSliceofInts(args)
	sortedStackA := sortSliceofInts(sliceToSort)
	IntArrayEquals(stackA, sortedStackA)
	input := Stacks{stackA, stackB}

	if len(stackA) == 3 {
		input.three()
	}
	if len(stackA) == 5 {
		input.five(sortedStackA)
	}
	if len(stackA) == 6 {
		input.six(sortedStackA)
	}
	if len(stackA) == 100 {
		input.hundred(sortedStackA)
	}

	for i := 0; i < len(input.stackA); i++ {
		fmt.Print(input.stackA[i], " ")
	}
	fmt.Println()
}

// reads the argument and returns it as an array of ints
func inputStringToSliceofInts(input string) []int {

	var intArgsArray []int

	argsStrings := strings.Split(input, " ")

	for i := 0; i < len(argsStrings); i++ {

		intArg, _ := strconv.Atoi(argsStrings[i])
		intArgsArray = append(intArgsArray, intArg)
	}
	return intArgsArray
}

// }

// rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func (s *Stacks) ra() {
	firstInt := s.stackA[0]
	s.stackA = s.stackA[1:]
	s.stackA = append(s.stackA, firstInt)
	fmt.Println("ra")
}

// rotate stack b
func (s *Stacks) rb() {

	firstInt := s.stackB[0]
	s.stackB = s.stackB[1:]
	s.stackB = append(s.stackB, firstInt)
	fmt.Println("rb")
}

// execute ra and rb
func (s *Stacks) rr() {

	s.ra()
	s.rb()
	fmt.Println("ra")
	fmt.Println("ra")
}

// reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func (s *Stacks) rra() {
	var newArray []int
	lastInt := s.stackA[len(s.stackA)-1]
	newArray = append(newArray, lastInt)
	s.stackA = s.stackA[:len(s.stackA)-1]
	s.stackA = append(newArray, s.stackA...)
	fmt.Println("rra")
}

// reverse rotate b
func (s *Stacks) rrb() {
	var newArray []int
	lastInt := s.stackB[len(s.stackB)-1]
	newArray = append(newArray, lastInt)
	s.stackB = s.stackB[:len(s.stackB)-1]
	s.stackB = append(newArray, s.stackB...)
	fmt.Println("rrb")
}

// execute rra and rrb
func (s *Stacks) rrr() {

	s.rra()
	s.rrb()
	fmt.Println("rra")
	fmt.Println("rrb")
}

// push the top first element of stack b to stack a
func (s *Stacks) pa() {
	var newArray []int
	firstIntB := s.stackB[0]
	newArray = append(newArray, firstIntB)
	s.stackA = append(newArray, s.stackA...)
	s.stackB = s.stackB[1:]
	fmt.Println("pa")
}

func (s *Stacks) pb() {
	var newArray []int
	firstIntA := s.stackA[0]
	newArray = append(newArray, firstIntA)
	s.stackB = append(newArray, s.stackB...)
	s.stackA = s.stackA[1:]
	fmt.Println("pb")
}

// swap first 2 elements of stack a
func (s *Stacks) sa() {

	indexOne := s.stackA[0]
	indexTwo := s.stackA[1]

	x := indexOne
	indexOne = indexTwo
	indexTwo = x

	s.stackA = append([]int{indexOne, indexTwo}, s.stackA[2:]...)
	fmt.Println("sa")
}

// swap first 2 elements of stack b
func (s *Stacks) sb() {

	indexOne := s.stackB[0]
	indexTwo := s.stackB[1]

	x := indexOne
	indexOne = indexTwo
	indexTwo = x

	s.stackB = append([]int{indexOne, indexTwo}, s.stackB[2:]...)
	fmt.Println("sb")
}
