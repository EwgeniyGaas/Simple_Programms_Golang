/*
Проверяет является ли строка введённая пользователем палиндромом
А роза упала на лапу Азора
*/


package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func palindrom() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n') // считает всю строку из ввода
	text = strings.TrimSpace(text) // удалит пробелы и непечатные знаки (\r\n) в начале и конце строки
	text = strings.ToLower(text) // к нижнему регистру
	text = strings.Replace(text, " ", "", -1) // удалит пробелы между строк
	characters := strings.Split(text, "") // разобьёт по отдельным буквам
	size := len(characters)
	for i := 0; i < size/2; i++ {
		if characters[i] == characters[size-i-1] {
			continue
		} else {
			fmt.Println("НЕТ")
			return
		}
	}
	fmt.Println("Палиндром")
}

func main() {
	for true {
		palindrom()
	}
}