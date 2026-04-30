package ListNode

//206. 反转链表
//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//链接：https://leetcode.cn/problems/reverse-linked-list/?envType=study-plan-v2&envId=top-100-liked
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var (
		preNode *ListNode
		curNode = head
	)
	for curNode.Next != nil {
		nextNode := curNode.Next

		curNode.Next = preNode
		preNode = curNode
		curNode = nextNode
	}
	curNode.Next = preNode

	return curNode
}
