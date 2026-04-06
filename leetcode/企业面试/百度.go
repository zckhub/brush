package main

import "fmt"
//长度为3的窗口从左到右滑过列表，返回每次窗口中的最大值
func main() {
    fmt.Println("Hello World!")
    nums := []int{1,3,2,-3,5,3,6,7,4,2,1,6,3}
    k := 3
    res := window2(nums,k)
    fmt.Println(res)
}

func window(nums []int,k int) []int{
    queue := make([]int,0,k)
    n := len(nums)
    res := make([]int,0,n-k)
    for i := 0;i<n;i++{
        //移除队头,保证k个元素
        for len(queue)>0 && queue[0]<i-k+1{
            queue = queue[1:]
        }

        //移除队尾
        for len(queue)>0 && nums[i] >= nums[queue[len(queue)-1]]{
            queue = queue[:len(queue)-1]
        }

        //加入队
        queue = append(queue, i)
        //形成窗口
        if i >= k-1{
            res = append(res, nums[queue[0]])
        }
    }
    return res
}
//nums := []int{5,3,4,-3,5,3,6,7,4,2,1,6,3}
func window2(nums []int,k int)[]int{
    queue := make([]int,k)
    n := len(nums)

    res := make([]int,0,n-k+1)
    for i := range nums{
        //删掉队列外的
        for len(queue)>0 && queue[0] <= i-k{
            queue = queue[1:]
        }

        //保证队头是最大的
        for  len(queue)>0 && nums[i]>=nums[queue[len(queue)-1]]{
            queue = queue[:len(queue)-1]
        }
        queue = append(queue, i)
        if i-k+1>=0{
            res = append(res, nums[queue[0]])
        }
    }
    return res
}