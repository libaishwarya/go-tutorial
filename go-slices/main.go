package main

import "fmt"

func main() {
	var s []string //declaring empty slice
	fmt.Println(s)

	s = []string{"Coffee", "Tea", "Juice"} //initializing slice
	fmt.Println(s)

	fmt.Println(s[1])
	s[1] = "Chai-Tea" //updating the slice value
	fmt.Println(s[1])

	s = append(s, "hot-chocolate", "hot-tea")
	fmt.Println(s)

	s1 := s
	s1[2] = "chai-latte"

	fmt.Println(s, s1) //index 2 in s is updated as well indicating slices are sharing the same data
	//----------------------------
	s2 := make([]int, 5, 7) //declaring using make function
	fmt.Println(len(s2), cap(s2))
	s2 = []int{1, 2, 3, 4, 5}
	fmt.Println(s2)
	//Traversing through slice
	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}
	for key, value := range s2 {
		fmt.Println(key, value)
	}

	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8} //string literal
	fmt.Printf("Address of sliceA is : %p\n", sliceA)
	fmt.Println(sliceA, len(sliceA), cap(sliceA))

	sliceB := sliceA[:5]
	fmt.Printf("Address of sliceB is : %p\n", sliceB)
	fmt.Println(sliceB, len(sliceB), cap(sliceB))

	sliceB = append(sliceB, 21, 43, 54, 42, 325, 23) //The address of sliceB changed after appending
	//because a new underlying array was allocated due to exceeding the original capacity.
	fmt.Printf("Address of sliceB is : %p\n", sliceB)
	fmt.Println(sliceB, len(sliceB), cap(sliceB))

	fmt.Println("----------------")

	slice := make([]int, 0, 3) // Capacity of 3
	fmt.Println("Initial Slice:", slice)
	fmt.Printf("Address of slice is : %p\n", slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// Appending elements beyond capacity

	fmt.Println("----------------")

	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println("Modified Slice:", slice)
	fmt.Printf("Address of slice is : %p\n", slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))
}
