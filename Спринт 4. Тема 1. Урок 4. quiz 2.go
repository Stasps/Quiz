//Спринт 4 → Тема 1: Работа со временем → Урок 4: Форматирование и парсинг времени
//https://practicum.yandex.ru/learn/go-developer-basic/courses/f8764346-5c9c-4d72-9e02-b5527dd2d5bf/sprints/835616/topics/023ca14e-81e0-41c6-ab6a-364ecd747ffc/lessons/382415d8-bbea-405a-96eb-ddc1c2e25b03/#d395f498-038d-4524-9c87-620392976763
//Quiz 2

package main

import (
    "fmt"
	"time"
)

// формат дней рождений
const layout = "02.01.2006"

func Age(birthday string) (int, error) {
	// напишите код функции
	birthdayTime, err := time.Parse(layout, birthday)
	if err != nil {
		return 0, err
	}
	now := time.Now()
	age := now.Year() - birthdayTime.Year()
	if now.Month() < birthdayTime.Month() {
		age--	
	}
	if now.Month() == birthdayTime.Month() {
		if now.Day() < birthdayTime.Day() {
			age--
		}
	}
	if age < 0 {
		return 0, fmt.Errorf("возраст не может быть отрицательным")
	}
	return age, nil
}

func main() {
    for _, v := range []string{"04.01.1969", "28.07.1995",
        "16.12.2007", "07.03.1947"} {
        age, err := Age(v)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Println(v, ":", age)
    }
} 