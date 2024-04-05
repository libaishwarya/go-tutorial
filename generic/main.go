package main

import "fmt"

type test struct {
}

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	fmt.Println(checkEqual(a, b))

	c := []string{"a", "b", "c"}
	d := []string{"a", "b", "c"}
	fmt.Println(checkEqual(c, d))

}

func checkEqual[V int | string | test | bool](a, b []V) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
