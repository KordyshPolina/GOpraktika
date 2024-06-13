package main

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
 "strings"
)

type queue struct {
 s    []interface{} // слайс в котором хранятся значения
 low  int           // индекс нижней границы очереди
 high int           // индекс верхней границы очереди
 size int           // размер очереди
}

func newQueue(size int) *queue {
 return &queue{
  s:    make([]interface{}, size),
  size: size,
  low:  -1,
  high: -1,
 }
}

// push - добавление в очередь значения
func push(q *queue, v interface{}) {
 if q.high == q.size-1 {
  return
 }
 q.high++
 q.s[q.high] = v
 if q.low == -1 {
  q.low = 0
 }
}

// pop - получение значения из очереди и его удаление
func pop(q *queue) interface{} {
 if q.low == -1 || q.low > q.high {
  return nil
 }
 v := q.s[q.low]
 q.low++
 return v
}

func main() {
 q := newQueue(5)
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
