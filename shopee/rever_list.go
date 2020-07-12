package shopee

import "leetcode_200/common"

func reverseLinkedList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	} // if>

	var pre, cur *common.ListNode
	pre, cur = nil, head
	for cur != nil {
		nextBackUp := cur.Next
		pre = cur.Next
		pre = cur
		cur = nextBackUp
	} // for>
	return pre
}
