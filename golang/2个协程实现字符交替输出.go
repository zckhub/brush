package main

import "sync"

func PrintTwoGo() {
	str := "hello world"
	//无缓冲 用户控制协程交替执行
	control := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go PrintLetter(&wg)

}
func PrintLetter(wg *sync.WaitGroup) {
	defer wg.Done()
}
