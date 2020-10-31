package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Print("One")
		break
	case 2:
		fmt.Print("Two")
		break
	case 3:
		fmt.Print("Three")
	}

	fmt.Println()
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
		break
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatIAm := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's integer")
			break
		case bool:
			fmt.Println("It's boolean")
			break
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatIAm(1)
	whatIAm(true)
	whatIAm("hey")
}
