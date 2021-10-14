package main

import "fmt"

type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

type PersonBuilder struct {
	person *Person // needs to be inited
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (it *PersonBuilder) Build() *Person {
	return it.person
}

func (it *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*it}
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(company string) *PersonJobBuilder {
	pjb.person.CompanyName = company
	return pjb
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (pab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	pab.person.Postcode = postcode
	return pab
}
func main() {
	pb := NewPersonBuilder()
	pb.
		Works().
		At("MTA")
	person := pb.Build()
	fmt.Println(person)
}
