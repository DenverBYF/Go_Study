package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return mergeKListsHelper(lists, 0, len(lists)-1)
}

func mergeKListsHelper(list []*ListNode, left, right int) *ListNode {
	if left == right {
		return list[left]
	}

	if left > right {
		return nil
	}
	mid := (left + right) / 2
	return mergeTwoList(mergeKListsHelper(list, left, mid), mergeKListsHelper(list, mid+1, right))
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
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
	} else {
		cur.Next = l2
	}
	return dummy.Next
}
