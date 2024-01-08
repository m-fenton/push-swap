package main

// sorts 6 numbers
func (s *Stacks) six(sortedStackA []int) {
	chunk1 := sortedStackA[:3]
	s.chunkBlaster(chunk1)
	s.three()
	for stackReformer := 0; len(s.stackB) != 0; stackReformer++ {
		if s.intIsHigherThanRestOfStack() {
			s.pa()
		} else {
			s.rb()
		}
	}

}

// performs steps 1 - 4 for the given chunk
func (s *Stacks) chunkBlaster(chunk []int) {

	for chunkTurn := 0; chunkTurn < len(chunk); chunkTurn++ {
		// step 1
		// Find holdfirst position
		_, hfInteger := s.isIntInChunkFromTop(chunk)
		// step 2
		// Find holdsecond position
		_, hsInteger := s.isIntInChunkFromBottom(chunk)
		// step 3
		// if it's quicker to move the number to the top of stack A then do that
		s.quickerUpOrDown(hfInteger, hsInteger)
		// step 4
		// checks for a sensible spot in B by checking if the number is lower/higher than all other numbers in B
		s.FindCorrectSpotinB()
	}
}

// searches stackA from the top until it finds a number that is in the inserted chunk
func (s *Stacks) isIntInChunkFromTop(chunk []int) (int, int) {
	var hold_first int
	var hfInteger int
	for i := 0; i < len(s.stackA); i++ {
		if isIntInChunk(s.stackA[i], chunk) {
			hold_first = s.stackA[i]
			hfInteger = i

			return hold_first, hfInteger
		}
	}
	return hold_first, hfInteger
}

// searches stackA from the bottom until it finds a number that is in the inserted chunk
func (s *Stacks) isIntInChunkFromBottom(chunk []int) (int, int) {
	var hold_second int
	var hsInteger int
	for i := len(s.stackA) - 1; i > 0; i-- {
		if isIntInChunk(s.stackA[i], chunk) {
			hold_second = s.stackA[i]
			hsInteger = len(s.stackA) - i

			return hold_second, hsInteger
		}
	}
	return hold_second, hsInteger
}

// Finds the first number from chunk 1, if any, that is in stack A
func isIntInChunk(inputInt int, chunk []int) bool {

	for arrayOfChunkNumber := 0; arrayOfChunkNumber < len(chunk); arrayOfChunkNumber++ {
		if inputInt == chunk[arrayOfChunkNumber] {
			return true
		}

	}
	return false
}

// checks to see if int that is about to be pushed is less than every number in B
func (s *Stacks) intIsLessThanStack(inputInt int) bool {

	for arrayOfChunkNumber := 0; arrayOfChunkNumber < len(s.stackB); arrayOfChunkNumber++ {
		if inputInt > s.stackB[arrayOfChunkNumber] {
			return false
		}

	}
	return true
}

// checks to see if int that is about to be pushed is higher than every number in B
func (s *Stacks) intIsHigherThanStack(inputInt int) bool {

	for arrayOfChunkNumber := 0; arrayOfChunkNumber < len(s.stackB); arrayOfChunkNumber++ {
		if inputInt < s.stackB[arrayOfChunkNumber] {
			return false
		}

	}
	return true
}

func (s *Stacks) quickerUpOrDown(turnsToTopUp int, turnsToTopDown int) {
	// if it's quicker to move the number to the top of stack A then do that
	if turnsToTopUp <= turnsToTopDown {
		for numberOfRAs := 0; numberOfRAs < turnsToTopUp; numberOfRAs++ {
			s.ra()
		}
		// else move it downwards
	} else {
		for numberOfRRAs := 0; numberOfRRAs < turnsToTopDown; numberOfRRAs++ {
			s.rra()
		}
	}
}

//Tries to make sure that the number from A is pushed to an efficient place in Stack B
func (s *Stacks) FindCorrectSpotinB() {

	if len(s.stackB) > 1 {
		// if the number at the top of A is less than every number in B then push it ver
		if s.intIsLessThanStack(s.stackA[0]) {
			s.pb()

		} else if s.intIsHigherThanStack(s.stackA[0]) {

			sortedB := sortSliceofInts(s.stackB)

			for i := 0; s.stackB[0] == sortedB[0]; i++ {
				s.rb()

			}
			s.pb()

			// rotate stack B until it's the number in A is larger than the top number of B
			// ... might need some work in order to avoid infinite loops
			for i := 0; s.stackA[0] < s.stackB[0]; i++ {
				s.rb()

			}
		} else {
			for rotatingB := 0; s.stackA[0] > s.stackB[0]; rotatingB++ {
				s.rb()

			}
			s.pb()

		}

		// once  the number at the top of stack A is larger than the number at the top of stack B then pb

	} else {
		// if stack B is 0 or 1 numbers long then just push
		s.pb()

	}
}

func (s *Stacks) intIsHigherThanRestOfStack() bool {

	for i := 0; i < len(s.stackB); i++ {
		if s.stackB[0] < s.stackB[i] {
			return false
		}

	}
	return true
}
