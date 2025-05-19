package main

import (
	"fmt"
	"sort"
)

func main() {
	res := groupAnagrams([]string{"tan", "tea", "aet", "ate", "nat", "bat"})
	fmt.Println("res", res)
}

//简洁版排序
func groupAnagrams1(strs []string) [][]string {
	hashMap := map[string][]string{}
	for i := range strs {
		tmp := []byte(strs[i])
		sort.Slice(tmp, func(j, k int) bool {
			return tmp[j] < tmp[k]
		})
		hashMap[string(tmp)] = append(hashMap[string(tmp)], strs[i])
		//fmt.Println("i",i,"tmp",tmp,"hash",hashMap)
	}
	var res [][]string
	for _, value := range hashMap {
		res = append(res, value)
	}
	return res
}

type Pair struct {
	Index    int
	ByteList []byte
}

//第一次自己做的，复杂排序eat tea tan ate nat bat
func mygroupAnagrams(strs []string) [][]string {
	pairs := make([]Pair, len(strs))
	for i := range strs {
		byteList := []byte(strs[i])
		//分别对所有字符串进行排序
		sort.Slice(byteList, func(i, j int) bool {
			return byteList[i] < byteList[j]
		})
		//记录原索引
		pairs[i] = Pair{Index: i, ByteList: byteList}
	}
	//再对整体结构体数组，按字母顺序排序
	sort.Slice(pairs, func(i, j int) bool {
		start := 0
		for start <= len(pairs[i].ByteList)-1 && start <= len(pairs[j].ByteList)-1 {
			if pairs[i].ByteList[start] < pairs[j].ByteList[start] {
				return true
			}
			if pairs[i].ByteList[start] > pairs[j].ByteList[start] {
				return false
			}
			if pairs[i].ByteList[start] == pairs[j].ByteList[start] {
				start++
			}
		}
		//前缀字符一样，长度短的在前面
		return len(pairs[i].ByteList) <= len(pairs[j].ByteList)
	})
	fmt.Println("pairs", pairs)
	var res [][]string
	tmpRes := []string{strs[pairs[0].Index]}
	for i := 1; i <= len(pairs)-1; i++ {
		if string(pairs[i].ByteList) == string(pairs[i-1].ByteList) {
			tmpRes = append(tmpRes, strs[pairs[i].Index])
		} else if len(tmpRes) > 0 {
			res = append(res, tmpRes)
			tmpRes = []string{strs[pairs[i].Index]}
		}
	}
	if len(tmpRes) > 0 {
		res = append(res, tmpRes)
	}
	return res
}

//用字节计数代替排序
func groupAnagrams(strs []string) [][]string {
	hashMap := map[[26]int][]string{}
	for i := range strs {
		tmp := [26]int{}
		for j := range strs[i] {
			tmp[int(strs[i][j]-'a')]++
		}
		hashMap[tmp] = append(hashMap[tmp], strs[i])
	}
	var res [][]string
	for _, value := range hashMap {
		res = append(res, value)
	}
	return res
}
