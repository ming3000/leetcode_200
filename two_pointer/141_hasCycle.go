package two_pointer

import "leetcode_200/common"

func hasCycle(head *common.ListNode) bool {
	if head == nil {
		return false
	} // if>
	slow, fast := head, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		} // if>>
		slow = slow.Next
		fast = fast.Next.Next
	} // for>

	return false
}
