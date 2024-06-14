package main

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
 "strings"
)

type singlyLinkedList struct {
    first *item
    last  *item
    size  int
}

type item struct {
    v    interface{}
    next *item
}

func newSinglyLinkedList() *singlyLinkedList {
    return &singlyLinkedList{}
}

// add - добавление значения в связный список
func add(l *singlyLinkedList, v interface{}) {
    newItem := &item{v: v}
    if l.first == nil {
        l.first = newItem
        l.last = newItem
    } else {
        l.last.next = newItem
        l.last = newItem
    }
    l.size++
}

// get - получение значения по индексу из связанного списка
func get(l *singlyLinkedList, idx int) interface{} {
    if idx < 0 || idx >= l.size {
        return nil
    }
    curr := l.first
    for i := 0; i < idx; i++ {
        curr = curr.next
    }
    return curr.v
}

// remove - удаление значения по индексу из списка
func remove(l *singlyLinkedList, idx int) {
    if idx < 0 || idx >= l.size {
        return
    }
    if idx == 0 {
        l.first = l.first.next
        l.size--
        return
    }
    prev := l.first
    for i := 0; i < idx-1; i++ {
        prev = prev.next
    }
    prev.next = prev.next.next
    if idx == l.size-1 {
        l.last = prev
    }
    l.size--
}

// values - получение слайса значений из списка
func values(l *singlyLinkedList) []interface{} {
    result := make([]interface{}, l.size)
    curr := l.first
    for i := 0; i < l.size; i++ {
        result[i] = curr.v
        curr = curr.next
    }
    return result
}


func main() {
 list := newSinglyLinkedList()
 reader := bufio.NewReader(os.Stdin)

 for {
  fmt.Println("Введите команду (add, get, remove, values, exit):")
  input, _ := reader.ReadString('\n')
  input = strings.TrimSpace(input)

  switch input {
  case "add":
   fmt.Println("Введите значение для добавления в список:")
   valueInput, _ := reader.ReadString('\n')
   valueInput = strings.TrimSpace(valueInput)
   add(list, valueInput)
   fmt.Println(valueInput, "добавлено в список.")
  case "get":
   fmt.Println("Введите индекс для получения значения из списка:")
   idxInput, _ := reader.ReadString('\n')
   idxInput = strings.TrimSpace(idxInput)
   idx, err := strconv.Atoi(idxInput)
   if err != nil {
    fmt.Println("Ошибка: введите целочисленное значение.")
    continue
   }
   result := get(list, idx)
   if result != nil {
    fmt.Println("Значение по индексу", idx, ":", result)
   } else {
    fmt.Println("Индекс вне диапазона.")
   }
  case "remove":
   fmt.Println("Введите индекс для удаления значения из списка:")
   idxInput, _ := reader.ReadString('\n')
   idxInput = strings.TrimSpace(idxInput)
   idx, err := strconv.Atoi(idxInput)
   if err != nil {
    fmt.Println("Ошибка: введите целочисленное значение.")
    continue
   }
   remove(list, idx)
   fmt.Println("Значение с индексом", idx, "удалено из списка.")
  case "values":
   vals := values(list)
   fmt.Println("Значения в списке:", vals)
  case "exit":
   fmt.Println("Программа завершена.")
   return
  default:
   fmt.Println("Неверная команда.")
  }
 }
}
