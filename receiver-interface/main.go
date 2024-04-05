package main

import (
	"bytes"
	"fmt"
	"strings"
)

type printer interface {
	Print() string
}

type user struct {
	id   int
	name string
}

func (u user) Print() string {
	return fmt.Sprintf("%v (%v)\n", u.name, u.id)
}

type menuItem struct {
	name  string
	price map[string]float64
}

func (mi menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(mi.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range mi.price {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}
	return b.String()
}

func main() {
	var p printer
	p = user{name: "aishu", id: 1}
	fmt.Println(p.Print())

	p = menuItem{name: "Coffee",
		price: map[string]float64{"small": 10,
			"medium": 20,
			"large":  30}}
	fmt.Println(p.Print())

	u, ok := p.(menuItem)
	fmt.Println(u, ok)
}
