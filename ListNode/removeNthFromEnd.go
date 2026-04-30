package ListNode

//19. 删除链表的倒数第 N 个结点
//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/solutions/2004057/ru-he-shan-chu-jie-dian-liu-fen-zhong-ga-xpfs/?envType=study-plan-v2&envId=top-100-liked
//维持左右指针之间的长度为n,当右节点走到底，左节点正好是倒数n+1节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dumpNode := new(ListNode)
	dumpNode.Next = head //创建一个假指针，处理删除头节点

	right := dumpNode
	for ; n > 0; n-- {
		right = right.Next
	}
	left := dumpNode

	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next

	return dumpNode.Next
}
