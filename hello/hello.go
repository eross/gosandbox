package main

import (
	"fmt"
	"time"

	"rsc.io/quote"
)

func main() {
	var BigArray [10000 * 1024 * 1024]byte
	for i := 0; i < 10000*1024*1024; i++ {
		BigArray[i] = 42
	}
	BigArray[200] = 10
	time.Sleep(10 * time.Second)
	fmt.Println(quote.Go())
}
