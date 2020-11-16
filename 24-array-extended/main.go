package main

import "fmt"

func arrPush(n int, arr []interface{}) []interface{} {
	newNumber := len(arr) + 1
	newArr := make([]interface{}, 0)
	if newNumber%3 == 0 && newNumber%5 == 0 {
		newArr = append(arr, "fizz-buzz")
	} else if newNumber%3 == 0 {
		newArr = append(arr, "fizz")
	} else if newNumber%5 == 0 {
		newArr = append(arr, "buzz")
	} else {
		newArr = append(arr, newNumber)
	}
	if len(arr) == n {
		return arr
	} else {
		return arrPush(n, newArr)
	}
}

func main() {
	arr := make([]interface{}, 0)
	arr2 := arrPush(15, arr)
	fmt.Println(arr2)
}
