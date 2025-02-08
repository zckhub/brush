package main

import (
	"fmt"
	"sync"
)

func main() {
	printfmt("hello world")
}
func printfmt(word string) {
	//交替打印字符串
	ch := make(chan int)
	strListCh := make(chan byte, len(word))

	for i := range word {
		strListCh <- word[i]
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go printA(wg, ch, strListCh, 1)
	wg.Add(1)
	go printA(wg, ch, strListCh, 2)
	ch <- 1
	wg.Wait()
}
func printA(wg *sync.WaitGroup, ch chan int, strListCh chan byte, flag int) {
	defer wg.Done()
	for {
		_, ok := <-ch
		if ok {
			//取出ch，最后写入ch
			//value, okList := <-strListCh
			//fmt.Println("printA value", string(value), "okList", okList)
			if len(strListCh) != 0 {
				fmt.Println("printA ", string(<-strListCh), "flag", flag)
			} else {
				close(ch)
				return
			}
			ch <- 1
		} else {
			fmt.Println("printA end ch")
			return
		}
	}
}
