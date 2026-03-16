package main

import "fmt"



 main() {
    // рекомендованные фильмы
    recommended := [...]string{"Хатико", "23", "Достучаться до небес",
        "Хакеры", "Трон", "1408"}
    // коллекция
    collection := [...]string{"Трон", "Военные игры", "Тихушники",
        "Джонни Мнемоник", "Хакеры", "Нирвана", "23", "Враг государства",
        "Взлом", "Пароль рыба-меч", "Сеть", "Кто я"}

	// допишите программу
	lenC := len(collection)

	for _, v := range recommended {
		for t := 0; t < lenC; t++ {
			if v == collection[t] { fmt.Println(collection[t]) }
		}
	}
}