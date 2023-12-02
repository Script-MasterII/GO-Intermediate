package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,
		name: name,
	}
}

func main() {
	// 1
	e := Employee{}
	fmt.Printf("%v\n", e)
	// 2
	e2 := Employee{
		id:   1,
		name: "Juan",
	}
	fmt.Printf("%v\n", e2)
	// 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 1
	e3.name = "Carlos"
	fmt.Printf("%v\n", *e3)
	// 4
	e4 := NewEmployee(1, "Enrique")
	fmt.Printf("%v\n", *e4)
}
