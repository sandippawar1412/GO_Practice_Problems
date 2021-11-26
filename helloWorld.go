package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	i := 0
	for i < 100 {
		if i%2 == 0 {
			fmt.Println("Even ", i)
		} else {
			fmt.Println("Odd ", i)
		}
		i += 1
	}
}
