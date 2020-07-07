package shopee

import "leetcode_200/common"

func sortLinkedList(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	} // if>

	// 查找链表中间节点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	} // for>

	rightList := sortLinkedList(slow.Next)
	slow.Next = nil

	leftList := sortLinkedList(head)

	sentry := &common.ListNode{}
	return mergeLinkedList(sentry, leftList, rightList)
}

func mergeLinkedList(sentry, left, right *common.ListNode) *common.ListNode {
	tempHead := sentry
	for left != nil && right != nil {
		if left.Val < right.Val {
			tempHead.Next = left
			left = left.Next
		} else {
			tempHead.Next = right
			right = right.Next
		} // else>>
		tempHead = tempHead.Next
		tempHead.Next = nil
	} // for>

	if left == nil {
		tempHead.Next = right
	}
	if right == nil {
		tempHead.Next = left
	}
	return tempHead.Next
}
