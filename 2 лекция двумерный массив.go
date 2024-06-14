package main

import (
 "fmt"
 "math/rand"
 "time"
)

func generateUniqueRandomNumbers(n, max int) []int {
 rand.Seed(time.Now().UnixNano())
 numbers := make([]int, n)
 used := make(map[int]bool)

 for i := 0; i < n; {
  num := rand.Intn(max)
  if !used[num] {
   numbers[i] = num
   used[num] = true
   i++
  }
 }

 return numbers
}

func main() {
 rows, cols := 5, 5
 maxNumber := rows * cols

 uniqueNumbers := generateUniqueRandomNumbers(rows*cols, maxNumber)

 matrix := make([][]int, rows)
 for i := range matrix {
  matrix[i] = make([]int, cols)
 }

 k := 0
 for i := 0; i < rows; i++ {
  for j := 0; j < cols; j++ {
   matrix[i][j] = uniqueNumbers[k]
   k++
  }
 }

 fmt.Println("Двумерный массив с уникальными случайными числами:")
 for _, row := range matrix {
  fmt.Println(row)
 }
}

