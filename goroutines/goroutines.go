package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("goroutines")
	go g1()
	time.Sleep(1000000000000)
}

func g1() {
	fmt.Println("g1 running")
}

func g2() {
	fmt.Println("g2 running")
}
