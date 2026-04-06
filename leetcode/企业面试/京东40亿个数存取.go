package main

import (
	"fmt"
)
func main() {
	newNum := &BitMap{Data: make([]byte, 10000000)}
	testNums := []int{100,200,400000,134214}
	for i := range testNums{
		newNum.SetMap(testNums[i])
	}
	check := 134214
	res := newNum.GetMap(check)
	fmt.Println(res)
	check = 13421411
	res = newNum.GetMap(check)
	fmt.Println(res)

}

type BitMap struct{
	Data []byte
}

func(b *BitMap) SetMap(num int){
	index := num / 8
	indexPos := num % 8
	tmp := byte(1) << indexPos
	oldByte := b.Data[index]
	newBtye := oldByte | tmp //对应位设置成1
	b.Data[index] = newBtye
}

func(b *BitMap) GetMap(num int)bool{
	index := num / 8
	indexPos := num % 8
	tmp := 1<<indexPos
	curByte := b.Data[index]
	//对应位置必须都是1
	res := tmp & int(curByte) !=0
	return res
}