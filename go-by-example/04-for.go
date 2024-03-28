// https://gobyexample.com/for

package main

import "fmt"

func main() {

	// basic iteration
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// one line iteration
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// break loop
	for {
		fmt.Println("loop")
		break
	}

	// pass within loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
