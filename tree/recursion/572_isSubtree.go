package recursion

import "leetcode_200/common"

func isSubtree(s *common.TreeNode, t *common.TreeNode) bool {
	if s == nil {
		return false
	} // if>

	return isSubtreeWithRoot(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSubtreeWithRoot(s *common.TreeNode, t *common.TreeNode) bool {
	if s == nil && t == nil {
		return true
	} //if>
	if s == nil || t == nil {
		return false
	} //if>
	if s.Val != t.Val {
		return false
	} //if>
	return isSubtreeWithRoot(s.Left, t.Left) && isSubtreeWithRoot(s.Right, t.Right)
}
