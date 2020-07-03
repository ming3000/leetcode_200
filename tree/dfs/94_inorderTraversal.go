package dfs

import "leetcode_200/common"

func inorderTraversal(root *common.TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*common.TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		} // for>>
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		root = node.Right
	} // for>

	return ret
}
