//Спринт 4 → Тема 1: Работа со временем → Урок 3: Стандарт времени UTC
//https://practicum.yandex.ru/learn/go-developer-basic/courses/f8764346-5c9c-4d72-9e02-b5527dd2d5bf/sprints/835616/topics/023ca14e-81e0-41c6-ab6a-364ecd747ffc/lessons/45b3b967-f7c1-4bac-8cd1-3afa31f5c82a/#942c7b97-8f17-4383-9961-32ecb6f4e110
//Quiz 2

//Спринт 4 → Тема 1: Работа со временем → Урок 3: Стандарт времени UTC
//https://practicum.yandex.ru/learn/go-developer-basic/courses/f8764346-5c9c-4d72-9e02-b5527dd2d5bf/sprints/835616/topics/023ca14e-81e0-41c6-ab6a-364ecd747ffc/lessons/45b3b967-f7c1-4bac-8cd1-3afa31f5c82a/#942c7b97-8f17-4383-9961-32ecb6f4e110
//Quiz 2

package main

import (
    "fmt"
    "time"
)

func whatTime(friend string) time.Time {
    var friends = map[string]string{
        "Серёга": "Омск",
        "Соня": "Москва",
        "Дима": "Екатеринбург",
        "Алина": "Новосибирск",
        "Егор": "Калининград",
    }
    var offsetUTC = map[string]int{
        "Санкт-Петербург": 3,
        "Москва":          3,
        "Самара":          4,
        "Новосибирск":     7,
        "Екатеринбург":    5,
        "Казань":          3,
        "Омск":            6,
        "Хабаровск":       10,
        "Пермь":           5,
        "Краснодар":       3,
        "Калининград":     2,
    }

	// напишите недостающий код
	utc := time.Now().UTC()
    // Проверяем существование ключа
    if _, ok := friends[friend]; ok {
        return utc.Add(time.Duration(offsetUTC[friends[friend]]) * time.Hour) 
    }
	
	// Если друга нет в списке, возвращаем UTC
    return utc
}

func main() {
    for _, friend := range []string{"Соня", "Дима", "Егор"} {
        t := whatTime(friend)
        fmt.Printf("%s: %d ч.\n", friend, t.Hour())
    }
}