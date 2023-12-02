package main

import (
	"fmt"
	"time"
)

func main() {
	times := 100000000

	// first loop
	t1 := time.Now()
	for i := 0; i < times; i++ {
		// do something
	}
	d1 := fmt.Sprint(time.Since(t1).Milliseconds()) + "ms"

	// second loop
	t2 := time.Now()
	for i := 0; i < times; i++ {
		// do something else
	}
	d2 := fmt.Sprint(time.Since(t2).Milliseconds()) + "ms"

	fmt.Println("first loop: ", d1)
	fmt.Println("second loop: ", d2)
}
