package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Hello, playground")

	pBirth := time.Now()

	pBirth2 := pBirth

	pBirth3 := pBirth.Add(time.Hour * 9)

	fmt.Println(pBirth)
	fmt.Println(pBirth2)
	fmt.Println(pBirth3)

	fmt.Println(math.Abs(pBirth.Sub(pBirth3).Hours()))
}
