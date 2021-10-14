package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old \n", p.name, p.age)
}

func NewPerson(name string, age int) *person {
	return &person{name, age}
}

func main() {
	p := NewPerson("James", 23)
	p.SayHello()
}
