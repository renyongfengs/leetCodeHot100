package ListNode

//2. 两数相加
//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//链接：https://leetcode.cn/problems/add-two-numbers/description/?envType=study-plan-v2&envId=top-100-liked
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := new(ListNode)
	forward := (l1.Val + l2.Val) / 10
	head.Val = (l1.Val + l2.Val) % 10
	l1 = l1.Next
	l2 = l2.Next

	node := head
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + forward
		//当前位置值
		node.Next = &ListNode{Val: sum % 10}
		//	进位值
		forward = sum / 10

		node = node.Next

		l1 = l1.Next
		l2 = l2.Next
	}
	//l1和l2的长度可能不一致
	for l1 != nil {
		sum := l1.Val + forward
		//当前位置值
		node.Next = &ListNode{Val: sum % 10}
		//进位值
		forward = sum / 10
		node = node.Next

		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + forward
		node.Next = &ListNode{Val: sum % 10}
		forward = sum / 10
		node = node.Next

		l2 = l2.Next
	}

	if forward > 0 {
		node.Next = &ListNode{Val: forward}
	}
	return head

}
