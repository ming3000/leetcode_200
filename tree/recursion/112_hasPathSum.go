package recursion

import "leetcode_200/common"

func hasPathSum(root *common.TreeNode, sum int) bool {
	if root == nil {
		return false
	} // if>
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	} // if>
	leftHasPathSum := hasPathSum(root.Left, sum-root.Val)
	rightHasPathSum := hasPathSum(root.Right, sum-root.Val)
	return leftHasPathSum || rightHasPathSum
}
