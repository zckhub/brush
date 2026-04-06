package main

import (
	"context"
	"fmt"
	"time"
)
//100000以内的数字两两相乘，一共多少种结果
func main() {
	fmt.Println(20000*20000 /8 / 1024/1024)
	start := time.Now()
	N := int(100000)
	res := count(N,120 *time.Second)
	duration := time.Since(start)
	fmt.Println("res= ",res)
	fmt.Println("dur= ",duration)
}

//100000*100000 /8 / 1024/1024
func count(N int,timeout time.Duration)int{
	
	bits := make([]uint64,(N*N+64)/64+1)
	ctx,cancel := context.WithTimeout(context.Background(),timeout)
	defer cancel()
	for i := 1;i<=N;i++{
		select {
		case <- ctx.Done():
			return -1
		default:
		}
		for j:= 1;j<=N;j++{
			mux := i*j
			tmpIndex := mux/64
			tmpPos := mux % 64
			mask := uint64(1) << tmpPos
			bits[tmpIndex] = bits[tmpIndex] | mask
		}
	}
	// fmt.Print(bits)

	resCnt := 0
	for _,value := range bits{
		for value >0 {
			resCnt++
			value &=value-1
		}
	}
	return resCnt
}