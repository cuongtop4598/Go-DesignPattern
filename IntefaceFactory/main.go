package main

import "fmt"

type Employee struct {
	Name     string
	Position string
	Salary   int
}

func NewEmployeeFactory(name, position string) func(salary int) *Employee {
	return func(salary int) *Employee {
		return &Employee{Name: name, Position: position, Salary: salary}
	}
}

func main() {
	dev := NewEmployeeFactory("cuong", "dev")
	dev1 := dev(1000)
	fmt.Println(dev1)
}
