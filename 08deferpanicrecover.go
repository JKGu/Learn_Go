package main

import (
	"fmt"
	"log"
)

func main() {
	a := 5
	fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println(a)
	fmt.Println("4")
	a = 999
	//print 1 4 5 3 2; defer will run right before function return
	//multiple defer: LIFO
	//used when you open a resource, and make sure to close it when done
	//defer is eager evaluated

	//panic
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	panic("exception") //happens after defer
	fmt.Println("end") //unreachable
}
