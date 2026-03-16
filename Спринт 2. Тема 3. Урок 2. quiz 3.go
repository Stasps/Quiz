//[Спринт 2/13 → Тема 3/6: Массивы и слайсы → Урок 2/5: Слайсы] 
// Quiz 3

package main

import "fmt"

func main() {
    // оценки по отдельным ученикам
    marks := [][]int{
        {5, 4, 5, 5},
        {3, 4, 4, 5, 3},
        {2, 3, 3},
        {5, 5, 4},
        {4, 3, 4, 4, 3},
    }
    var average float32 // итоговый средний балл

	// добавьте код для подсчёта среднего балла
	var sum int
	var count int

	for i := 0; i < 5; i++ {
		for _, v := range marks[i] {
//			fmt.Printf("%d ", v) // для проверки
			sum += v // суммируем оценки
			count++ // считаем количество оценок
			}
//		fmt.Println()
	}
	


	average = float32(sum) / float32(count)

    fmt.Printf("%.2f", average)
}