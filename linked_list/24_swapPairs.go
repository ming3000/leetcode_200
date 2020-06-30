package linked_list

import "leetcode_200/common"

func swapPairs(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	} // if>

	sentry := &common.ListNode{Next: head}
	pre := sentry
	for pre != nil && pre.Next != nil && pre.Next.Next != nil {
		first := pre.Next
		second := pre.Next.Next

		first.Next = second.Next
		second.Next = first
		pre.Next = second
		pre = first
	} // for>

	return sentry.Next
}
