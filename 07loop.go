package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	} // i will be 5
	for i < 5 {
		fmt.Println(i)
		i++
	}
	for i < 5 {
		fmt.Println(i)
		i++
	}
	i = 100
	for { //while true
		if i == 1000 {
			break
		}
		i++
	}

Loop: //breakout label for nested loops
	for i := 1; i < 5; i++ {
		for j := 5; j < 100; j++ {
			if j == 90 {
				break Loop
			}
		}
	}

	s := []int{1, 2, 3, 4}
	for k, v := range s {
		fmt.Println(k, v)
	}
	for k := range s {
		fmt.Println(k)
	}
	for _, v := range s {
		fmt.Println(v)
	}

}
