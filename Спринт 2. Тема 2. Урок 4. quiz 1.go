package main

import "fmt"

func main() {
	a := 3
	for a < 10 {
		b := 3
		fmt.Printf("%dx%d = %d\n", a, b, a*b)
		b +=2
		fmt.Printf("%dx%d = %d\n", a, b, a*b)
		b +=2
		fmt.Printf("%dx%d = %d\n", a, b, a*b)
		b +=2
		fmt.Printf("%dx%d = %d\n", a, b, a*b)
		a +=2
	}
}