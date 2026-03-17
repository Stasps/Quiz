//Спринт 4 → Тема 1: Работа со временем → Урок 2: Работа со временем
//https://practicum.yandex.ru/learn/go-developer-basic/courses/f8764346-5c9c-4d72-9e02-b5527dd2d5bf/sprints/835616/topics/023ca14e-81e0-41c6-ab6a-364ecd747ffc/lessons/ec09dbb8-9994-41c2-9d2c-a76c4721e314/#c4786592-af41-481b-9186-3ab86e0d420a
//Quiz 2

package main

import (
    "fmt"
    "time"
)

func main() {
    // создайте переменные типа time.Time
	var go120 = time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC)
	var go121 = time.Date(2023, time.August, 8, 0, 0, 0, 0, time.UTC)
    // получите интервал между датами
	sub := go121.Sub(go120)
	// вычислите количество дней
	days := sub.Hours()/24

	fmt.Printf("Между выходами версий прошло %d дней", int(days))
} 