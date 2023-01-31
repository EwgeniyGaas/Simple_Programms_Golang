package main

import "fmt"

func main() {
	var (
		num1 float64
		num2 float64
		sign string
		result float64
		buffer_rest_str string // пожирает остаток строки, в том числе перенос строки
	)
	for true {
		fmt.Print("\nВведите выражение: ")
		fmt.Scanf("%v%1s%v%v", &num1, &sign, &num2, buffer_rest_str)
		
		switch sign {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "%":
				if num2 == 0{
					fmt.Println("Ошибка. На ноль делить нельзя.")
					continue
				}
				buf1 := float64(int(num1))
				buf2 := float64(int(num2))
				if buf1 == num1 && buf2 == num2 {
					result = float64(int(num1) % int(num2))
				} else {
					fmt.Println("Ошибка. Нахождение остатка доступно только для целых чисел.")
					continue
				}
			case "/":
				if num2 != 0 {
					result = num1 / num2
				} else {
					fmt.Println("Ошибка. На ноль делить нельзя.")
					continue
				}
			default:
				fmt.Println("Ошибка. Операция невозможна.")
				continue
		}
		fmt.Println("Ответ:", result)
	}
}