package main

import (
	"fmt" //format
)

func main() {
	//ARRAY -fixed size
	grades := [...]int{97, 85, 93}
	//grades := [3]int
	//var grades [3]int
	//matrix := [3][3]int
	copy := grades //make a copy of the entire array
	copy[2] = 5
	pointer := &grades //pointer reference
	pointer[1] = 6
	fmt.Printf("Grades: %v  Length: %v \n", grades, len(grades))
	fmt.Printf("%v\n%v\n", copy, pointer)

	// SLICE  backed by array but dynamic
	s := []int{97, 85, 93}
	s2 := s // slice are referenced
	s2[2] = 5
	// a := s[3:6] slices of s
	fmt.Printf("%v\n%v\n", s, s2)

	//a := make([]int, 3)
	a := make([]int, 3, 100)
	a = append(a, 1) //if not fit, copy everything to a larger array
	fmt.Println(a)
	fmt.Printf("%v %v", len(a), cap(a))
	a = append(a, []int{1, 2, 3, 4}...) //... spread operator, con two slices
	fmt.Println(a)
}
