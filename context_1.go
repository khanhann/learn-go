package main

import (
	"context"
	"fmt"
)

type userKey int

func main() {
	fmt.Println("main")

	bg := context.Background()

	const uk userKey = 0

	ctx := context.WithValue(bg, uk, "abcdef")

	if s, ok := ctx.Value(uk).(string); ok {
		fmt.Println("User key: ", s)
	}

	if s, ok := ctx.Value(uk).(string); ok {
		fmt.Println("User not found", s)
	}

}
