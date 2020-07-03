package dfs

import "leetcode_200/common"

func preorderTraversal(root *common.TreeNode) []int {
	if root == nil {
		return nil
	} // if>

	ret := make([]int, 0)
	stack := make([]*common.TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		theRoot := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, theRoot.Val)

		// pre-order, so right child first
		if theRoot.Right != nil {
			stack = append(stack, theRoot.Right)
		} // if>>
		if theRoot.Left != nil {
			stack = append(stack, theRoot.Left)
		}
	} // for>
	return ret
}
