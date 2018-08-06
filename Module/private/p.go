package private

import (
	"fmt"
)

type P struct {
	firstName string
	lastName string
}

func New (firstName string, lastName string) *P {
	return &P{
		firstName:   firstName,
		lastName: lastName,
	}
}

func (p *P) Bar() {
	fmt.Println("Fullname is " + p.firstName + " " + p.lastName)
}