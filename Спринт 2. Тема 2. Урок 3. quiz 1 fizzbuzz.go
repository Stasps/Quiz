package main

import "fmt"

func main() {

	for i:=1 ; i <= 100 ; i++ {
		switch i%3 == 0 {	
			case true:
			fmt.Printf ("Fizz")
		}
		switch i%5 == 0 {	
			case true:
			fmt.Printf ("Buzz")
		}

		if (i%3 != 0 ) && (i%5 != 0) { fmt.Printf ("%d\n", i)
		} else {
			fmt.Printf ("\n")
}
}
}