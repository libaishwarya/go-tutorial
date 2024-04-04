package main

import "fmt"

func main() {
	dividend, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	dividend, divisor = 10, 0
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))
}

func divide(divided, divisor int) int {
	defer func() {
		if message := recover(); message != nil {
			fmt.Println(message)
		}
	}()
	return divided / divisor
}
