package dfs

import "leetcode_200/common"

// 前序遍历为 root -> left -> right，后序遍历为 left -> right -> root。
// 可以修改前序遍历成为 root -> right -> left，那么这个顺序就和后序遍历正好相反。
func postorderTraversal(root *common.TreeNode) []int {
	if root == nil {
		return nil
	} //if>

	ret := make([]int, 0)
	stack := make([]*common.TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		} // if>>
		ret = append(ret, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		} // if>>
		if node.Right != nil {
			stack = append(stack, node.Right)
		} // if>>
	} // for>

	for i := 0; i < len(ret)/2; i++ {
		j := len(ret) - i - 1
		ret[i], ret[j] = ret[j], ret[i]
	} // for>
	return ret
}
