package main

/*
1.删除的是叶子节点-直接删除
2.如果存在左子树 - 结点值替换为前驱结点，再删除前驱结点
	由于左子树存在，找前驱直接找左子树最大
3。如果存在右子树-结点值替换为后继，再删除后继
*/
func deleteNode(root *TreeNode, key int) *TreeNode {
	node := searchBST(root, key)
	if node.Right == nil && node.Left == nil {
		node = nil
		return root
	}
	if node.Left != nil {
		leftMaxNode := findMaxNode(node.Left)
		node.Val = leftMaxNode.Val
		leftMaxNode = nil
	}
	if node.Right != nil {
		rightMixNode := findMinNode(node.Right)
		node.Val = rightMixNode.Val
		rightMixNode = nil
	}
	return root
}
