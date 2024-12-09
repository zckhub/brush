package main

import (
	"fmt"
	"sync"
	"time"
)

/*编写一个程序限制10个goroutine执行，一次执行10个
每执行完一个goroutine就放一个新的goroutine进来
*/
func main5() {
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

func main() {
	chanNum := 10
	ch := make(chan int, chanNum)
	var wg sync.WaitGroup
	//ch <- 1
	for i := 0; i < chanNum; i++ {
		ch <- i
	}
	for j := 0; j < chanNum*10; j++ {
		wg.Add(1)
		go task1(ch, j, &wg)
	}
	wg.Wait()
	fmt.Println("end")
}

func task1(ch chan int, j int, group *sync.WaitGroup) {
	defer group.Done()
	<-ch
	fmt.Println("j=", j)
	time.Sleep(2 * time.Second)
	ch <- j
}
