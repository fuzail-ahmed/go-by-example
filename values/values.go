package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(1.2)
	fmt.Println("String")
	fmt.Println(true)

	// String concatination
	fmt.Println("Happy" + " " + "Coding :)")

	// int
	fmt.Println("1+1 = ", 1+1)

	// Float
	fmt.Println("7.3/3.0 = ", 7.3/3.0)

	// Boolean
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
