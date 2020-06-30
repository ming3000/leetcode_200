package linked_list

import "leetcode_200/common"

func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	sentry := &common.ListNode{Next: head}
	var pre, cur *common.ListNode
	cur = sentry
	count := 1

	for head != nil {
		if count >= n {
			pre = cur
			cur = cur.Next
		} // if>>
		head = head.Next
		count++
	} // for>

	pre.Next = pre.Next.Next
	return sentry.Next
}
