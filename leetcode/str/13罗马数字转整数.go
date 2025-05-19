package main

import "fmt"

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

//找到后比前小的，计算之前的数字，再加后面的数字
//MCMXCIV  M = 1000, CM = 900, XC = 90, IV = 4
func romanToInt(s string) int {
	hashMap := map[byte]int{}
	hashMap['I'] = 1
	hashMap['V'] = 5
	hashMap['X'] = 10
	hashMap['L'] = 50
	hashMap['C'] = 100
	hashMap['D'] = 500
	hashMap['M'] = 1000
	var res int
	for i := range s {
		if i < len(s)-1 && hashMap[s[i]] < hashMap[s[i+1]] {
			res -= hashMap[s[i]]
		} else {
			res += hashMap[s[i]]
		}

	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var cnt int
	var res, preNode, rightNode, tmpPre, leftNode, startNode *ListNode
	if left != 0 {
		res = head
	}
	startNode = &ListNode{}
	startNode.Next = head
	preNode = startNode

	for cnt <= right {
		cnt++
		if head != nil && cnt < left {
			preNode = head
			head = head.Next
		}
		if head != nil && cnt == left {
			leftNode = head
			tmpPre = head
			rightNode = head //初始化right 可能left==right
			head = head.Next
		}
		if head != nil && cnt > left && cnt <= right {
			tmpNext := head.Next
			head.Next = tmpPre
			tmpPre = head
			rightNode = head

			head = tmpNext
		}
		//最后一位时，head为right.Next
	}

	preNode.Next = rightNode //->5
	leftNode.Next = head     //3->nil
	if preNode == startNode {
		return rightNode
	}
	return res
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	startNode := &ListNode{}
	startNode.Next = head
	pre := startNode
	for head != nil {
		tail := pre
		fmt.Println("head1", head.Val)
		for i := 0; i < k; i++ {

			tail = tail.Next
			if tail == nil {
				return startNode.Next
			}
		}
		fmt.Println("head2", head.Val)
		nex := tail.Next
		fmt.Println("head", head.Val, "tail", tail.Val, "nex", nex.Val)

		head, tail = reverseNode(head, tail)
		fmt.Println("head", head.Val, "tail", tail.Val)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
		fmt.Println("head", head.Val, "tail", tail.Val)
	}
	return startNode.Next
}

func reverseNode(head *ListNode, tail *ListNode) (resHead *ListNode, resTail *ListNode) {
	cur := head
	pre := tail.Next //末尾元素需要指向的
	for cur != tail {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	resHead = tail
	resTail = head

	return

}
