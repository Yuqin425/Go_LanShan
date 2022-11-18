package main

import (
	"context"
	"fmt"
	"github.com/eiannone/keyboard"
	"time"
)

var (
	shot  = make(chan string, 1)
	aim   = make(chan string, 1)
	shoot = make(chan string, 1)
	point = "start"
)

func main() {
	shot <- point
	ctx, cancel := context.WithCancel(context.Background())
	fire(ctx)
	var (
		imput rune
		err   error
	)
	for imput != 'q' {
		imput, _, err = keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
	}
	cancel()
	fmt.Println()
	fmt.Println("打炮结束")
}

func fire(ctx context.Context) {
	for i := 1; i <= 10; i++ {
		go func() {
			for {
				select {
				case <-shot:
					time.Sleep(time.Second)
					fmt.Print("装弹 -> ")
					aim <- point
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	for i := 1; i <= 5; i++ {
		go func() {
			for {
				select {
				case <-aim:
					time.Sleep(time.Second)
					fmt.Print("瞄准 -> ")
					shoot <- point
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	for i := 1; i <= 3; i++ {
		go func() {
			for {
				select {
				case <-shoot:
					time.Sleep(time.Second)
					fmt.Println("发射！")
					shot <- point
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}
