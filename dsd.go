package dsd

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name,
		age,
	}
}

func (p *Person) GoSchool() {
	fmt.Printf("Today, %s %d years old go to school", p.name, p.age)

}
