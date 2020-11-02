package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println(i)

	zeroval(i)
	fmt.Println("zero val: ", i)

	zeroptr(&i)
	fmt.Println("zero ptr: ", i)

	fmt.Println("zero ptr: ", &i)
}
