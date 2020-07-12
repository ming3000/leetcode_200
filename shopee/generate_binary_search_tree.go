package shopee

import (
	"fmt"
	"sort"
	"strings"
)

// ListNode Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 归并排序无序链表
func sortLinkedList(head *ListNode) *ListNode {
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

	sentry := &ListNode{}
	return mergeLinkedList(sentry, leftList, rightList)
}

func mergeLinkedList(sentry, left, right *ListNode) *ListNode {
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

func midNode(n *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	slow, fast := n, n
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	} // for>
	return pre, slow
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	pre, mid := midNode(head)
	if pre != nil {
		pre.Next = nil
	}

	root := &TreeNode{Val: mid.Val}
	if head == mid {
		// only one node
		return root
	}

	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)
	return root
}

func unsortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	sortedList := sortLinkedList(head)
	return sortedListToBST(sortedList)
}

func GenerateKeyList(in map[int64]interface{}) []interface{} {
	if in == nil {
		return nil
	}
	ret := make([]interface{}, 0)
	for k, _ := range in {
		ret = append(ret, k)
	}
	return ret
}

func GenerateValueList(in map[int64]interface{}) []interface{} {
	if in == nil {
		return nil
	}
	ret := make([]interface{}, 0)
	for _, v := range in {
		ret = append(ret, v)
	}
	return ret
}

func FixStr(inStr string) string {
	if inStr == "" || len(inStr) == 0 {
		return ""
	}

	dstStr := strings.TrimSpace(strings.ToLower(inStr))
	return dstStr
}

func FixStr2(inStr string) string {
	if inStr == "" || len(inStr) == 0 {
		return ""
	}

	lowStr := strings.ToLower(inStr)
	ret := make([]string, 0)
	for i := 0; i < len(lowStr); i++ {
		if (lowStr[i] >= 'a' && lowStr[i] <= 'z') || (lowStr[i] >= '0' && lowStr[i] <= '9') {
			ret = append(ret, fmt.Sprintf("%c", lowStr[i]))
		}
	}
	sort.Strings(ret)
	return strings.Join(ret, "")
}
