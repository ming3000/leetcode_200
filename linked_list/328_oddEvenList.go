package linked_list

import "leetcode_200/common"

func oddEvenList(head *common.ListNode) *common.ListNode {
	oddList, evenList := &common.ListNode{}, &common.ListNode{}
	oddCur, evenCur := oddList, evenList

	i := 0
	for head != nil {
		nextBack := head.Next
		if i%2 != 0 {
			oddCur.Next = head
			oddCur = oddCur.Next
			oddCur.Next = nil
		} else {
			evenCur.Next = head
			evenCur = evenCur.Next
			evenCur.Next = nil
		} // else>>
		i++
		head = nextBack
	} // for>

	evenCur.Next = oddList.Next
	return evenList.Next
}
