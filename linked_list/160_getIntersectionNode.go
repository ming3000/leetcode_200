package linked_list

import "leetcode_200/common"

func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	ha, hb := headA, headB
	la, lb := 0, 0
	for ha != nil {
		la++
		ha = ha.Next
	} // for>
	for hb != nil {
		lb++
		hb = hb.Next
	} // for>

	var longList, shortList *common.ListNode
	var lengthDiff int
	if la > lb {
		longList = headA
		shortList = headB
		lengthDiff = la - lb
	} else {
		longList = headB
		shortList = headA
		lengthDiff = lb - la
	} // else>

	for longList != nil {
		if lengthDiff > 0 {
			lengthDiff--
			longList = longList.Next
			continue
		} // if>>
		if longList == shortList {
			return longList
		} else {
			shortList = shortList.Next
			longList = longList.Next
		} // else>>
	} // for>

	return nil
}
