// https://gobyexample.com/pointers

package main

import "fmt"

// doesn't actually change the value passed
func zeroval(ival int) {
	ival = 0
}

// actually updates the value in memory since it's passed the memory address
func zerovalpointer(pointer *int) {
	*pointer = 0
}

func main() {
	i := 1

	fmt.Println("init:", i)
	zeroval(i)
	fmt.Println("val:", i)
	zerovalpointer(&i)
	fmt.Println("pointer:", i)
}
