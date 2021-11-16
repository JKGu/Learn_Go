package main

import (
	"fmt"     //format
	"strconv" //string conversion
)

// var in the package
var (
	x    float32 = 100
	name string  = "jay"
)
var y int = 999 // := cannot be used here

//visibility
var aX int = 100 // scoped to the package
var AX int = 100 // scoped globally

func main() {
	//comment
	//var i float32 , assign later
	var i float32 = 40 // var always has to be used; scoped to the main block
	j := 43
	j = 99
	fmt.Println(i + x)
	fmt.Printf("%v, %T", i, j)

	// shadowing, x covered by innermost scope
	// cannot redelcare var but shadow ok
	fmt.Println(x)
	var x int = 100
	// x := 100   won't work
	fmt.Println(x)

	i = float32(aX)           //conversion
	var s string = string(AX) // returns "d" (ASCII)
	s = strconv.Itoa(AX)      //Integer to Ascii
	fmt.Println(s)
}
