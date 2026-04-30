package ListNode

//148. 排序链表
//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//链接：https://leetcode.cn/problems/sort-list/description/?envType=study-plan-v2&envId=top-100-liked
//讲解：https://leetcode.cn/problems/sort-list/?envType=study-plan-v2&envId=top-100-liked
func sortList(head *ListNode) *ListNode {
	// 如果链表为空或者只有一个节点，无需排序
	if head == nil || head.Next == nil {
		return head
	}

	// 找到中间节点 middle，并断开 middle 与其前一个节点的连接
	middle := middleNode(head)
	//排序
	head = sortList(head)
	middle = sortList(middle)
	// 合并
	return mergeList(head, middle)
}

//找到链表的中间结点 head 的前一个节点，并断开 head 与其前一个节点的连接。
func middleNode(node *ListNode) *ListNode {
	pre, slow, fast := node, node, node
	for fast != nil && fast.Next != nil {
		pre = slow // 记录 slow 的前一个节点
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil // 断开 slow 的前一个节点和 slow 的连接
	return slow
}

func mergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := ListNode{}
	cur := &dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}
