package ListNode

//160. 相交链表
//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
//图示两个链表在节点 c1 开始相交：
//链接：https://leetcode.cn/problems/intersection-of-two-linked-lists/description/?envType=study-plan-v2&envId=top-100-liked
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	//如果A和B相交，需要知道A和B差几个节点
	cur1, cur2 := headA, headB

	length := 0
	for cur1.Next != nil {
		cur1 = cur1.Next
		length++
	}
	for cur2.Next != nil {
		cur2 = cur2.Next
		length--
	}
	//最后一个节点不相等肯定不相交
	if cur1 != cur2 {
		return nil
	}
	//相交，区分出长短节点
	longHead := ternary(length > 0, headA, headB)
	shortHead := ternary(longHead == headA, headB, headA)
	//取绝对值
	if length < 0 {
		length = -length
	}
	//长的节点先走length步
	for length > 0 {
		length--
		longHead = longHead.Next
	}

	for longHead != shortHead {
		longHead = longHead.Next
		shortHead = shortHead.Next
	}

	return longHead
}

func ternary(v bool, a, b *ListNode) *ListNode {
	if v {
		return a
	}
	return b
}

type ListNode struct {
	Val  int
	Next *ListNode
}
