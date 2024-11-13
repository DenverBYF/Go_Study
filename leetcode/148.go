package leetcode

// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 将链表分为两部分
	mid := getMid(head)
	rightStart := mid.Next
	mid.Next = nil
	left := sortList(head)
	right := sortList(rightStart)
	return mergeList(left, right)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

func getMid(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
