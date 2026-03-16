//Спринт 2 → Тема 3: Массивы и слайсы → Урок 2: Слайсы
//q1

package main

import "fmt"

func main() {
    list := []int{1, 5, 4, 8, 7, 6, 0, 2, 9}
    var (
        index  int
        isLoop bool  // должна быть true, если есть петля
    )
	// добавьте в код поиск петли
	var count int
    for index >= 0 && index < len(list) {
		index = list[index]
		count++
		if count >= len(list) {break}
    }
	isLoop = count >= len(list) 
    fmt.Println("В последовательности имеется петля?", isLoop)
}