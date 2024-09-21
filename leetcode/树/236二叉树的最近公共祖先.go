package main

/*
两种方法都会把所有节点都访问一遍，时间和空间复杂度都是o(n)：
1。从上到下，递归遍历是否包含p,q。由于递归是反向执行，第一次返回的就是实际结果
2。用一个map记录所有结点的父节点，再用一个map记录 遍历p visited过的父节点，
遍历q判断父节点是否存在于visited，如果存在则返回
*/
var res *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res = nil
	hasPQ(root, p, q)
	return res
}

//返回 是否有p,是否有q,p或q的祖先
func hasPQ(node, p, q *TreeNode) (hasP bool, hasQ bool) {
	if node == nil {
		return false, false
	}

	leftHasP, leftHasQ := hasPQ(node.Left, p, q)

	rightHasP, rightHasQ := hasPQ(node.Right, p, q)
	hasP = leftHasP || rightHasP || node.Val == p.Val
	hasQ = rightHasQ || leftHasQ || node.Val == q.Val

	if hasP && hasQ && res == nil {
		res = node
	}
	return
}
