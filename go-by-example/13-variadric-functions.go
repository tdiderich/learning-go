// https://gobyexample.com/variadic-functions

package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println(nums)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	a := sum(1, 2, 3, 4)
	fmt.Println(a)

	b := sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(b)

	n := []int{7, 8, 9}
	// must use slice to pass existing list
	c := sum(n...)
	fmt.Println(c)
}
