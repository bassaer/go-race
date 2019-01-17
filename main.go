package main

import (
	"fmt"
)

func main() {
	value := 0

	go func() {
		value++
	}()

	if value == 0 {
		fmt.Printf("value is %#v\n", value)
	}
}
