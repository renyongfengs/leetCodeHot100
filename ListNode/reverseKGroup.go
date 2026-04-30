package ListNode

//25. K 个一组翻转链表
//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//链接：https://leetcode.cn/problems/reverse-nodes-in-k-group/?envType=study-plan-v2&envId=top-100-liked

func reverseKGroup(head *ListNode, k int) *ListNode {
	//统计节点个数
	n := 0
	for cur := head; cur != nil; {
		n++
		cur = cur.Next
	}

	dump := &ListNode{
		Next: head,
	}
	p0 := dump
	cur := dump.Next
	pre := new(ListNode)
	for ; n >= k; n -= k {

		for i := 0; i < k; i++ {
			next := cur.Next

			cur.Next = pre
			pre = cur
			cur = next
		}

		tmp := dump.Next

		dump.Next.Next = cur
		dump.Next = pre
		dump = tmp
	}

	return p0.Next
}
