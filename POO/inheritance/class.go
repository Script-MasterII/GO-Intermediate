package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Employee struct {
	id int
}
type FullTimeEmployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.age = 12
	ftEmployee.id = 8
	ftEmployee.name = "Juan"

	fmt.Println(ftEmployee)
}
