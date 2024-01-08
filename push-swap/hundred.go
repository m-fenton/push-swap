package main

//searches stackA from the top until it finds a number that is in the inserted chunk, returns that number and its position in its []int
func (s *Stacks) holdFirst(chunk1 []int) (int, int) {
	var hold_first int
	var i int

	for i := 0; i < len(s.stackA); i++ {

		for j := 0; j < len(chunk1); j++ {

			if s.stackA[i] == chunk1[j] {

				hold_first = s.stackA[i]

				return hold_first, i
			}
		}
	}

	return hold_first, i
}

//searches stackA from the bottom until it finds a number that is in the inserted chunk, returns that number and its position in its []int
func (s *Stacks) holdSecond(chunk1 []int) (int, int) {
	var hold_second int
	var i int
	//chunk1 := sortedStackA[0:20]

	for i := len(s.stackA) - 1; i > 0; i-- { // step 2

		for j := 0; j < len(chunk1); j++ {
			if s.stackA[i] == chunk1[j] {

				hold_second = s.stackA[i]

				//hold_second = chunk1[j]
				//fmt.Println("This is index [i]:", i)
				return hold_second, i
			}
		}
	}

	return hold_second, i
}

func (s *Stacks) compareIndex2(first int, second int, stackA []int) bool {

	if first < len(stackA)-second {
		for i := 0; i < first; i++ {
			s.ra()

		}
		s.pb()
		return true
	}

	x := len(stackA) - second

	for i := 0; i < x; i++ {
		s.rra()

	}
	s.pb()

	return false
}

// sorts 100 ints
func (s *Stacks) hundred(sortedStackA []int) {

	chunk1 := sortedStackA[0:20]
	chunk2 := sortedStackA[20:40]
	chunk3 := sortedStackA[40:60]
	chunk4 := sortedStackA[60:80]
	chunk5 := sortedStackA[80:100]

	//step 6 (do steps 1-4 for each chunk)
	s.chunkBlasterHundred(chunk1)
	s.chunkBlasterHundred(chunk2)
	s.chunkBlasterHundred(chunk3)
	s.chunkBlasterHundred(chunk4)
	s.chunkBlasterHundred(chunk5)

	for i := 0; i < 100; i++ {

		s.largestNumber(sortedStackA)
		sortedStackA = sortedStackA[:len(sortedStackA)-1]
	}
}

// finds each number in the chunk within A, then pushes it into stackB
func (s *Stacks) chunkBlasterHundred(chunk []int) {
	//step 5 - does steps 1 - 4 for each number in the inserted chunk
	for i := 0; i < len(chunk); i++ {
		//step 1 - scan stackA from top to confirm one of the numbers from chunk1 exist there and call it hold-first (holdFirst)
		_, first := s.holdFirst(chunk)
		//step 2 - scan stackA from the bottom to confirm one of the numbers from chunk1 exist there and call it hold-second (holdSecond)
		_, second := s.holdSecond(chunk)
		//step 3 and 4 - compare the moves necessary to get each of these to the top of the stack (compareIndex2)
		s.compareIndex2(first, second, s.stackA)
	}
}

//step 7 - find biggest number in stack B, get it's index, get it to the top either ra or rra similar to above functions and then push to a (pa)
// keep repeating until B is empty
func (s *Stacks) largestNumber(sortedStackA []int) {

	var index int
	max := sortedStackA[len(sortedStackA)-1]

	for i := 0; i < len(s.stackB); i++ {
		if max == s.stackB[i] {
			index = i

			if index < len(s.stackB)/2 {
				for j := 0; j < index; j++ {
					s.rb()

				}
				s.pa()

			} else {
				for k := 0; k < len(s.stackB)-index; k++ {
					s.rrb()

				}
				s.pa()

			}

		}

	}
}
