package main

import "fmt"

func main() {
	var arr = [2]int{1, 2} //with defined length
	arr2 := [3]int{4, 5, 6}
	fmt.Println(arr)
	fmt.Println(arr2)

	var arr3 = [...]int{1, 2, 3, 4, 5} //with inferred length
	arr4 := [2]int{3, 2}
	fmt.Println(arr3)
	fmt.Println(arr4)

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"} //array of string
	fmt.Println(cars)

	//Traversing through array
	fmt.Println("Array items are: ")
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

}
