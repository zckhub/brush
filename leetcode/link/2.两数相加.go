package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := &ListNode{}
	head = cur
	var add int
	var residual int
	for l1 != nil || l2 != nil {
		cur.Val = residual
		cur.Next = &ListNode{}
		cur = cur.Next
		if l1 == nil && l2 != nil {
			residual = (l2.Val + add) % 10
			add = (l2.Val + add) / 10
			l2 = l2.Next
		}
		if l2 == nil && l1 != nil {
			residual = (l1.Val + add) % 10
			add = (l1.Val + add) / 10
			l1 = l1.Next
		}
		if l1 != nil && l2 != nil {
			residual = (l1.Val + l2.Val + add) % 10
			add = (l1.Val + l2.Val + add) / 10
			l1 = l1.Next
			l2 = l2.Next
		}

	}
	cur.Val = residual
	if add != 0 {
		cur.Next = &ListNode{}
		cur = cur.Next
		cur.Val = add
	}
	return head.Next
}
