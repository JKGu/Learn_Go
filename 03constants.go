package main

import (
	"fmt"
)

const (
	bb = iota
	cc
	dd
)
const (
	aa = iota
)

func main() {
	//typed constants are like immutable var
	//untyped constants are like literals
	const myConst = 42
	// const myConst int = 42+1 //cannot evaluate at compile time
	var b int16 = 27
	fmt.Println(myConst + b) // compiler can infer that myConst is int 16
	var c int32 = 6
	fmt.Println(myConst + c) //still works, compiler simply replaces constant symbols

	fmt.Println("---------")
	fmt.Println(bb) //0
	fmt.Println(cc) //1
	fmt.Println(dd) //2
	fmt.Println(aa) //0
}
