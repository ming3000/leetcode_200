package linked_list

import "leetcode_200/common"

func mergeTwoLists(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	sentry := &common.ListNode{}
	cur := sentry

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		} // else>>
		cur = cur.Next
		cur.Next = nil
	} // for>

	if l1 != nil {
		cur.Next = l1
	} // if>
	if l2 != nil {
		cur.Next = l2
	} // if>

	return sentry.Next
}
