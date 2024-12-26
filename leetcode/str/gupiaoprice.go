package main

import "fmt"

type LRUCache struct {
	Capacity int
	Head     *LRULink         //链表
	Tail     *LRULink         //链表结尾，用于删除最少访问的数据
	LRUhash  map[int]*LRULink //key,链表节点指针
	CurLen   int
}
type LRULink struct {
	Key   int
	Value int
	Pre   *LRULink
	Next  *LRULink
}

func Constructor(capacity int) LRUCache {
	head := &LRULink{}
	tail := &LRULink{}
	head.Next = tail
	tail.Pre = head
	lruCache := LRUCache{
		Capacity: capacity,
		Head:     head,
		Tail:     tail,
		LRUhash:  map[int]*LRULink{},
		CurLen:   0,
	}
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	point := this.LRUhash[key]
	if point == nil {
		return -1
	}
	//把point从现有位置放最前面，

	point.Pre.Next = point.Next
	point.Next.Pre = point.Pre

	point.Next = this.Head.Next
	point.Pre = this.Head
	this.Head.Next.Pre = point
	this.Head.Next = point

	//fmt.Println("head",this.Head.Next.Value)

	//fmt.Println("tail",this.Tail.Pre.Value)

	return point.Value
}

func (this *LRUCache) Put(key int, value int) {
	point := this.LRUhash[key]
	if point != nil {
		//替换原value,更新key到link最前面
		point.Value = value
		point.Pre.Next = point.Next
		point.Next.Pre = point.Pre

		point.Next = this.Head.Next
		point.Pre = this.Head
		this.Head.Next.Pre = point
		this.Head.Next = point
		return
	}

	if this.Capacity <= this.CurLen {
		//容量超了，删除最后一位
		//fmt.Println("this.Tail.Pre.Pre",this.Tail.Pre.Value)
		delete(this.LRUhash, this.Tail.Pre.Key)

		this.Tail.Pre.Pre.Next = this.Tail
		this.Tail.Pre = this.Tail.Pre.Pre
		this.CurLen--

	}
	//添加新元素到第一位
	this.CurLen++
	newPoint := &LRULink{Value: value, Key: key}
	newPoint.Next = this.Head.Next
	newPoint.Pre = this.Head
	this.Head.Next.Pre = newPoint
	this.Head.Next = newPoint
	this.LRUhash[key] = newPoint

}

func main() {
	obj := Constructor(2)

	obj.Put(1, 1)
	fmt.Println(obj.Get(1))

	obj.Put(2, 2)
	fmt.Println(obj.Get(2))
	obj.Put(3, 3)
	fmt.Println(obj.Get(1))

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
