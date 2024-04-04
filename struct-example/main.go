package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	in := bufio.NewReader(os.Stdin)    //wrap the keyboard ip with buffered IO reader
	choice, _ := in.ReadString('\n')   //reads whatever the user provides
	choice = strings.TrimSpace(choice) //eliminates any while spaces
	type menuItem struct {
		name   string
		prices map[string]float64
	}
	menu := []menuItem{
		{name: "coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "tea", prices: map[string]float64{"single": 1.65, "double": 1.80, "triple": 1.95}},
	}
	fmt.Println(menu)

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}

	}
}
