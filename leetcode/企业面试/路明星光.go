package main

import "fmt"

// 面向对象设计 - 键值对类
// 题目描述
// 设计一个管理数据键值对的类（Class），语言不限：​

// 背景：​
// 1. 这个类需要以任意形式存储形如 (key, value, timestamp) 的数据，其中 key 和 value 均为字符串格式，timestamp 为整数类型的时间戳​
// 2. 有一个数据源会以串行的方式（无需考虑并行化产生的竞争），调用该类的接口，进行数据的存储和查询操作​
// 3. 对于数据的存储操作，保证调用方提供的 timestamp 是严格递增的。而对于查询操作，则没有限制​

// 需求：​
// 1. 提供一个插入数据的接口：给定 (key, value, timestamp)，对该数据进行存储​
// 2. 提供一个查询最新数据的接口：给定 key，得到最新的 value​
// 3. 提供一个支持查询历史数据的接口：给定 key 和 timestamp，返回时间小于等于 timestamp 的最新的 value​
// 4. 需要对「背景中第 3 条」进行严格的检查，保证调用方没有进行错误的调用​​​

func main(){
    kv := NewKeyValueStore()
    resBool := kv.Insert("name","123",100)
    fmt.Println(resBool)
    resBool = kv.Insert("name","233",200)
    fmt.Println(resBool)

    resValue,resBool := kv.GetLatest("name")
    fmt.Println(resValue,resBool)
    resValue,resBool = kv.GetLatest("age")
    fmt.Println(resValue,resBool)

    resBool = kv.Insert("name","433",150)
    fmt.Println(resBool)

    resBool = kv.Insert("name","324",350)
    fmt.Println(resBool)

    resBool = kv.Insert("name","2345",450)
    fmt.Println(resBool)
    resValue,resBool = kv.GetLatest("name")
    fmt.Println(resValue,resBool)

    resValue,resBool = kv.Get("name",400)
    fmt.Println(resValue,resBool)
}
type KeyValueStore struct{
    data map[string][]TimeValueData
}

type TimeValueData struct{
    Value string
    Timestamp int
}

func NewKeyValueStore() *KeyValueStore{
    res := &KeyValueStore{
        data: make(map[string][]TimeValueData),
    }
    return res
}

func(this *KeyValueStore) GetLatest(key string)(resValue string,resBool bool){
    values,ok := this.data[key]
    if !ok{
        resBool = false
        return
    }
    if len(values) == 0{
        resBool = false
        return
    }

    resValue = values[len(values)-1].Value
    resBool = true
    return
}

func(this *KeyValueStore) Insert(key string,value string,timestamp int)(resBool bool){
    values,_ := this.data[key]
    if len(values) != 0{
        if timestamp <= values[len(values)-1].Timestamp{
            resBool = false
            return 
        }
    }
    this.data[key] = append(this.data[key], TimeValueData{Value: value,Timestamp: timestamp})
    resBool = true
    return
}

func(this *KeyValueStore) Get(key string,timestamp int)(resValue string,resBool bool){
    values,ok := this.data[key]
        if !ok{
        resBool = false
        return
    }
    if len(values) == 0{
        resBool = false
        return
    }

    if timestamp < values[0].Timestamp{
        resBool = false
        return 
    }

    //二分法找到timestamp的位置
    left:=0
    right := len(values)-1
    for left<right{
        mid := (left+right)/2
        if values[mid].Timestamp<timestamp{
            resValue = values[mid].Value
            resBool = true
            left = mid+1
        }else{
            right = mid-1
        }
    }
    return
}