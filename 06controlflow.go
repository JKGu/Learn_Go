package main

import (
	"fmt"
)

func main() {
	// && || !
	if 6 < 8 {
		fmt.Println("hi")
	} else if 5 < 4 {
	} else {
	}

	switch 222 {
	case 1111:
		//
	case 222:
		//
	case 1, 55, 6:
		//
		fallthrough //keyword
	case 0:
		//still get evaluated due to fallthrough
	default:
		//
	}

	switch i := 2 + 3; i { //i get matched
	case 5:
	default:
	}

	i := 999
	switch {
	case i < 1000:
		//
	case i < 1000000: //overlap allowed here
		//
	}

	var ii interface{} = 1.0
	switch ii.(type) {
	case int:
		fmt.Println("int")
		if true {
			break
		}
	}

	m := map[string]int{
		"s1":         555,
		"s2":         99021,
		"dsfjkldsfl": 321312,
	}
	if pop, ok := m["s1"]; ok {
		fmt.Println(pop)
	}

}
