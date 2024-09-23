package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，
另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。*/
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go GenNum(ch, &wg)
		wg.Add(1)
		go FmtNum(ch, &wg)
		time.Sleep(1 * time.Second)

	}
	wg.Wait()
}

func GenNum(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 设置随机种子，确保每次运行生成不同的随机数序列
	rand.Seed(time.Now().UnixMilli())
	// 生成范围在 0 到 100 之间的随机整数
	randomInt := rand.Intn(101)
	fmt.Println("gen rand", randomInt)
	ch <- randomInt
}

func FmtNum(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	num := <-ch
	fmt.Println("fmt rand", num)
}
