package main

import (
	"fmt"

	"github.com/hagan/puppy"
)

func main() {
	puppy.From13()

	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	// also like this
	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())
}
