package main

import "fmt"

func main() {
	for i := 1000000; i < 1000200; i++ {
		fmt.Printf("%d - %b - %#X \n", i, i, i)
	}
}
