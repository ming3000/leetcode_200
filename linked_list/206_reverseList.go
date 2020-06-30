package linked_list

import (
	"leetcode_200/common"
)

// 头插法
func reverseList(head *common.ListNode) *common.ListNode {
	sentry := &common.ListNode{}
	for head != nil {
		nextBack := head.Next
		head.Next = sentry.Next
		sentry.Next = head
		head = nextBack
	} // for>
	return sentry.Next
}

// 双指针法
func reverseList2(head *common.ListNode) *common.ListNode {
	var pre, cur *common.ListNode
	pre, cur = nil, head
	for cur != nil {
		nextBack := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextBack
	} // for>
	return pre
}
