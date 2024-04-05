package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structExample/menu"
)

var in = bufio.NewReader(os.Stdin)

func main() {

loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add Item")
		fmt.Println("q) Quit")

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			err := menu.AddItem()
			if err != nil {
				fmt.Println(fmt.Errorf("invalid input: %w", err))
			}
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}
