package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("fmt")

	ctx := context.WithValue(context.Background(), "abc", 123)

	fmt.Println(ctx.Value("abc"))
}
