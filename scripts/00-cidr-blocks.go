// creates a list of subnet ranges from 10.10.0.0/24 to 10.10.255.0/24

package main

import (
	"fmt"
)

func main() {
	for n := 0; n <= 255; n++ {
		fmt.Printf("10.10.%v.0/24\n", n)
	}
}
