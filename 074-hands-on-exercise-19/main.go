package main

import (
  "fmt"
  "math/rand"
)

func main() {
  x := rand.Intn(400)
  fmt.Println("The value of x is %v\t", x)
  if x <= 100 {
    fmt.Println("is less than 100")
  } else if x >= 101 && x <= 200 {
    fmt.Println("is between 101 and 200")
  } else if x >= 201 && x <= 250 {
    fmt.Println("this was between 201 and 250")
  } else {
    fmt.Println("this was more than 250")
  }

  fmt.Println(rand.Intn(3))
  fmt.Println(rand.Intn(3))
  fmt.Println(rand.Intn(3))
  fmt.Println(rand.Intn(3))
  fmt.Println(rand.Intn(3))
}
