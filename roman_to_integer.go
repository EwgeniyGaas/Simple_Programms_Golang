/*
Осуществляет преобразование римских чисел в арабские
Проверяет последовательность цифр в римском числе на коррекность
https://leetcode.com/problems/roman-to-integer/

string can contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M')
-------------------
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
-------------------
I can be placed before V (5) and X (10) to make 4 and 9. 
X can be placed before L (50) and C (100) to make 40 and 90. 
C can be placed before D (500) and M (1000) to make 400 and 900.
*/

package main

import (
	"fmt"
	"strings"
	)

// Проверяет что текущая римская цифра это или "I" или "X" или "C"	
func check(v int) bool {
	if v == 1 || v == 10 || v == 100 {
		return true
	} else {
		return false
	}
}

// Проверяет последовательность цифр в римском числе на коррекность
func isRomanValueCorrect(values []int, i int) bool {
	// если значение меньше, чем следующее в 5 или 10 раз, и это "I", "X" или "C" значит перед нами вариант типа "IV", "IX"
	if check(values[i]) == true && (values[i+1] / values[i]) == 10 || (values[i+1] / values[i]) == 5 {
		if (i-1) >= 0 { // есть ли впереди ещё цифры? проверить что не три цифры подряд на уменьшение
			if values[i-1] >= values[i] { // напр. число типа "IXL", (L / X == 10) - correct, но "I" - третья цифра подряд на уменьшение - incorrect
				return true
			} else {
				return false  // неправильная последовательность символов в римском числе
			}
		} else {
			return true
		}
	} else {
		return false // неправильная последовательность символов в римском числе
	}
}

// Преобразует римское число в арабские
func romanToInt(s string) int {
	var values []int
	var result int
	
	slice_str := strings.Split(s, "")
	
	for i := range(slice_str) { // преобразует римские цифры в соответствующие арабские числа
		switch slice_str[i] {
			case "I": values = append(values, 1)
			case "V": values = append(values, 5)
			case "X": values = append(values, 10)
			case "L": values = append(values, 50)
			case "C": values = append(values, 100)
			case "D": values = append(values, 500)
			case "M": values = append(values, 1000)
			default: return 0
		}
	}
	result = values[len(values)-1]  // последний элемент. Цикл ниже начнём с предпоследнего
	for i := len(values)-2; i >= 0; i-- {
		if values[i] < values[i+1] { // если предыдущая цифра < последующей, то проверяем число на корректность прежде чем отнимать
			if isRomanValueCorrect(values, i) == true {
				result -= values[i]
			} else {
				return 0  // неправильная последовательность символов в римском числе
			}
		} else {
			result += values[i] // если предыдущая цифра > или == последующей, то всё ок
		}
	}
	return result
}

func main() {
	fmt.Println("I =", romanToInt("I"))             // -> 1
	fmt.Println("III =", romanToInt("III"))         // -> 3
	fmt.Println("IX =", romanToInt("IX"))           // -> 9
	fmt.Println("X =", romanToInt("X"))             // -> 10
	fmt.Println("XI =", romanToInt("XI"))           // -> 11
	fmt.Println("XLVII =", romanToInt("XLVII"))     // -> 47
	fmt.Println("LIV =", romanToInt("LIV"))         // -> 54
	fmt.Println("IXL =", romanToInt("IXL"))         // -> 0  incorrect value
	fmt.Println("VL =", romanToInt("VL"))           // -> 0  incorrect value
	fmt.Println("XL =", romanToInt("XL"))           // -> 40
	fmt.Println("MCMXCIV =", romanToInt("MCMXCIV")) // -> 1994
	fmt.Println("MMXXIII =", romanToInt("MMXXIII")) // -> 2023
	fmt.Println("IIL =", romanToInt("IIL"))	        // -> 0  incorrect value
	fmt.Println("MMXXMII =", romanToInt("MMXXMII")) // -> 0  incorrect value
	fmt.Println("DM =", romanToInt("DM"))           // -> 0  incorrect value
	fmt.Println("MD =", romanToInt("MD"))           // -> 1500
}
