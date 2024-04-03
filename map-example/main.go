package main

import "fmt"

func main() {
	var m map[string][]string
	fmt.Println(m)

	m = map[string][]string{
		"coffee": {"coff", "expresso", "cappi"},
		"tea":    {"chai", "latee", "hot-tea"},
	}
	fmt.Println(m["coffee"])

	m["other"] = []string{"Hot-chocolate"} //adding a key value in map
	fmt.Println(m)
	//maps do not always print out in the order that we add the indices but array and slices will
	//always print out in the order that we add the indices

	delete(m, "tea") //deleting a key
	fmt.Println(m)

	v, ok := m["tea"]
	fmt.Println(v, ok)

	m2 := m
	m2["coffee"] = []string{"coffee"}
	m["tea"] = []string{"ice-tea"}
	fmt.Println(m)
	fmt.Println(m2)

}
