// quiz 3

package main

import "fmt"

func main() {
    count := 2  // количество бактерий
    day := 1    // счётчик дней
	
	fmt.Printf("day\tmol start\tdied\tadd\ttotal\n") //день, мотекул в начале, умерло, добавилось, всего
	for day <= 30 {
		fmt.Printf("%d\t %d\t", day, count) //день, молекул в начале дня
		if day >= 11 {
				fmt.Printf("\t%d\t", count/10)
				count -= count/10
		} else {fmt.Printf("\t0\t")}
		{fmt.Printf("%d\t", count/2)}
		count += count/2
fmt.Printf("%d\n", count)
		day++
	}

	fmt.Println("Количество бактерий через 30 дней:", count)
}