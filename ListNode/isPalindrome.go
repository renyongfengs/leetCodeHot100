package ListNode

//234. 回文链表
//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//链接：https://leetcode.cn/problems/palindrome-linked-list/description/?envType=study-plan-v2&envId=top-100-liked

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	//快指针走两步，慢指针走一步，快指针走到底，慢指针到终点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//	反转slow到最后的链表
	var (
		pre *ListNode
	)
	for slow.Next != nil {
		nextNode := slow.Next

		slow.Next = pre
		pre = slow
		slow = nextNode
	}
	slow.Next = pre

	for slow != nil {
		if head.Val != slow.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}
	return true
}
