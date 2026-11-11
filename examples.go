package main

import (
  "fmt"
)

func add(a int, b int) int {
  return a + b
}

func main() {
  var result int
  result = add(  1 ,2)
  fmt.Println("Result is: ", result )
}
