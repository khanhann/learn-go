package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	name string
}

func main() {
	deadline := time.Now().Add(150 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	ch := make(chan data, 1)
	go func() {
		time.Sleep(50 * time.Millisecond)

		ch <- data{"123"}
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}
