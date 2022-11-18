package main

import (
	"fmt"
	"time"
)

var (
	Wallet = 0
	vivo50 = make(chan int)
)

func main() {
	for i := 0; i < 10000; i++ {
		go vPaopao50()
	}
	go func() {
		for {
			Wallet += <-vivo50
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("泡泡现在有", Wallet, "元")
}

func vPaopao50() {
	vivo50 <- 50
}
