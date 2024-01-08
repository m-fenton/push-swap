package main

// sorts five numbers
func (s *Stacks) five(sortedStackA []int) {

	s.lowestNumber(sortedStackA)
	s.lowestNumber(sortedStackA[1:])
	s.three()
	s.pa()
	s.pa()

}

// finds the lowest number in stack A, then pushes it over
func (s *Stacks) lowestNumber(sortedStackA []int) {

	var index int
	min := sortedStackA[0]

	for i := 0; i < len(s.stackA); i++ {
		if min == s.stackA[i] {
			index = i

			if index < len(s.stackA)/2 {
				for j := 0; j < index; j++ {
					s.ra()
				}
				s.pb()
			} else {

				for k := 0; k < len(s.stackA)-index; k++ {
					s.rra()
				}
				s.pb()

			}

		}

	}
}
