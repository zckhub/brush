package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Right, val)
	}
	if root.Val < val {
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
