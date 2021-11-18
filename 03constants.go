package main

import (
	"fmt"
)

const a int16 = 27

func main() {
	const myConst = 42
	// const myConst int = 42+1 //cannot evaluate at compile time
	var b int16 = 27
	fmt.Println(myConst + b) // compiler can infer that myConst is int 16
	var c int32 = 6
	fmt.Println(myConst + c)
}
