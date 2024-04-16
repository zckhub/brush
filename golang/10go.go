package main

import (
	"fmt"
	"sync"
	"time"
)

//编写一个程序限制10个goroutine执行，每执行完一个goroutine就放一个新的goroutine进来

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go task(i, ch, &wg)
	}
	for k := 0; k < 10; k++ {
		ch <- k
	}

	wg.Wait()
}

func task(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ch
	time.Sleep(3 * time.Second)
	fmt.Println(i)

	ch <- i

}
