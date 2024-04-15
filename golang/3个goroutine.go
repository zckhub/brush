package main

import (
	"fmt"
	"sync"
)

func main() {
	//启动3个goroutine 循环100次顺序打印123
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go PrintOne(ch1, ch2, 1, &wg)
	go PrintOne(ch2, ch3, 2, &wg)
	go PrintOne(ch3, ch1, 3, &wg)
	ch1 <- 1
	wg.Wait()
}

func PrintOne(chin chan int, chout chan int, number int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-chin
		fmt.Println(number)
		chout <- 0
	}
	//最后要把channel资源释放
	if number == 1 {
		<-chin
	}
}
