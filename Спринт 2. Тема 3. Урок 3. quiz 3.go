//[Спринт 2/13 → Тема 3/6: Массивы и слайсы → Урок 3/5: Функции для работы со слайсами]
// https://practicum.yandex.ru/learn/go-developer-basic/courses/f8764346-5c9c-4d72-9e02-b5527dd2d5bf/sprints/835614/topics/8f94bc33-ec99-4156-8851-819f0c83e6c0/lessons/e236aa72-c288-4845-aefc-56ffd38e6b3b/#b9896463-c0fa-4466-9bbf-246afcd38cb5
// Quiz 3

package main

import "fmt"

func main() {
    words := []string{"сервер", "cистема", "специалист", "слайc", "процессор",
        "масcив", "строка", "максимум", "cпоcоб", "парсер", "условие"}
    var mistakes []string
	
//	var wrong rune = 99 // руна 99 - латинская "c", а 63 = "?"
//	quest := rune('?')
//	fmt.Printf("rune %d = %c, а ? = %d\n",wrong,wrong,quest)

	for _, v := range words {
		found := false // Сбрасываем found для каждого нового слова
		wordRunes := []rune(v) // Преобразуем слово в срез рун для возможности замены
		for i, w := range wordRunes {
				if w == 99 { // руна 99 - латинская "c", а 63 = "?"
					found = true
					wordRunes[i] = 63
				}
			}
			if found == true {
				correctedWord := string(wordRunes) // Преобразуем измененный срез рун обратно в строку
				mistakes = append(mistakes, correctedWord)
				}
		}

    fmt.Println(mistakes)
}