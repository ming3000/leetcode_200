package recursion

import "leetcode_200/common"

func mergeTrees(t1 *common.TreeNode, t2 *common.TreeNode) *common.TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} // if>
	if t1 == nil {
		return t2
	} // if>
	if t2 == nil {
		return t1
	} // if>
	newRoot := &common.TreeNode{Val: t1.Val + t2.Val}
	newRoot.Left = mergeTrees(t1.Left, t2.Left)
	newRoot.Right = mergeTrees(t1.Right, t2.Right)
	return newRoot
}
