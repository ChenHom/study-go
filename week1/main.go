package main

import (
	"fmt"
	concepts "week1/concepts/struct"
)

func main() {
	concepts.Embedded()

	fmt.Println("")

	r := concepts.Rectangle{Width: 10, Height: 5}
	fmt.Printf("Area is: %d\n", r.Area())

	r.Scale(10)
	fmt.Printf("Area is: %d\n", r.Area())
	fmt.Printf("Width: %d, Height: %d\n", r.Width, r.Height)

	fmt.Println("")
}
