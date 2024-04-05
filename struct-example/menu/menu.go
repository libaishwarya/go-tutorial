package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

type menu []menuItem

func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
	}
}

func (m *menu) add() error {
	fmt.Println("Enter the name of the new item")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)

	for _, item := range data {
		if item.name == name {
			return errors.New("menu already exists")

		}
	}
	*m = append(*m, menuItem{name: name, prices: make(map[string]float64)})
	return nil

}

var in = bufio.NewReader(os.Stdin)

func Print() {
	data.print()
}

func AddItem() error {
	return data.add()

}
