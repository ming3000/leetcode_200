package linked_list

import "leetcode_200/common"

func deleteDuplicates(head *common.ListNode) *common.ListNode {
	sentry := &common.ListNode{Next: head}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		} // else>>
	} // for>

	return sentry.Next
}
