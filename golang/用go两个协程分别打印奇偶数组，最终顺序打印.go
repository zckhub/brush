package main

import (
	"fmt"
	"time"
)

func main() {
	oddCh := make(chan struct{})
	evenCh := make(chan struct{})
	go func() {
		for i := 1; i <= 100; i += 2 {
			<-oddCh
			fmt.Println(i)
			evenCh <- struct{}{}
		}
	}()

	go func() {
		oddCh <- struct{}{}
		for i := 2; i <= 100; i += 2 {
			<-evenCh
			fmt.Println(i)
			oddCh <- struct{}{}
		}
	}()
	time.Sleep(5 * time.Second)
}
