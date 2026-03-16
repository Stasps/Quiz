//[Спринт 2/13 → Тема 3/6: Массивы и слайсы → Урок 2/5: Слайсы] 
// Quiz 2

package main

import "fmt"

func main() {
	s := "Мезчхе%зцчхкьексцє%з%9%ьеце%ме%ширус%йусе"
	r := []rune(s)
	var msg string
	for _, v := range r {
		v -=5
		msg += string(v)
		}
	fmt.Println(msg)
}