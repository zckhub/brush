package main

import (
	"fmt"
	"sync"
)

//执行命令for i in {1..10}; do go run 2000个人抢500张票.go | grep 购买成功| wc -l;done
//2000个协程抢占500个资源，如何避免竞争
func main() {
	ticker := 500
	var mu sync.Mutex
	var wg sync.WaitGroup
	//2000个协程并发抢票
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(userId int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if ticker > 0 {
				ticker--
				fmt.Printf("用户 %d 购买成功，剩余票数 %d\n", userId, ticker)
			} else {
				fmt.Printf("用户 %d 购买失败\n", userId)
			}
		}(i)
	}
	wg.Wait()
}
