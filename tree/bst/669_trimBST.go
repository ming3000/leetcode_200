package bst

import "leetcode_200/common"

func trimBST(root *common.TreeNode, L int, R int) *common.TreeNode {
	if root == nil {
		return nil
	} // if>

	if root.Val > R {
		return trimBST(root.Left, L, R)
	} // if>
	if root.Val < L {
		return trimBST(root.Right, L, R)
	} // if>

	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
