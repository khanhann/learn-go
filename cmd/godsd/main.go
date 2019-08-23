package main

import (
	"fmt"
	"sync"

	dsd "github.com/duynhanf/learn-go"
)

func main() {
	fmt.Println("godsd started")
	dsd.NewPerson("Nhan", 27).GoSchool()

	fmt.Println("godsd ended")

	var locker sync.Locker
	locker.Lock()
	cond := sync.NewCond(locker)

	fmt.Println(cond)
}
