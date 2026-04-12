package main

import (
	"fmt"
)
func main() {
	var res byte
	var resBool bool
	capacity := 3
	rb := NewRingBuf(capacity+1)
	rb.Write('a')
	rb.Write('b')
	rb.Write('c')
	res,resBool = rb.Read()
	fmt.Println(string(res),resBool)

	rb.Write('d')
	rb.Write('c')
	res,resBool = rb.Read()
	fmt.Println(string(res),resBool)
	// rb.Write('e')
	// res,resBool = rb.Read()
	// fmt.Println(string(res),resBool)
}

type RingBuffer struct{
	buffer []byte
	capacity int
	head int
	tail int
}

func NewRingBuf(capacity int) (*RingBuffer){
	if capacity<1{
		return nil
	}
	res := &RingBuffer{
		buffer: make([]byte,capacity),
		capacity: capacity}
	return res
}

func (r *RingBuffer) Write(b byte){
	//队列已经满
	if (r.tail +1)%r.capacity ==r.head{
		fmt.Println("队列已经满，不能加入",string(b))
		return
	}
	r.buffer[r.tail] = b
	r.tail = (r.tail+1)%r.capacity
	return
}

func(r *RingBuffer) Read()(res byte,resBool bool){
	if r.head == r.tail{
		resBool = false
		return
	}
	b := r.buffer[r.head]
	r.head = (r.head+1) %r.capacity
	return b,true
}

