package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return root
}

//找val的父节点
func findParent(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Right.Val == val || root.Left.Val == val {
		return root
	}
	if root.Val < val {
		return findParent(root.Right, val)
	}
	return findParent(root.Left, val)
}

/*
寻找val的前驱结点
1.有左子树 -- 左子树的最大值
2。无左子树 -- 这时要找左边第一个结点：一直向上找父节点，
直到某个父节点是右子树，这个父节点的父节点就是答案
*/
func findPredecesor(root *TreeNode, val int) *TreeNode {
	node := searchBST(root, val)
	if node.Left != nil {
		return findMaxNode(node.Left)
	}
	return node
}

func findMaxNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right != nil {
		return findMaxNode(root.Right)
	}
	return root.Right
}

/*
寻找后继结点
*/
func findSuccessor(root *TreeNode, val int) *TreeNode {
	node := searchBST(root, val)
	if node.Right != nil {
		return findMinNode(node.Right)
	}
	return node
}

func findMinNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		return findMinNode(root.Left)
	}
	return root.Left
}
