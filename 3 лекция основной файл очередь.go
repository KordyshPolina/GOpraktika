
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "yourproject/queue" 
)

func main() {
    q := queue.NewQueue(5)
     reader := bufio.NewReader(os.Stdin)

 for {
  fmt.Println("Введите команду (push, pop, exit):")
  input, _ := reader.ReadString('\n')
  input = strings.TrimSpace(input)

  switch input {
  case "push":
   fmt.Println("Введите значение для добавления в очередь:")
   valueInput, _ := reader.ReadString('\n')
   valueInput = strings.TrimSpace(valueInput)
   value, err := strconv.Atoi(valueInput)
   if err != nil {
    fmt.Println("Ошибка: введите целочисленное значение.")
    continue
   }
   push(q, value)
   fmt.Println(value, "добавлено в очередь.")
  case "pop":
   popped := pop(q)
   if popped != nil {
    fmt.Println("Извлечено из очереди:", popped)
   } else {
    fmt.Println("Очередь пуста.")
   }
  case "exit":
   fmt.Println("Программа завершена.")
   return
  default:
   fmt.Println("Неверная команда.")
  }
 }
}