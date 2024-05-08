package main

import (
  "fmt"
  "math/rand"
)

func main() {
  for x := 0; x <= 15; x++ {
    r1 := rand.Intn(10)
    r2 := rand.Intn(10)
    fmt.Printf("%v : r1: %v, r2: %v\n", x, r1, r2)
    if r1 < 4 && r2 < 4 {
      fmt.Println("both are less than 4")
    } else if r1 > 6 && r2 > 6 {
      fmt.Println("both are greater than 6")
    } else if r1 >= 4 && r2 <= 6 {
      fmt.Println("r1 is between 4 and 6, inclusive")
    } else if r2 != 5 {
      fmt.Println("r2 is not 5")
    }
    else {
      fmt.Println("no match made")
    }
  }
}
