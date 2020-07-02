package recursion

import "leetcode_200/common"

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	} // if>
	return common.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
