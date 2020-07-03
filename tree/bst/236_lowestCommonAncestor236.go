package bst

import "leetcode_200/common"

func lowestCommonAncestor236(root, p, q *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	} // if>

	left := lowestCommonAncestor236(root.Left, p, q)
	right := lowestCommonAncestor236(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} // if>

	if left == nil {
		return right
	} else {
		return left
	}
}
