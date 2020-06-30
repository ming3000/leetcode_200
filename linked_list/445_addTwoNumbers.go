package linked_list

import "leetcode_200/common"

func addTwoNumbers445(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	l1Stack := make([]*common.ListNode, 0)
	l2Stack := make([]*common.ListNode, 0)

	l1Back, l2Back := l1, l2
	for l1Back != nil {
		l1Stack = append(l1Stack, l1Back)
		l1Back = l1Back.Next
	} // for>
	for l2Back != nil {
		l2Stack = append(l2Stack, l2Back)
		l2Back = l2Back.Next
	} // for>

	carry := 0
	sentry := &common.ListNode{}
	cur := sentry
	for len(l1Stack) != 0 || len(l2Stack) != 0 || carry != 0 {
		curSum := carry
		if len(l1Stack) != 0 {
			curSum += l1Stack[len(l1Stack)-1].Val
			l1Stack = l1Stack[:len(l1Stack)-1]
		} // if>>
		if len(l2Stack) != 0 {
			curSum += l2Stack[len(l2Stack)-1].Val
			l2Stack = l2Stack[:len(l2Stack)-1]
		} // if>>
		temp := &common.ListNode{Val: curSum % 10, Next: cur.Next}
		cur.Next = temp
		carry = curSum / 10
	} // for>
	return sentry.Next
}
