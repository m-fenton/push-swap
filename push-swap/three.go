package main

func (s *Stacks) three() {

	if s.stackA[0] > s.stackA[1] && s.stackA[1] < s.stackA[2] && s.stackA[0] < s.stackA[2] {
		s.sa()
	} else if s.stackA[0] > s.stackA[1] && s.stackA[1] > s.stackA[2] && s.stackA[0] > s.stackA[2] {
		s.sa()
		s.rra()
	} else if s.stackA[0] > s.stackA[1] && s.stackA[1] < s.stackA[2] && s.stackA[0] > s.stackA[2] {
		s.ra()
	} else if s.stackA[0] < s.stackA[1] && s.stackA[1] > s.stackA[2] && s.stackA[0] < s.stackA[2] {
		s.sa()
		s.ra()
	} else if s.stackA[0] < s.stackA[1] && s.stackA[1] > s.stackA[2] && s.stackA[0] > s.stackA[2] {
		s.rra()
	}
}
