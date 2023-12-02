package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}
func (e *Employee) SetName(name string) {
	e.name = name
}
func (e *Employee) GetId() int {
	return e.id
}
func (e *Employee) GetName() string {
	return e.name
}
func (e Employee) String() string {
	return fmt.Sprintf("%v %v", e.id, e.name)
}
func main() {
	e := Employee{
		id: 5, name: "juan",
	}
	fmt.Println(e)

	e.SetId(55)
	fmt.Println(e)

	e.SetName("diego")
	fmt.Println(e)

	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
