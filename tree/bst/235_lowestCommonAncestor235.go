package bst

import "leetcode_200/common"

func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	node := root
	for node != nil {
		if node.Val < p.Val && node.Val < q.Val {
			node = node.Right
		} else if node.Val > p.Val && node.Val > q.Val {
			node = node.Left
		} else {
			return node
		}
	} // for>

	return nil
}
