package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	ValueMap  map[int]int
	ValueList []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{ValueMap: map[int]int{},
		ValueList: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.ValueMap[val]; ok {
		return false
	}
	this.ValueList = append(this.ValueList, val)
	this.ValueMap[val] = len(this.ValueList) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.ValueMap[val]
	if !ok {
		return false
	}
	delete(this.ValueMap, val)
	//保持元素索引位置不变
	this.ValueList[index] = this.ValueList[len(this.ValueList)-1]
	this.ValueList = this.ValueList[:len(this.ValueList)-1]
	if index < len(this.ValueList) {
		this.ValueMap[this.ValueList[index]] = index
	}
	fmt.Println("ValueMap", this.ValueMap, "ValueList", this.ValueList)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	randIndex := rand.Intn(len(this.ValueList))
	return this.ValueList[randIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
