//instead of zero based index we have our own key type/index type (customize key type)
//maps are reference type ie. they dont hold their own data but they point to data
//..stored somewhere

package main

import "fmt"

func main() {
	var m map[string]int //map keyword followed by key type(string)and value type(int)
	fmt.Println(m)
	m = map[string]int{"foo": 1, "bar": 2, "sun": 32, "hen": 54}
	fmt.Println(m)

	fmt.Println(m["bar"]) //Accessing value in map
	m["bar"] = 99         //update value in map
	fmt.Println(m)

	m["gun"] = 6 //add value to map
	fmt.Println(m)

	delete(m, "bar") //deleting value
	fmt.Println(m)

	fmt.Println(m["bar"])
	v, ok := m["bar"]  //
	fmt.Println(v, ok) //key is not present in map so ok returns False and v returns 0 since not found
	g, ok := m["foo"]
	fmt.Println(g, ok)

	m2 := m
	m["foo"], m2["gun"], m["bar"] = 65, 83, 32
	fmt.Println("m value is: ", m)
	fmt.Println("m2 value is: ", m2)

}
