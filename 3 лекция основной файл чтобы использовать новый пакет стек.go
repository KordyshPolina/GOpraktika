package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    
    "yourproject/stack"
)

func main() {
    s := stack.NewStack(5)
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
            stack.Push(s, value)
            fmt.Println(value, "добавлено в стек.")
        case "pop":
            popped := stack.Pop(s)
            if popped != nil {
                fmt.Println("Извлечено из стека:", popped)
            } else {
                fmt.Println("Стек пуст.")
            }
        case "peek":
            top := stack.Peek(s)
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
