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
	endDate string
}

func (ftEployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporatyEmployee struct {
	Person
	Employee
	taxtRate int
}

func (tEployee TemporatyEmployee) getMessage() string {
	return "Tmporaty Employee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

type PrintInfo interface {
	getMessage() string
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.age = 12
	ftEmployee.id = 8
	ftEmployee.name = "Juan"

	fmt.Println(ftEmployee)
	tEmployee := TemporatyEmployee{}
	getMessage(ftEmployee)
	getMessage(tEmployee)
}
