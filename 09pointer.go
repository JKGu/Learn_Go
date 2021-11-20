package main

import (
	"fmt"
)

func main() {
	var a int = 42  //& addressof operator
	var b *int = &a //b is a pointer to int, and pointing to a's address
	a = 27
	*b = 999
	fmt.Println(a, b, *b) //dereferencing operator for stored data

	x := [3]int{1, 2, 3}
	y := &x[0]
	z := &x[1] //-4? no go does not allow math op on pointers
	fmt.Println(x, y, z)

	//all assignments op in Go are copy operations
	//slices, maps contain internal pointers

}
