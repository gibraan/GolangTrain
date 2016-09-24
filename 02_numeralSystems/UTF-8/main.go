package main

import "fmt"

func main() {
	for i := 200; i < 300; i++ {
		fmt.Printf("%d - %b - %#X %q \n", i, i, i, i)
	}
}
