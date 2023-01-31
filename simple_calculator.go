/*
Простой калькулятор. Операция с 2-мя числами. Целыми и дробными.
Сложение, вычитание, умножение, деление, остаток от деления.
Обрабатывает выражения вида:  1+2  2-1  5*3  3/5  7%4
*/

package main

import "fmt"

func main() {
	var (
		num1 float64
		num2 float64
		sign string
		result float64
		rest_str string // пожирает остаток строки, что-бы цикл не крутил пустые итерации
	)
	for true {
		num1, num2, sign, result = 0, 0, "", 0
		fmt.Print("\nВведите выражение: ")
		fmt.Scanf("%v%1s%v%v", &num1, &sign, &num2, &rest_str)
		
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
				buf1 := float64(int(num1)) // отбрасываем дробную часть, для проверки на целое число
				buf2 := float64(int(num2))
				if buf1 == num1 && buf2 == num2 { // проверка на целое число
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