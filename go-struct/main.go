package main

import "fmt"

func main() {

	// 	var s struct {
	// 		name string
	// 		id   int
	// 	}
	// 	fmt.Println(s)
	// 	s.name = "Mohan"    //Assign value to field
	// 	fmt.Println(s.name) //Query value from field
	//

	type myStruct struct { //custom type based on struct
		name string
		id   int
	}
	var s myStruct
	fmt.Println((s))

	//struct literal
	s = myStruct{
		name: "Aishu",
		id:   2,
	}
	//structs are copied by value and we dont have data sharing.
	s2 := s
	s.name = "Mohan"
	fmt.Println(s, s2)

}
