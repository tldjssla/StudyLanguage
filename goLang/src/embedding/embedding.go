package main

import (
	"fmt"
	_ "math"
	_ "unsafe"
)

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("zz")
}

type Student struct {
	//p      Person //has a
	Person //is a
	school string
	grade  int
}

func main() {
	var s Student
	s.Person.greeting()
	s.greeting()
}
