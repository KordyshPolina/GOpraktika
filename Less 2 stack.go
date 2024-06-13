package main

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
 "strings"
)

type stack struct {
 s    []interface{}
 head int
}

func newStack(size int) *stack {
 return &stack{
  s:    make([]interface{}, size),
  head: -1,
 }
}

func push(s *stack, v interface{}) {
 s.head++
 s.s[s.head] = v
}

func pop(s *stack) interface{} {
 if s.head < 0 {
  return nil
 }
 v := s.s[s.head]
 s.head--
 return v
}

func peek(s *stack) interface{} {
 if s.head < 0 {
  return nil
 }
 return s.s[s.head]
}

func main() {
 s := newStack(5)
 reader := bufio.NewReader(os.Stdin)

 for {
  fmt.Println("Введите команду (push, pop, peek, exit):")
  input, _ := reader.ReadString('\n')
  input = strings.TrimSpace(input)

  switch input {
  case "push":
   fmt.Println("Введите значение для добавления в стек:")
   valueInput, _ := reader.ReadString('\n')
   valueInput = strings.TrimSpace(valueInput)
   value, err := strconv.Atoi(valueInput)
   if err != nil {
    fmt.Println("Ошибка: введите целочисленное значение.")
    continue
   }
   push(s, value)
   fmt.Println(value, "добавлено в стек.")
  case "pop":
   popped := pop(s)
   if popped != nil {
    fmt.Println("Извлечено из стека:", popped)
   } else {
    fmt.Println("Стек пуст.")
   }
  case "peek":
   top := peek(s)
   if top != nil {
    fmt.Println("Значение на вершине стека:", top)
   } else {
    fmt.Println("Стек пуст.")
   }
  case "exit":
   fmt.Println("Программа завершена.")
   return
  default:
   fmt.Println("Неверная команда.")
  }
 }
}
