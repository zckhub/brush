package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var sycnTest Sync
	for i := 0;i<10;i++{
		go func(i int){
			val := sycnTest.Exec("张三",func ()(int, error)  {
				fmt.Println(" 协程",i,"真正去查询张三")
				return i,nil
			})
			fmt.Println("协程",i,"拿到结果",val)
		}(i)
	}
	time.Sleep(3*time.Second)
}

type Call struct{
	WG sync.WaitGroup
	Val int
	Err error
}

type Sync struct{
	Mu sync.Mutex
	M map[string]*Call //key 正在执行的call
}

func(this *Sync) Exec(key string,callBack func() (int,error)) int{
	this.Mu.Lock()
	if this.M == nil{
		this.M = make(map[string]*Call)
	}

	//已经有人在查key
	if value,ok := this.M[key];ok{
		this.Mu.Unlock()
		value.WG.Wait()
		return value.Val
	}

	//没人查过，写入
	writeCall := &Call{}
	writeCall.WG.Add(1)
	this.M[key] = writeCall
	this.Mu.Unlock()

	//查数据
	writeCall.Val,writeCall.Err = callBack()
	writeCall.WG.Done()
	this.Mu.Lock()
	delete(this.M,key)
	this.Mu.Unlock()
	return writeCall.Val

}