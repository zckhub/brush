package main

import "fmt"

func main(){
	testSlice2()
}
func testSlice(){
	slice1 := make([]int64,0,2)
	slice1 = append(slice1, 100) //100
	fmt.Println(slice1)
	slice2 := slice1
	slice2 = append(slice2, 200) //底层数组变成了 【100 200] 但是slice1长度为1，还是显示【100】
	slice1 = append(slice1, 300) //底层数组变成了[100,300] 因为slice1长度是1，所以append是从第1位开始修改，把200改成了300
	fmt.Println(slice1)
	fmt.Println(slice2)
}
func testSlice2(){
	slice1 := make([]int64,0,2)
	slice1 = append(slice1, 100) //slice1: 底层数组[100, 0]，长度=1
	slice2 := slice1 // 
	slice2 = append(slice2, 200) //底层数组变为：[100, 200]  slice2 长度变为2  slice1 长度仍为1

	slice3 := slice1[1:] //slice1[1:] 从索引1开始，slice1长度为1，所以slice3长度为0 ｜ slice3共享底层数组，起始位置是索引1，容量=1
	slice3 = append(slice3, 300) //100 300
	slice3 = append(slice3, 400) //slice3 扩容，创建新数组 300,400
	// fmt.Println(slice3)
	slice3[0] = 500
	fmt.Println(slice1)
	fmt.Println(slice2)	
	fmt.Println(slice3)

}

//当channels中的任意一个channel有输出时，该函数返回channel close
func or(channels ...<-chan interface{}) <-chan interface{

}
//Go 协程的数量变少一些。
//那你监听这一个 channel，你是用一个协程去监听它吗？
