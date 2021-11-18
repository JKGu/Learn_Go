package main

import (
	"fmt"
)

func main() {
	var n bool = false
	m := 1 == 2
	fmt.Println(n, m)

	// int8, int16, int32, int64

	// uint8, 16, 32

	// + - * / %

	//(int will be viewed as binary)
	// and &  or |  xor ^  andnot &^

	a := 8              //2^3
	fmt.Println(a << 3) //2^3 * 2^3  (>>)

	// float32 64
	x := 13.7e72 //will overflow if x declared as float32
	fmt.Println(x)

	//string immutable
	s := "dafljdklfads"
	b := []byte(s)
	fmt.Printf("%v, %T\n", string(s[2]), s[2]) //unit8 byt
	fmt.Println(b)
	fmt.Println(s[5:9])
	//var r rune = 'a' //int32
	//fmt.Printf("&v,%T\n", r, r)

}
