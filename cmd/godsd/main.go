package main

import (
	"fmt"

	dsd "github.com/duynhanf/learn-go"
)

func main() {
	fmt.Println("godsd started")
	dsd.NewPerson("Nhan", 27).GoSchool()

	fmt.Println("godsd ended")
}
