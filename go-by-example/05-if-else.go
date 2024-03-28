// https://gobyexample.com/if-else

package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "is less than 10")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// not on site

	for n := 0; n <= 20; n++ {
		if n < 5 {
			fmt.Println(n, "is less than 5")
		} else if n%5 == 0 {
			fmt.Println(n, "is divisible by 5")
		} else {
			fmt.Println(n, "is greater than 5 and not divisible by 5")
		}
	}
}
