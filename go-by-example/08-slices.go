// https://gobyexample.com/slices

package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[:5]
	fmt.Println("sl1:", l)

	l = s[2:]
	fmt.Println("sl2:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declare:", t)

	twoD := make([][]int, 3)
	fmt.Println("declare:", twoD)

	for i := 0; i <= 2; i++ {
		twoD[i] = make([]int, 3)
		for j := 0; j <= 2; j++ {
			twoD[i][j] = i * j
		}
	}

	fmt.Println("set:", twoD)
}
