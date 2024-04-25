package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		time.Sleep(12 * time.Second)
		c <- 0
	}()

	for {
		select {
		case <-c:
			fmt.Println("end")
			return
		default:
			time.Sleep(1000 * time.Millisecond)
			// 1, 2, 1, 4
			fmt.Println("test")
		}
	}

	c = make(chan int)

	go func() {
		time.Sleep(12 * time.Second)
		c <- 0
	}()

	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-c:
			fmt.Println("end")
			return
		case <-ticker.C:
			fmt.Println("test")
		}
	}

}
