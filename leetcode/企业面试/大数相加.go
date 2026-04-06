package main

import (
	"fmt"
)
func main() {
	fmt.Println("hello")
	nums1 := "99"
	nums2 :=  "99"
	res := addString(nums1,nums2)
	fmt.Println(res)
}
// 进位tmp
//   "231312" 
// "4143434"
func addString(nums1 string,nums2 string) string{
	i := len(nums1) - 1
	j := len(nums2) - 1
	tmpJinwei := 0 //进位
	res := make([]byte,0)
	for i >=0 || j >=0 || tmpJinwei>0{
		tmpi:=0
		if i >=0{
			tmpi = int(nums1[i] - '0')
			i--
		} 
		tmpj := 0
		if j >= 0{
			tmpj = int(nums2[j] - '0')
			j--
		}
		tmpSum := tmpi+tmpj+tmpJinwei
		tmpJinwei = tmpSum/10
		curRes := tmpSum % 10
		res = append(res, byte(curRes)+'0')
	}

	//翻转
	left,right := 0,len(res)-1
	for left < right{
		res[left],res[right] = res[right],res[left]
		left++
		right--
	}
	return string(res)
}