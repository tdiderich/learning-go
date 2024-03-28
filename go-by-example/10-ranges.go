// https://gobyexample.com/range

package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for index, num := range nums {
		fmt.Println("index:", index)
		sum += num
		fmt.Println("updated sum:", sum)
	}
	fmt.Println("final sum:", sum)

	for i, n := range nums {
		if n == 3 {
			fmt.Println("n is 3 at index", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	// key, value
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}

	// key
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
