package bst

import "leetcode_200/common"

func findMidAndPre(n *common.ListNode) (*common.ListNode, *common.ListNode) {
	var pre *common.ListNode
	fast, slow := n, n
	for slow != nil && fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	} // for>
	return slow, pre
}

func sortedListToBST(head *common.ListNode) *common.TreeNode {
	if head == nil {
		return nil
	} // if>

	mid, midPre := findMidAndPre(head)
	if midPre != nil {
		midPre.Next = nil
	} // if>

	root := &common.TreeNode{Val: mid.Val}
	if head == mid {
		// only one node
		return root
	} // if>

	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)
	return root
}
