package main

import (
 "fmt"
 "strings"
)

var romanNumeralMap = map[string]int{
 "I": 1,
 "V": 5,
 "X": 10,
 "L": 50,
 "C": 100,
 "D": 500,
 "M": 1000,
}

func romanToArabic(roman string) int {
 result := 0
 prevValue := 0

 for i := len(roman) - 1; i >= 0; i-- {
  value := romanNumeralMap[string(roman[i])]
  if value < prevValue {
   result -= value
  } else {
   result += value
  }
  prevValue = value
 }

 return result
}

func main() {
 var input string
 for {
  fmt.Print("Введите римское число (I - MMM): ")
  fmt.Scanln(&input)
  
  // Проверка на пустой ввод
  if input == "" {
   fmt.Println("Пожалуйста, введите римское число.")
   continue
  }

  arabic := romanToArabic(strings.ToUpper(input))
  if arabic >= 1 && arabic <= 3000 {
   fmt.Printf("Арабское представление числа %s: %d\n", input, arabic)
  } else {
   fmt.Println("Пожалуйста, введите римское число от I до MMM.")
  }
 }
}

