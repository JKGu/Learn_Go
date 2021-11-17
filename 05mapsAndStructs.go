package main

import (
	"fmt"
	"reflect"
)

//STRUCT
type Doctor struct {
	number int
	name   string
	Id     int //capital export
}

// go doesn't have inheritance, but compositions (ISA relation) (embedding)
type Animal struct {
	Type string `required max:"100"` //tag
}
type Bird struct {
	Animal //"has an animal (characteristic)"
	CanFly bool
}

// embedding is not good for behavior modeling, use interface instead

func main() {

	//MAP

	m := make(map[string]int)
	// slice cannot be key; array can be key
	m = map[string]int{
		"s1":         555,
		"s2":         99021,
		"dsfjkldsfl": 321312,
	}
	m["xxx"] = 324234
	delete(m, "xxx")
	v, ok := m["55"] // use ok to test if in this map
	fmt.Println(m["s1"], v, ok)
	// passing a map: pass by reference

	d := Doctor{
		number: 5,
		name:   "Jon",
		Id:     6,
	} // passing a struct is passing the copy of a struct
	fmt.Println(d, d.name)

	b := Bird{}
	b.CanFly = true
	fmt.Println(b.Type)

	b2 := Bird{
		Animal: Animal{Type: "dsfsdf"},
		CanFly: false,
	}
	fmt.Println(b2.Type)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Type")
	fmt.Println(field.Tag)
}
