package linked_list

import "leetcode_200/common"

func isPalindrome(head *common.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	} // if>
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	} // for>

	// 奇数链表取下一点, fast.next == nil
	if fast != nil {
		slow = slow.Next
	} // if>

	// 后半段反转链表
	var pre *common.ListNode
	for slow != nil {
		nextBack := slow.Next
		slow.Next = pre
		pre = slow
		slow = nextBack
	} // for>

	for pre != nil {
		if pre.Val != head.Val {
			return false
		} // if>>
		pre = pre.Next
		head = head.Next
	} // for>
	return true
}
