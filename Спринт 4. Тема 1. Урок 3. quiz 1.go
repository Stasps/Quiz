//Спринт 4 → Тема 1: Работа со временем → Урок 3: Стандарт времени UTC
//https://practicum.yandex.ru/learn/go-developer-basic/courses/f8764346-5c9c-4d72-9e02-b5527dd2d5bf/sprints/835616/topics/023ca14e-81e0-41c6-ab6a-364ecd747ffc/lessons/45b3b967-f7c1-4bac-8cd1-3afa31f5c82a/#99c11c38-2dd1-4c9a-80c3-a4149d981e27
//Quiz 1

package main

import (
    "fmt"
    "time"
)

func whatTime(city string) time.Time {
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
    // напишите код, который берёт текущее UTC-время
    // и добавляет к нему сдвиг города в часах
	utc := time.Now().UTC()
	return utc.Add(time.Duration(offsetUTC[city]) * time.Hour)

}

func main() {
    for _, city := range []string{"Екатеринбург", "Москва", "Хабаровск"} {
        t := whatTime(city)
        fmt.Printf("%s: %d ч.\n", city, t.Hour())
    }
}