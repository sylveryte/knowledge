package main

import "fmt"

// type Stack[T any] struct {
// 	vals []T
// }

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) push(val T) {
	s.vals = append(s.vals, val)
}
func (s *Stack[T]) print() {
	fmt.Println(s.vals)
}
func (s *Stack[T]) pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	r := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return r, true
}

func (s *Stack[T]) contains(val T) bool {

	for _, v := range s.vals {
    // doens't work with any
		if v == val {
			return true
		}
	}

	return false
}

func main() {
	j := Stack[int]{
		vals: []int{},
	}
	j.push(34)
	j.push(23)
	j.push(87)
	j.push(948)
	j.push(21)
	j.print()
	v, i := j.pop()
	fmt.Println("popping ", v, i)
	j.print()
  fmt.Println("23 exists? ",j.contains(23))
}
