package main

import "fmt"

//两个无序链表合成一个有序链表
func main() {
	l1List := []int{3, 1, 4, 5}
	l2List := []int{2, 6, 1, 7}
	l1 := genList(l1List)
	l2 := genList(l2List)
	l1Pre := l1
	for l1.Next != nil {
		l1 = l1.Next
	}
	l1.Next = l2
	resHead := sortList(l1Pre)
	for resHead != nil {
		fmt.Println("res", resHead.Val)
		resHead = resHead.Next
	}
}

func genList(nums []int) *ListNode {
	l1 := &ListNode{}
	l1Pre := l1
	for i := range nums {
		l1.Next = &ListNode{Val: nums[i]}
		l1 = l1.Next
	}
	l1 = l1Pre.Next
	return l1
}

//对单个链表归并排序
func sortList(l1 *ListNode) *ListNode {
	if l1.Next == nil {
		return l1
	}
	mid := findMid(l1)
	right := mid.Next
	mid.Next = nil

	left := sortList(l1)
	right = sortList(right)
	res := merge(left, right)
	return res
}

func merge(l1, l2 *ListNode) *ListNode {
	//合并有序链表
	pre := &ListNode{}
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next

	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return pre.Next
}

//快慢指针找中点
func findMid(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
