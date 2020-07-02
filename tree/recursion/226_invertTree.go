package recursion

import "leetcode_200/common"

func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	} // if>
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
