//Ссылка на урок: [Спринт 2 → Тема 2 → Урок 3:

package main

import "fmt"

func main() {
    var hidden string // результирующая строка со звёздочками
    email := `vasyapupkin33@mail.ru`

	var unHide bool

	for i, v := range email {
		if string(v) == "@" { unHide = true }
		if (i <= 1) { hidden += string(v)
					} else {
		if unHide == false { hidden += "*"
						   } else { hidden += string(v) }
}
		i++
	}
	fmt.Println(hidden)
}