package linked_list

import "leetcode_200/common"

func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	sentry := &common.ListNode{Next: head}

	fast := head
	for n > 0 {
		n--
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return sentry.Next
}
