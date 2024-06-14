package main

import "fmt"

type Employee struct {
    Name     string // имя
    Age      int    // возраст
    Position string // позиция
    Salary   int    // зарплата
}

var commands = `
1 - Добавление нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func main() {
    const size = 512
    empls := make([]*Employee, size)
    var count int

    for {
        cmd := 0
        fmt.Print(commands)
        fmt.Scanf("%d", &cmd)

        switch cmd {
        case 1:
            // Добавляем нового сотрудника
            if count < size {
                empl := &Employee{}
                fmt.Println("\nИмя:")
                fmt.Scanf("%s", &empl.Name)
                fmt.Println("Возраст:")
                fmt.Scanf("%d", &empl.Age)
                fmt.Println("Позиция:")
                fmt.Scanf("%s", &empl.Position)
                fmt.Println("Зарплата:")
                fmt.Scanf("%d", &empl.Salary)
                empls[count] = empl
                count++
            } else {
                fmt.Println("Достигнуто максимальное количество сотрудников")
            }
        case 2:
            // Удаляем сотрудника
            var index int
            fmt.Println("Введите номер сотрудника для удаления:")
            fmt.Scanf("%d", &index)
            if index >= 1 && index <= count {
                copy(empls[index-1:], empls[index:])
                empls[count-1] = nil
                count--
            } else {
                fmt.Println("Неверный номер сотрудника")
            }
        case 3:
            // Вывод списка сотрудников
            fmt.Println("Список сотрудников:")
            for i := 0; i < count; i++ {
                fmt.Printf("Сотрудник %d: Имя: %s, Возраст: %d, Позиция: %s, Зарплата: %d\n", i+1, empls[i].Name, empls[i].Age, empls[i].Position, empls[i].Salary)
            }
        case 4:
            fmt.Println("Выход из программы")
            return
        }
    }
}
