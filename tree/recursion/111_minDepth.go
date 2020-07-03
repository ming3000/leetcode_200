package recursion

import "leetcode_200/common"

func minDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	} // if>
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	} // if>
	return common.Min(left, right) + 1
}
