//Спринт 2 → Тема 5: Мапы → Урок 1: Знакомство с мапами
//q3

package main

import "fmt"

func main() {
    // оценки по отдельным ученикам
    marks := map[string][]int{
        "Светлана": {5, 4, 5, 5, 4},
        "Артём": {3, 4, 4, 5, 3},
        "Александр": {2, 3, 3, 4},
        "Ольга": {5, 5, 4, 4},
        "Мария": {4, 3, 4, 4, 3, 5},
    }
    student := "Светлана"
    var average float32 // средний балл

    // допишите недостающий код
    // ...
	marks[student] = append(marks[student], 5)
	
	// fmt.Println(marks[student]) для проверки

	// добавьте код для подсчёта среднего балла
	var sum int
	var count int

		for _, v := range marks[student] {
			sum += v // суммируем оценки
			count++ // считаем количество оценок
			}

	average = float32(sum)/float32(count)

    fmt.Printf("%.2f\n", average)
}