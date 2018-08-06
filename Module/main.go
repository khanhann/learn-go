package main

import (
	"fmt"
	"github.com/khanhann/LearnGo/Module/private"
)

func main() {
	fmt.Println("Module test")

	p := private.New("Yamamoto","Ishihara")
	p.Bar()
}