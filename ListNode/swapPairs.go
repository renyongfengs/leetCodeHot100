package ListNode

//24. 两两交换链表中的节点
//其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//链接：https://leetcode.cn/problems/swap-nodes-in-pairs/description/?envType=study-plan-v2&envId=top-100-liked

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dumpNode := new(ListNode)
	dumpNode.Next = head

	tmp := dumpNode

	for tmp.Next != nil && tmp.Next.Next != nil {
		next1, next2 := tmp.Next, tmp.Next.Next

		tmp.Next = next2
		next1.Next = next2.Next
		next2.Next = next1

		tmp = tmp.Next.Next
	}

	return dumpNode.Next
}
