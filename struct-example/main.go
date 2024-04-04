package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type menuItem struct {
		name   string
		prices map[string]float64
	}
	menu := []menuItem{
		{name: "coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "tea", prices: map[string]float64{"single": 1.65, "double": 1.80, "triple": 1.95}},
	}
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	fmt.Println("2) Add Item")
	fmt.Println("q) Quit")
	in := bufio.NewReader(os.Stdin)

loop:
	for {
		choice, _ := in.ReadString('\n')
		switch strings.TrimSpace(choice) {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, cost := range item.prices {
					fmt.Printf("\t%10s%10.2f\n", size, cost)
				}
			}
		case "2":
			fmt.Println("Enter the name of the new item")
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")

		}
	}
}
