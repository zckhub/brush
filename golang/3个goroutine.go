package main

import (
	"fmt"
	"sync"
)

func main1() {
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
	//在第一个协程内最后要把ch1取出，否则第三个协程写入了ch1，没有地方取出造成死锁。
	if number == 1 {
		<-chin
	}
}

func main3() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(3)
		go FmtGo(ch1, ch2, 1, &wg, i)
		go FmtGo(ch2, ch3, 2, &wg, i)
		go FmtGo(ch3, ch1, 3, &wg, i)
	}
	ch1 <- 1
	<-ch1

	wg.Wait()

}
func FmtGo(ch1 chan int, ch2 chan int, fmtNum int, group *sync.WaitGroup, i int) {
	defer group.Done()
	<-ch1
	fmt.Println(fmtNum, "i=", i)
	ch2 <- 1
}

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go printNum(ch1, ch2, 1, &wg)
	go printNum(ch2, ch3, 2, &wg)
	go printNum(ch3, ch1, 3, &wg)
	ch1 <- 1
	wg.Wait()
	fmt.Println("end")
}

func printNum(ch1 chan int, ch2 chan int, num int, group *sync.WaitGroup) {
	defer group.Done()
	for i := 0; i < 100; i++ {
		<-ch1
		fmt.Println(num, "i=", i)
		ch2 <- 1
	}
	if num == 1 {
		<-ch1
	}

}
