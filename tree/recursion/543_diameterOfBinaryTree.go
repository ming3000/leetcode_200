package recursion

import "leetcode_200/common"

func diameterOfBinaryTree(root *common.TreeNode) int {
	var max int
	depth(root, &max)
	return max
}

func depth(root *common.TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	leftDepth := depth(root.Left, max)
	rightDepth := depth(root.Right, max)
	*max = common.Max(*max, leftDepth+rightDepth)
	return common.Max(leftDepth, rightDepth) + 1
}
