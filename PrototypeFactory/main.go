package main

import "fmt"

type Employee struct {
	Name         string
	Position     string
	AnnualIncome int
}

const (
	Developer = iota
	Manager
	Tester
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{Position: "developer"}
	case Manager:
		return &Employee{Position: "manager"}
	case Tester:
		return &Employee{Position: "tester"}
	default:
		return &Employee{Position: "nil"}
	}
}

func main() {
	em := NewEmployee(1)
	em.Name = "Cuong"
	em.AnnualIncome = 10000
	fmt.Println(em)
}
